# Kratos Project Template

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker


```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

### kafka 创建topic 
```
kafka-topics --create --topic order-topic --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092
```
###topic 列表

```
 kafka-topics --list  --bootrap-server localhost:9092

```

### order mapping 
```
PUT orders/_mappings
{
    "properties": {
      "id":{
        "type":"integer",
        "index": true      
      },
      
      "order_id":{
        "type":"text",
        "index":true
      },
      
      "uid":{
        "type":"integer",
        "index":true 
      },
      
      "user_mate_id":{
        "type":"integer",
        "index":true
      },
      
      "service_category_id":{
        "type":"integer",
        "index":true
      },
      
      "start_time":{
        "type":"date",
        "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
      },
      
      "end_time":{
        "type":"date",
        "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
      },
      
      "status":{
        "type":"integer"
      },
      "amount":{
        "type":"float",
        "index": false
      },
      
      "discount":{
        "type":"integer",
        "index": false
      },
      
      "avatar":{
        "type":"text",
        "index":false
      },
      
      "link_url":{
        "type":"text",
        "index" :false
      },
      
      "is_order_after":{
        "type":"integer"
      },
      
      "gender":{
        "type":"integer"
      },
      
      "level":{
        "type":"integer",
        "index":false
      },
      
      "duration":{
        "type":"integer",
        "index":false
      },
      
      "service_category_name":{
        "type":"text",
        "index":false
      },
      
      "wechat":{
        "type":"text",
        "index":false
      },
      
      "note":{
        "type":"text",
        "index":false
      },
      "payment":{
        "type":"integer",
        "index":false
      }
      
    }
}
```