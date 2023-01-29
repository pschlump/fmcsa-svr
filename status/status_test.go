package status

// Copyright (c) Philip Schlump, 2023.
// This file is MIT licensed, see ../LICENSE.mit

import (
	"os"
	"testing"

	"github.com/pschlump/fmcsa-svr/config"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestStorageDriverExist(t *testing.T) {
	cfg := config.LoadTestConfig()
	cfg.StatEngine = "should-procude-error"
	err := InitAppStatus(cfg)
	assert.Error(t, err)
}

func TestStatForMemoryEngine(t *testing.T) {
	var val int64
	cfg := config.LoadTestConfig()
	cfg.StatEngine = "memory"
	err := InitAppStatus(cfg)
	assert.Nil(t, err)

	StatStorage.AddTotalCount(100)
	StatStorage.AddFmcsaSuccess(200)
	StatStorage.AddFmcsaError(300)
	StatStorage.AddCacheSuccess(400)
	StatStorage.AddCacheError(500)

	val = StatStorage.GetTotalCount()
	assert.Equal(t, int64(100), val)
	val = StatStorage.GetFmcsaSuccess()
	assert.Equal(t, int64(200), val)
	val = StatStorage.GetFmcsaError()
	assert.Equal(t, int64(300), val)
	val = StatStorage.GetCacheSuccess()
	assert.Equal(t, int64(400), val)
	val = StatStorage.GetCacheError()
	assert.Equal(t, int64(500), val)
}

/*
func TestRedisServerSuccess(t *testing.T) {
	cfg := config.LoadTestConfig()

	err := InitAppStatus(cfg)

	assert.NoError(t, err)
}

func TestRedisServerError(t *testing.T) {
	cfg := config.LoadTestConfig()

	err := InitAppStatus(cfg)

	assert.Error(t, err)
}

func TestStatForRedisEngine(t *testing.T) {
	var val int64
	cfg := config.LoadTestConfig()
	err := InitAppStatus(cfg)
	assert.Nil(t, err)

	assert.Nil(t, StatStorage.Init())
	StatStorage.Reset()

	StatStorage.AddTotalCount(100)
	StatStorage.AddIosSuccess(200)
	StatStorage.AddIosError(300)
	StatStorage.AddAndroidSuccess(400)
	StatStorage.AddAndroidError(500)

	val = StatStorage.GetTotalCount()
	assert.Equal(t, int64(100), val)
	val = StatStorage.GetIosSuccess()
	assert.Equal(t, int64(200), val)
	val = StatStorage.GetIosError()
	assert.Equal(t, int64(300), val)
	val = StatStorage.GetAndroidSuccess()
	assert.Equal(t, int64(400), val)
	val = StatStorage.GetAndroidError()
	assert.Equal(t, int64(500), val)
}

func TestDefaultEngine(t *testing.T) {
	var val int64
	// defaul engine as memory
	cfg := config.LoadTestConfig()
	err := InitAppStatus(cfg)
	assert.Nil(t, err)

	StatStorage.Reset()

	StatStorage.AddTotalCount(100)
	StatStorage.AddIosSuccess(200)
	StatStorage.AddIosError(300)
	StatStorage.AddAndroidSuccess(400)
	StatStorage.AddAndroidError(500)

	val = StatStorage.GetTotalCount()
	assert.Equal(t, int64(100), val)
	val = StatStorage.GetIosSuccess()
	assert.Equal(t, int64(200), val)
	val = StatStorage.GetIosError()
	assert.Equal(t, int64(300), val)
	val = StatStorage.GetAndroidSuccess()
	assert.Equal(t, int64(400), val)
	val = StatStorage.GetAndroidError()
	assert.Equal(t, int64(500), val)
}

func TestStatForBoltDBEngine(t *testing.T) {
	var val int64
	cfg := config.LoadTestConfig()
	cfg.Stat.Engine = "boltdb"
	err := InitAppStatus(cfg)
	assert.Nil(t, err)

	StatStorage.Reset()

	StatStorage.AddTotalCount(100)
	StatStorage.AddIosSuccess(200)
	StatStorage.AddIosError(300)
	StatStorage.AddAndroidSuccess(400)
	StatStorage.AddAndroidError(500)

	val = StatStorage.GetTotalCount()
	assert.Equal(t, int64(100), val)
	val = StatStorage.GetIosSuccess()
	assert.Equal(t, int64(200), val)
	val = StatStorage.GetIosError()
	assert.Equal(t, int64(300), val)
	val = StatStorage.GetAndroidSuccess()
	assert.Equal(t, int64(400), val)
	val = StatStorage.GetAndroidError()
	assert.Equal(t, int64(500), val)
}
*/
