package redis

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pschlump/fmcsa-svr/config"

	// "github.com/go-redis/redis/v9"
	"github.com/go-redis/redis"
)

// New func implements the storage interface for fmcsa_svr (https://github.com/pschlump/fmcsa_svr)
func New(config *config.ConfigData) *Storage {
	return &Storage{
		ctx:    context.Background(),
		config: config,
	}
}

// Storage is interface structure
type Storage struct {
	ctx    context.Context
	config *config.ConfigData
	client redis.Cmdable
}

func (s *Storage) Add(key string, count int64) {
	// s.client.IncrBy(s.ctx, key, count)
	s.client.IncrBy(key, count)
}

func (s *Storage) Set(key string, count int64) {
	// s.client.Set(s.ctx, key, count, 0)
	s.client.Set(key, count, 0)
}

func (s *Storage) Get(key string) int64 {
	// val, _ := s.client.Get(s.ctx, key).Result()
	val, _ := s.client.Get(key).Result()
	count, _ := strconv.ParseInt(val, 10, 64)
	return count
}

// Init client storage.
// Uses:
//	RedisConnectHost string `json:"redis_host" default:"$ENV$REDIS_HOST"`
//	RedisConnectAuth string `json:"redis_auth" default:"$ENV$REDIS_AUTH"`
//	RedisConnectPort string `json:"redis_port" default:"6379"`
//	RedisCluster string `json:"redis_cluster" default:"no"`
func (s *Storage) Init() error {
	if s.config.RedisCluster == "no" {
		s.client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    strings.Split(s.config.RedisConnectHost, ","),
			Password: s.config.RedisConnectAuth,
		})
	} else {
		s.client = redis.NewClient(&redis.Options{
			Addr:     s.config.RedisConnectHost,
			Password: s.config.RedisConnectAuth,
			DB:       1,
		})
	}

	// if err := s.client.Ping(s.ctx).Err(); err != nil {
	if err := s.client.Ping().Err(); err != nil {
		return err
	}

	return nil
}

// Close the storage connection
func (s *Storage) Close() error {
	switch v := s.client.(type) {
	case *redis.Client:
		return v.Close()
	case *redis.ClusterClient:
		return v.Close()
	case nil:
		return nil
	default:
		// this will not happen anyway, unless we mishandle it on `Init`
		panic(fmt.Sprintf("invalid redis client: %v", reflect.TypeOf(v)))
	}
}
