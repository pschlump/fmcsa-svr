package memory

// Copyright (c) Philip Schlump, 2023.
// This file is MIT licensed, see ../../LICENSE.mit

import (
	"os"
	"sync"

	"github.com/pschlump/dbgo"
	"go.uber.org/atomic"
)

// New func implements the storage interface for fmcsa-svr (https://github.com/pschlump/fmcsa-svr)
func New() *Storage {
	return &Storage{
		Data: make(map[string]interface{}),
	}
}

// Storage is interface structure
type Storage struct {
	// Mem sync.Map
	Data map[string]interface{}
	Safe sync.Mutex
}

func (s *Storage) getValueBtKey(key string) *atomic.Int64 {
	s.Safe.Lock()
	defer s.Safe.Unlock()
	// if val, ok := s.Mem.Load(key); ok {
	if val, ok := s.Data[key]; ok {
		return val.(*atomic.Int64)
	}
	val := atomic.NewInt64(0)
	// s.Mem.Store(key, val)
	s.Data[key] = val
	return val
}

func (s *Storage) Add(key string, count int64) {
	s.getValueBtKey(key).Add(count)
	if db1 {
		dbgo.Fprintf(os.Stderr, "at:%(LF) key ->%s<- add %d data-after %s\n", key, count, dbgo.SVarI(s.Data))
	}
}

func (s *Storage) Set(key string, count int64) {
	s.getValueBtKey(key).Store(count)
	if db1 {
		dbgo.Printf("%(red)Data is: %s at:%(LF)\n", dbgo.SVarI(s.Data))
	}
}

func (s *Storage) Get(key string) int64 {
	if db1 {
		dbgo.Printf("%(red)Data is: %s at:%(LF)\n", dbgo.SVarI(s.Data))
	}
	return s.getValueBtKey(key).Load()
}

// Init client storage.
func (*Storage) Init() error {
	return nil
}

// Close the storage connection
func (s *Storage) Close() error {
	dbgo.Printf("%(red)In Call Close() Data is: %s at:%(LF)\n", dbgo.SVarI(s.Data))
	return nil
}

const db1 = false
