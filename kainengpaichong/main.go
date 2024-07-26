package main

import (
	"kainengpaichong/model"
	"kainengpaichong/router"
	"kainengpaichong/until"
)

func main() {
	r := router.Router()
	model.Init()
	model.NewFromplan()
	cfg, err := until.ParseConfig("./config/config.json")
	if err != nil {
		panic(err)
	}
	r.Run(cfg.APPHOST + ":" + cfg.APPPORT)
}
