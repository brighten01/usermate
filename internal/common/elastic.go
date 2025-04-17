package common

import (
	"errors"
	"time"
)

var (
	// ErrInvalidConfig 配置错误
	ErrInvalidConfig = errors.New("invalid elasticsearch configuration")
	// ErrClientCreation 客户端创建错误
	ErrClientCreation = errors.New("failed to create elasticsearch client")
	// ErrDataUnmarshal 数据解析错误
	ErrDataUnmarshal = errors.New("failed to unmarshal data")
	// ErrDataWrite 数据写入错误
	ErrDataWrite = errors.New("failed to write data")
	// ErrDataSearch 数据搜索错误
	ErrDataSearch = errors.New("failed to search data")
)

const (
	// DefaultPageSize 默认分页大小
	DefaultPageSize = 10
	// DefaultPage 默认页码
	DefaultPage = 1
	// OrderIndex 订单索引名称
	OrderIndex = "orders"
	// DefaultCacheExpiration 默认缓存过期时间
	DefaultCacheExpiration = 5 * time.Minute
	// DefaultCacheCleanupInterval 默认缓存清理间隔
	DefaultCacheCleanupInterval = 10 * time.Minute
)
