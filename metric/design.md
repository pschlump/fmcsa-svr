
Take `metric/track-items.json` and for each item

(1). Create a Get{name} Set{name} function for ...



(2). Create constants for ... into ./core/core-constants-generated.go
	--domain fmcsa-svr
	Cammel to Snake Case | s/_/-/g + "-key"

```
package core

const (
	TotalCountKey   = "fmcsa-svr-total-count-key"
	FmcsaSuccessKey = "fmcsa-svr-fmcsa-success-key"
	FmcsaErrorKey   = "fmcsa-svr-fmcsa-error-key"
	CacheSuccessKey = "fmcsa-svr-cache-success-key"
	CacheErrorKey   = "fmcsa-svr-cache-error-key"
)


```
