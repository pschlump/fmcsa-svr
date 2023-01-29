package file_store

// Copyright (c) Philip Schlump, 2023.
// This file is MIT licensed, see ../../LICENSE.mit

import (
	"sync"

	"github.com/pschlump/fmcsa-svr/config"
	"go.uber.org/atomic"
)

// New func implements the storage interface for fmcsa-svr (https://github.com/pschlump/fmcsa-svr)
func New(gCfg *config.ConfigData) *Storage {
	return &Storage{
		storagerFileName: gCfg.StatFileLocaiton,
	}
}

// Storage is interface structure
type Storage struct {
	mem              sync.Map
	storagerFileName string
}

func (s *Storage) getValueBtKey(key string) *atomic.Int64 {
	if val, ok := s.mem.Load(key); ok {
		return val.(*atomic.Int64)
	}
	val := atomic.NewInt64(0)
	s.mem.Store(key, val)
	return val
}

func (s *Storage) Add(key string, count int64) {
	s.getValueBtKey(key).Add(count)
}

func (s *Storage) Set(key string, count int64) {
	s.getValueBtKey(key).Store(count)
}

func (s *Storage) Get(key string) int64 {
	return s.getValueBtKey(key).Load()
}

// Init client storage.
func (*Storage) Init() error {
	return nil
}

// Close the storage connection
func (*Storage) Close() error {
	return nil
}
