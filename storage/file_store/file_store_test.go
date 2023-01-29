package file_store

// Copyright (c) Philip Schlump, 2023.
// This file is MIT licensed, see ../../LICENSE.mit

import (
	"sync"
	"testing"

	"github.com/pschlump/fmcsa-svr/config"
	"github.com/pschlump/fmcsa-svr/core"

	"github.com/stretchr/testify/assert"
)

func TestMemoryEngine(t *testing.T) {
	var val int64

	cfg := config.LoadTestConfig()
	cfg.StatFileLocaiton = "./testdata/test-save-cache.json"

	memory := New(cfg)
	err := memory.Init()
	assert.Nil(t, err)

	memory.Add(core.CacheSuccessKey, 10)
	val = memory.Get(core.CacheSuccessKey)
	assert.Equal(t, int64(10), val)
	memory.Add(core.CacheSuccessKey, 10)
	val = memory.Get(core.CacheSuccessKey)
	assert.Equal(t, int64(20), val)

	memory.Set(core.CacheSuccessKey, 0)
	val = memory.Get(core.CacheSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			memory.Add(core.CacheSuccessKey, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	val = memory.Get(core.CacheSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, memory.Close())
}
