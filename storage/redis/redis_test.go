package redis

// Copyright (c) Philip Schlump, 2023.
// This file is MIT licensed, see ../../LICENSE.mit

import (
	"sync"
	"testing"

	"github.com/pschlump/fmcsa-svr/config"
	"github.com/pschlump/fmcsa-svr/core"

	"github.com/stretchr/testify/assert"
)

func TestRedisServerError(t *testing.T) {
	cfg := config.LoadTestConfig()

	redis := New(cfg)
	err := redis.Init()

	assert.Error(t, err)
}

func TestRedisEngine(t *testing.T) {
	var val int64

	cfg := config.LoadTestConfig()

	redis := New(cfg)
	err := redis.Init()
	assert.Nil(t, err)

	redis.Add(core.HuaweiSuccessKey, 10)
	val = redis.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	redis.Add(core.HuaweiSuccessKey, 10)
	val = redis.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	redis.Set(core.HuaweiSuccessKey, 0)
	val = redis.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			redis.Add(core.HuaweiSuccessKey, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	val = redis.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, redis.Close())
}
