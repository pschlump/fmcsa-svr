package config

// Copyright (c) Philip Schlump, 2023.
// This file is MIT licensed, see ../LICENSE.mit

import (
	"fmt"
	"os"
	"sync"

	"github.com/pschlump/ReadConfig"
)

type ConfigData struct {
	// do not change - do not edit next line.
	Status string `json:"status" default:"success"`

	StatEngine string `json:"stat_engine" default:"memory"` // memory, file, redis

	StatFileLocaiton string `json:"stat_file_location" default:"./data/stat.json"`

	// Redis Connection Info
	RedisConnectHost string `json:"redis_host" default:"$ENV$REDIS_HOST"`
	RedisConnectAuth string `json:"redis_auth" default:"$ENV$REDIS_AUTH"`
	RedisConnectPort string `json:"redis_port" default:"6379"`
	RedisCluster     string `json:"redis_cluster" default:"no"`
}

var doOnce sync.Once
var gCfg *ConfigData

func LoadTestConfig() (rv *ConfigData) {
	doOnce.Do(func() {
		rv = &ConfigData{}
		gCfg = rv

		err := ReadConfig.ReadFile("./testdata/cfg.json", rv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read confguration: %s error %s\n", "./destdata/cfg.json", err)
			os.Exit(1)
		}

	})
	return gCfg
}
