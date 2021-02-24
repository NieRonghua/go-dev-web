package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	BindAddr string `default:"0.0.0.0:7788" split_words:"true"`
	//db
	MysqlUser     string `split_words:"true" default:"cdp"`
	MysqlPassword string `split_words:"true" default:"_Q4NVU@yhPUJ"`
	MysqlHost     string `split_words:"true" default:"cdpapib-m.dbsit.sfcloud.local:3306"`
	MysqlSchema   string `split_words:"true" default:"cdp"`

	//redis
	RedisAddrs    []string `split_words:"true" default:"Vp15A26i-1.cachesit.sfcloud.local:8080,Vp15A26i-2.cachesit.sfcloud.local:8080,Vp15A26i-3.cachesit.sfcloud.local:8080,Vp15A26i-4.cachesit.sfcloud.local:8080,Vp15A26i-5.cachesit.sfcloud.local:8080,Vp15A26i-6.cachesit.sfcloud.local:8080"`
	RedisPassword string   `split_words:"true" default:"UIbjKf6ZktsjU75xvaBE9Mel"`

	Debug   bool   `default:"false" split_words:"true"`
	EnvType string `default:"local" split_words:"true"`

	LogFilePath string `split_words:"true" default:"/go/release/logs"`
}

var cfg *Config

//
func ParseEnvConfig() error {
	cfg = new(Config)
	err := envconfig.Process("API", cfg)
	fmt.Printf("get config %+v\n", *cfg)
	return err
}
