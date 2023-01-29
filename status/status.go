package status

import (
	"errors"
	"os"

	"github.com/pschlump/dbgo"
	"github.com/pschlump/fmcsa-svr/config"
	"github.com/pschlump/gorush/storage/file_store"
	"github.com/pschlump/gorush/storage/memory"
	"github.com/pschlump/gorush/storage/redis"
	"github.com/thoas/stats"
)

var logFilePtr = os.Stderr

// Stats provide response time, status code count, etc.
var Stats *stats.Stats

// StatStorage implements the storage interface
var StatStorage *StateStorage

// App is status structure
type App struct {
	Version        string      `json:"version"`
	SuccessTasks   int         `json:"success_tasks"`
	FailureTasks   int         `json:"failure_tasks"`
	SubmittedTasks int         `json:"submitted_tasks"`
	TotalCount     int64       `json:"total_count"`
	Fmcsa          FmcsaStatus `json:"fmcsa"`
}

// AndroidStatus is android structure
type FmcsaStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

// InitAppStatus for initialize app status
func InitAppStatus(conf *config.ConfigData) error {

	// logx.LogAccess.Info("Init App Status Engine as ", conf.StatEngine)
	dbgo.Fprintf(logFilePtr, "Init App Status Engine as %s\n", conf.StatEngine)

	// var store core.Storage
	//nolint:goconst
	switch conf.StatEngine {
	case "memory":
		store = memory.New()
	case "redis":
		store = redis.New(conf)
	case "file":
		store = file_store.New(conf)
	default:
		// logx.LogError.Error("storage error: can't find storage driver")
		dbgo.Fprintf(logFilePtr, "storage error: invalid storage driver, should be 'memory', 'redis', 'file', is:%s\n", conf.StatEngine)
		return errors.New("can't find storage driver")
	}

	StatStorage = NewStateStorage(store)

	if err := StatStorage.Init(); err != nil {
		// logx.LogError.Error("storage error: " + err.Error())
		dbgo.Fprintf(logFilePtr, "storage error: %s\n", err)
		return err
	}

	Stats = stats.New()

	return nil
}
