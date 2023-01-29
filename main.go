package main

// See LICENSE - Apache 2.0 Licnesed.
// Copyright (C) Philip Schlump, 2022.

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/brandenc40/qcmobile"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/pschlump/ReadConfig"
	"github.com/pschlump/dbgo"
	"github.com/pschlump/filelib"
	"github.com/pschlump/fmcsa-svr/config"
	"github.com/pschlump/fmcsa-svr/metric"
	"github.com/pschlump/fmcsa-svr/status"
)

var HostPort = flag.String("hostport", "127.0.0.1:10042", "Host/Port to listen on")
var Dir = flag.String("dir", "./www", "Directory to server static assets from")
var ValidationKey = flag.String("key", "FMCSA.l4BUEAzcBvqu8_C1JFxi8vMFr5g", "Authentication Key")
var DbFlagParam = flag.String("db_flag", "", "Additional Debug Flags")
var Version = flag.Bool("version", false, "Report version of code and exit")
var Comment = flag.String("comment", "", "Unused comment for ps.")
var Cache = flag.String("cache", "./cache", "Cached Data based on previous calls")
var Cfg = flag.String("cfg", "cfg.json", "config file for this call")

var DbOn map[string]bool = make(map[string]bool)
var logFilePtr = os.Stderr
var gCfg config.ConfigData

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

	if Cfg == nil {
		fmt.Printf("--cfg is a required parameter\n")
		os.Exit(1)
	}

	// logx.SetLogFilePtr(os.Stderr)

	if *Version {
		fmt.Printf("Version (Git Commit): %s\n", GitCommit)
		os.Exit(0)
	}

	SetVersion(GitCommit)

	PrintVersion(logFilePtr)

	// -----------------------------------------------------------------------------------
	// Save PID to log directory for possible kill later
	// -----------------------------------------------------------------------------------
	os.Mkdir("./log", 0755)
	ioutil.WriteFile("./log/pid_file", []byte(fmt.Sprintf("%d\n", os.Getpid())), 0644)

	// ------------------------------------------------------------------------------
	// Read in Configuration
	// ------------------------------------------------------------------------------
	err := ReadConfig.ReadFile(*Cfg, &gCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read confguration: %s error %s\n", *Cfg, err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", dbgo.SVarI(gCfg))

	status.InitAppStatus(&gCfg)

	DebugFlagProcess(*DbFlagParam, DbOn)
	os.MkdirAll(*Cache, 0755)

	if gCfg.StatEngine == "file" {
		os.MkdirAll(filepath.Dir(gCfg.StatFileLocaiton), 0755) // string `json:"stat_file_location" default:"./data/stat.json"`
	}

	// ------------------------------------------------------------------------------
	//fmt.Printf("DbOn=%s\n", dbgo.SVarI(DbOn))
	//os.Exit(1)
	// ------------------------------------------------------------------------------

	Key := os.Getenv("FMCSA_WebKey")
	if Key == "" {
		dbgo.Fprintf(os.Stderr, "Not setup correctly - missing environment variable 'FMCSA_WebKey'\n")
		os.Exit(1)
	}

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile(*Dir, true)))

	// ------------------------------------------------------------------------------
	// Support metrics
	// ------------------------------------------------------------------------------
	m := metric.NewMetrics()
	prometheus.MustRegister(m)

	// ------------------------------------------------------------------------------
	// Create Routes
	// ------------------------------------------------------------------------------
	router.GET("/api/v1/status", func(c *gin.Context) {
		c.JSON(http.StatusOK /*200*/, gin.H{
			"status": "success",
			"msg":    "Hello Silly World!",
		})
	})
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK /*200*/, gin.H{
			"status": "success",
			"msg":    "Hello Silly World!",
		})
	})

	router.GET("/api/v1/mc-number-data", func(c *gin.Context) {

		type ApiGetMc struct {
			Mc string `json:"mc" form:"mc" binding:"required"`
		}

		status.StatStorage.AddTotalCount(int64(1))

		auth := c.Request.Header.Get("X-Authentication")
		if db8 {
			fmt.Printf("auth=->%s<- %T ->%s<-\n", auth, auth, dbgo.SVarI(c.Request.Header))
		}
		if auth != *ValidationKey {
			c.JSON(http.StatusUnauthorized /*401*/, gin.H{
				"status": "error",
				"msg":    "401 - Invalid x-authentication header",
			})
			return
		}

		Key := os.Getenv("FMCSA_WebKey")

		var pp ApiGetMc
		if err := BindFormOrJSON(c, &pp); err != nil {
			return
		}

		fmt.Printf("At Top mc=%+v\n", pp)

		// cleanup mc number so if "MC 43565" just use the number, trim, spaces MC- remove etc.
		var re = regexp.MustCompile(`(^[ \t]*)([mM][cC])?([^0-9]*)?`)
		pp.Mc = re.ReplaceAllString(pp.Mc, "")
		fmt.Printf("After mc=->%s<-\n", pp.Mc)

		var carrierData string
		carrierFileName := fmt.Sprintf("%s/%s.json", *Cache, pp.Mc)
		cacheHit := false
		dbgo.Printf("%(yellow)file name ->%s<- at:%(LF)\n", carrierFileName)

		if filelib.Exists(carrierFileName) {
			if buf, err := ioutil.ReadFile(carrierFileName); err != nil {
				dbgo.Fprintf(os.Stderr, "%(red)Error reading cached file ->%s<- error :%s\n", carrierFileName, err)
				status.StatStorage.AddCacheError(int64(1))
			} else {
				cacheHit = true
				carrierData = string(buf)
				dbgo.Printf("%(green)Got it %s from cache\n", pp.Mc)
				status.StatStorage.AddCacheSuccess(int64(1))
			}
		}

		if !cacheHit {
			cfg := qcmobile.Config{
				Key:        Key,
				HTTPClient: &http.Client{},
			}
			client := qcmobile.NewClient(cfg)
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			defer cancel()
			xcarrier, err := client.GetCarrier(ctx, pp.Mc)
			if err != nil {
				status.StatStorage.AddFmcsaError(int64(1))
				c.JSON(http.StatusOK /*200*/, gin.H{
					"status": "error",
					"msg":    "Invalid MC number",
				})
				return
			}
			status.StatStorage.AddFmcsaSuccess(int64(1))
			carrierData = dbgo.SVarI(xcarrier)
			dbgo.Printf("%(yellow)Got it %s from server\n", pp.Mc)
		}

		dbgo.Printf("%(yellow)Update cache for %s file name ->%s<- at:%(LF)\n", pp.Mc, carrierFileName)
		ioutil.WriteFile(carrierFileName, []byte(carrierData), 0644)

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(http.StatusOK /*200*/, `{"status":"success","data":`+carrierData+"}\n")
		return
	})

	// Prometheus metric data
	router.GET("/metric", metricsHandler)
	router.GET("/status-server", appStatusHandler())

	router.Run(*HostPort) // listen and serve on 0.0.0.0:9090
}

func DebugFlagProcess(DbFlagParam string, DbOn map[string]bool) {
	for _, s := range strings.Split(DbFlagParam, ",") {
		DbOn[s] = true
	}
}

func BindFormOrJSON(c *gin.Context, bindTo interface{}) (err error) {
	content_type := c.Request.Header.Get("Content-Type")
	content_type = strings.ToLower(content_type)
	method := c.Request.Method

	if method == "POST" || method == "PUT" {
		if strings.HasPrefix(content_type, "application/json") {
			if err = c.ShouldBindJSON(bindTo); err != nil {
				dbgo.Printf("%(red)In BindFormOrJSON at:%(LF) err=%s\n", err)
				c.JSON(http.StatusNotAcceptable, gin.H{ // 406
					"status": "error",
					"msg":    fmt.Sprintf("Error: %s", err),
				})
				return
			}
		} else {
			if err = c.ShouldBind(bindTo); err != nil {
				dbgo.Printf("%(red)In BindFormOrJSON at:%(LF) err=%s\n", err)
				c.JSON(http.StatusNotAcceptable, gin.H{ // 406
					"status": "error",
					"msg":    fmt.Sprintf("Error: %s", err),
				})
				return
			}
		}
	} else {
		if err = c.ShouldBind(bindTo); err != nil {
			dbgo.Printf("%(red)In BindFormOrJSON at:%(LF) err=%s\n", err)
			c.JSON(http.StatusNotAcceptable, gin.H{ // 406
				"status": "error",
				"msg":    fmt.Sprintf("Error: %s", err),
			})
			return
		}
	}
	dbgo.Printf("%(cyan)Parameters: %s at %s\n", dbgo.SVarI(bindTo), dbgo.LF(2))
	return
}

const db8 = true
