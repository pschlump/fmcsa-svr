package status

// Copyright (c) Philip Schlump, 2023.
// This file is MIT licensed, see ../LICENSE.mit

import (
	"github.com/pschlump/fmcsa-svr/core"
)

type StateStorage struct {
	store core.Storage
}

func NewStateStorage(store core.Storage) *StateStorage {
	return &StateStorage{
		store: store,
	}
}

func (s *StateStorage) Init() error {
	return s.store.Init()
}

func (s *StateStorage) Close() error {
	return s.store.Close()
}

// PJS - generated function.

// Reset Client storage.
func (s *StateStorage) Reset() {
	s.store.Set(core.TotalCountKey, 0)
	s.store.Set(core.FmcsaSuccessKey, 0)
	s.store.Set(core.FmcsaErrorKey, 0)
	s.store.Set(core.CacheSuccessKey, 0)
	s.store.Set(core.CacheErrorKey, 0)
}

// PJS - generated functions.

// AddTotalCount record push notification count.
func (s *StateStorage) AddTotalCount(count int64) {
	s.store.Add(core.TotalCountKey, count)
}

// AddFmcsaSuccess record push notification count.
func (s *StateStorage) AddFmcsaSuccess(count int64) {
	s.store.Add(core.FmcsaSuccessKey, count)
}

// AddFmcsaError record push notification count.
func (s *StateStorage) AddFmcsaError(count int64) {
	s.store.Add(core.FmcsaErrorKey, count)
}

// AddCacheSuccess record push notification count.
func (s *StateStorage) AddCacheSuccess(count int64) {
	s.store.Add(core.CacheSuccessKey, count)
}

// AddCacheError record push notification count.
func (s *StateStorage) AddCacheError(count int64) {
	s.store.Add(core.CacheErrorKey, count)
}

// GetTotalCount show success counts of total requests to server.
func (s *StateStorage) GetTotalCount() int64 {
	return s.store.Get(core.TotalCountKey)
}

// GetFmcsaSuccess show success counts of total requests to FMCSA remote server.
func (s *StateStorage) GetFmcsaSuccess() int64 {
	return s.store.Get(core.FmcsaSuccessKey)
}

// GetFmcsaError show success counts of total requests to FMCSA remote server.
func (s *StateStorage) GetFmcsaError() int64 {
	return s.store.Get(core.FmcsaErrorKey)
}

// GetCacheSuccess show success counts of total requests to FMCSA remote server.
func (s *StateStorage) GetCacheSuccess() int64 {
	return s.store.Get(core.FmcsaSuccessKey)
}

// GetCacheError show success counts of total requests to FMCSA remote server.
func (s *StateStorage) GetCacheError() int64 {
	return s.store.Get(core.CacheSuccessKey)
}
