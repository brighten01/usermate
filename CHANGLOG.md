
//接入nacos  不分代码
```azure
package test

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
	"strings"
	"usermate/internal/conf"
	"usermate/pkg/zlog"
)

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"usermate/internal/data"

	"os"
	"strings"
	"usermate/pkg/zlog"

	"github.com/nacos-group/nacos-sdk-go/clients/config_client"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/go-kratos/kratos/v2/registry"

	nacosConfig "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"usermate/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "genie.infra.v1"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	nacos            bool
	nacosLogDir      string
	nacosCacheDir    string
	nacosServer      string
	nacosPort        uint64
	nacosNamespaceId string
	nacosGroupId     string = "infra-service"
	dataIds                 = []string{
		"config.yaml",
		"registry.yaml",
	}
	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf=config.yaml")
	flag.BoolVar(&nacos, "nacos", false, "use nacos, eg: -nacos=true")
	flag.StringVar(&nacosServer, "nacosServer", "127.0.0.1", "nacos host, eg: -nacosServer=127.0.0.1")
	flag.Uint64Var(&nacosPort, "nacosPort", 8848, "nacos port, eg: -nacosPort 8488")
	flag.StringVar(&nacosNamespaceId, "nacosNamespaceId", "namespaceId", "nacos namespaceId, eg: -nacosNamespaceId=id")
	flag.StringVar(&nacosCacheDir, "nacosCacheDir", "./logs", "nacos cacheDir, eg: -nacosCacheDir=./logs")
	flag.StringVar(&nacosLogDir, "nacosLogDir", "./logs", "nacos logDir, eg: -nacoslogDir=./logs")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()

	var c config.Config
	if nacos {
		nacosClient := NewConfigClient(nacosServer, nacosPort, nacosNamespaceId, nacosLogDir, nacosCacheDir)
		configSources := make([]config.Source, 0)
		for _, dataId := range dataIds {
			configSources = append(configSources, nacosConfig.NewConfigSource(nacosClient, nacosConfig.WithGroup(nacosGroupId), nacosConfig.WithDataID(dataId)))
		}
		c = config.New(
			config.WithSource(configSources...),
		)
	} else {
		localDir := flagconf
		// 判断working dir
		dir, _ := os.Getwd()
		if !strings.Contains(dir, "/cmd/") {
			localDir = "./configs"
		}
		c = config.New(
			config.WithSource(
				file.NewSource(localDir),
			),
		)
	}
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	zlog.Init(Name, bc.Log.Filename, int(bc.Log.MaxSize), int(bc.Log.MaxBackup), int(bc.Log.MaxAge), bc.Log.Compress)
	defer zlog.Sync()
	logger := log.With(zlog.NewZapLogger(zlog.STDInstance()),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
		//"dayu_trace_id", dayu_trace.TraceID(),
	)

	if err := data.NewTrace(bc.Jaeger); err != nil {
		panic(err)
	}
	c.Watch("oss", func(key string, value config.Value) {
		value.Scan(&bc.Oss)
		fmt.Printf("oss change: %+v\n", bc.Oss)
	})

	app, cleanup, err := initApp(bc.Server, bc.Data, bc.Nacos, bc.Oss, bc.Services, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func NewConfigClient(address string, port uint64, namespaceId string, logDir string, cacheDir string) config_client.IConfigClient {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(address, port),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         namespaceId, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              logDir,
		CacheDir:            cacheDir,
		LogLevel:            "debug",
	}
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	return client
}

```