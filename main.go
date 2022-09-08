package main

// See LICENSE - Apache 2.0 Licnesed.
// Copyright (C) Philip Schlump, 2022.

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/brandenc40/qcmobile"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/pschlump/dbgo"
)

var HostPort = flag.String("hostport", "127.0.0.1:10042", "Host/Port to listen on")
var Dir = flag.String("dir", "./www", "Directory to server static assets from")
var ValidationKey = flag.String("key", "FMCSA.l4BUEAzcBvqu8+C1JFxi8vMFr5g=", "Authentication Key")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "fmcsa : Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()
	if len(fns) != 0 {
		fmt.Printf("Extra arguments are not supported [%s]\n", fns)
		os.Exit(1)
	}

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile(*Dir, true)))
	router.GET("/api/v1/status", func(c *gin.Context) {
		c.JSON(http.StatusOK /*200*/, gin.H{
			"status": "success",
			"msg":    "Hello Silly World!",
		})
	})
	router.GET("/api/v1/mc-number-data", func(c *gin.Context) {
		// xyzzy TODO

		// xyzzy - Validate callers key - get headers.
		// xyzzy - get parameters mc=
		// xyzzy - cleanup mc number so if "MC 43565" just use the number, trim, spaces MC- remove etc.

		Key := os.Getenv("xyzzy")

		cfg := qcmobile.Config{
			Key:        Key,
			HTTPClient: &http.Client{},
		}
		client := qcmobile.NewClient(cfg)
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
		carrier, err := client.GetCarrier(ctx, "53467")
		if err != nil {
			c.JSON(http.StatusOK /*200*/, gin.H{
				"status": "error",
				"msg":    "Invalid MC number",
			})
			return
		}

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(http.StatusOK /*200*/, `{"status":"success",`+dbgo.SVarI(carrier)+"}\n")
		return

		//c.JSON(http.StatusOK /*200*/, gin.H{
		//	"status": "success",
		//	"msg":    "Hello Silly World!",
		//})
	})
	router.GET("/status", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(http.StatusOK /*200*/, dbgo.SVarI(c))
	})
	router.Run(*HostPort) // listen and serve on 0.0.0.0:9090
}
