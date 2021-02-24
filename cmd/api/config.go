package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	BindAddr string `default:"0.0.0.0:7788" split_words:"true"`
	//db
	MysqlUser     string `split_words:"true" default:"user"`
	MysqlPassword string `split_words:"true" default:"pwd"`
	MysqlHost     string `split_words:"true" default:"host"`
	MysqlSchema   string `split_words:"true" default:"schema"`

	//redis
	RedisAddrs    []string `split_words:"true" default:"host:3306"`
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
