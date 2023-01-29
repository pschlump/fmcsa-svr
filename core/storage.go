package core

const (
	TotalCountKey   = "fmcsa-svr-total-count-key"
	FmcsaSuccessKey = "fmcsa-svr-fmcsa-success-key"
	FmcsaErrorKey   = "fmcsa-svr-fmcsa-error-key"
	CacheSuccessKey = "fmcsa-svr-cache-success-key"
	CacheErrorKey   = "fmcsa-svr-cache-error-key"
)

// Storage interface
type Storage interface {
	Init() error
	Add(key string, count int64)
	Set(key string, count int64)
	Get(key string) int64
	Close() error
}
