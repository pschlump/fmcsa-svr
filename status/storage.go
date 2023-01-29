package status

import (
	"github.com/pschlump/gorush/core"
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

// Reset Client storage.
func (s *StateStorage) Reset() {
	s.store.Set(core.TotalCountKey, 0)
	s.store.Set(core.IosSuccessKey, 0)
	s.store.Set(core.IosErrorKey, 0)
	s.store.Set(core.AndroidSuccessKey, 0)
	s.store.Set(core.AndroidErrorKey, 0)
	s.store.Set(core.HuaweiSuccessKey, 0)
	s.store.Set(core.HuaweiErrorKey, 0)
}

// AddTotalCount record push notification count.
func (s *StateStorage) AddTotalCount(count int64) {
	s.store.Add(core.TotalCountKey, count)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *StateStorage) AddIosSuccess(count int64) {
	s.store.Add(core.IosSuccessKey, count)
}

// AddIosError record counts of error iOS push notification.
func (s *StateStorage) AddIosError(count int64) {
	s.store.Add(core.IosErrorKey, count)
}
