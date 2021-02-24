package main

import (
	"flag"
	"go-dev-web/pkg/plog"
	//"go-dev-web/pkg/redismodels"

	"go-dev-web/pkg/api"

	"go-dev-web/pkg/models"
)

func main() {
	var err error
	flag.Parse()

	if err = ParseEnvConfig(); err != nil {
		panic(err)
	}

	if err = plog.Init(cfg.LogFilePath, cfg.EnvType); err != nil {
		panic(err)
	}

	if err = models.InitMysqlDB("default", &models.MysqlDBWrap{
		Host:     cfg.MysqlHost,
		Username: cfg.MysqlUser,
		Password: cfg.MysqlPassword,
		Schema:   cfg.MysqlSchema,
	}); err != nil {
		plog.Panic(err)
	}

	//if err = redismodels.InitRedisDB("default", &redismodels.RedisWrap{
	//	Addrs:    cfg.RedisAddrs,
	//	Password: cfg.RedisPassword,
	//}); err != nil {
	//	plog.Panic(err)
	//}

	plog.Info("start api server")

	api.ServeAPI(&api.Config{
		BindAddr: cfg.BindAddr,
	})
}
