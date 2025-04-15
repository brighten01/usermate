package utils

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	epoch          = 1672531200000 // 2023-01-01 00:00:00 UTC
	datacenterBits = 5
	workerBits     = 5
	sequenceBits   = 12
	maxSequence    = 1<<sequenceBits - 1
)

type Snowflake struct {
	mu            sync.Mutex
	lastTimestamp int64
	sequence      int64
	datacenterID  int64
	workerID      int64
	redisClient   *redis.Client
}

func New(datacenterID, workerID int64, redisAddr string) *Snowflake {
	return &Snowflake{
		datacenterID: datacenterID,
		workerID:     workerID,
		redisClient:  redis.NewClient(&redis.Options{Addr: redisAddr}),
	}
}

func (s *Snowflake) NextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / 1e6
	ctx := context.Background()

	if now < s.lastTimestamp {
		return 0, errors.New("clock moved backwards")
	}

	if now == s.lastTimestamp {
		// 使用Redis原子操作获取序列号
		seq, err := s.redisClient.Incr(ctx, "snowflake:sequence").Result()
		if err != nil {
			return 0, err
		}
		s.sequence = seq & maxSequence
		if s.sequence == 0 {
			now = s.waitNextMillis(now)
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = now

	// 组合ID各部分‌:ml-citation{ref="6,7" data="citationList"}
	id := ((now - epoch) << (datacenterBits + workerBits + sequenceBits)) |
		(s.datacenterID << (workerBits + sequenceBits)) |
		(s.workerID << sequenceBits) |
		s.sequence

	return id, nil
}

func (s *Snowflake) waitNextMillis(last int64) int64 {
	now := time.Now().UnixNano() / 1e6
	for now <= last {
		time.Sleep(100 * time.Microsecond)
		now = time.Now().UnixNano() / 1e6
	}
	return now
}

func GenerateId() string {
	workerId := int64(math.Abs(float64(rand.Int())))
	snowflack := New(111, workerId, "")
	nextId, err := snowflack.NextID()
	if err != nil {
		//toodo log
	}
	nextIdstr := strconv.Itoa(int(nextId))
	order_pre := time.Now().Format("20060102150304")
	var builder strings.Builder
	builder.WriteString(order_pre)
	builder.WriteString(nextIdstr)
	return builder.String()

}
