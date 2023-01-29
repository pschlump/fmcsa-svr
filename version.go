package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

var version string

// SetVersion for setup version string.
func SetVersion(ver string) {
	version = ver
}

// GetVersion for get current version.
func GetVersion() string {
	return version
}

// PrintVersion provide print server engine
func PrintVersion(fp *os.File) {
	fmt.Fprintf(fp, "Fmcsa-Svr %s, Compiler: %s %s, Copyright (C) 2023 Philip Schlump.\n", version, runtime.Compiler, runtime.Version())
}

// VersionMiddleware : add version on header.
func VersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		c.Header("X-FMCDA-SVR-VERSION", version)
		c.Next()
	}
}
