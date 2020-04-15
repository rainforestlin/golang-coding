package main

import (
	"github.com/go-ini/ini"
	"log"
)

var Cfg *ini.File

type RabbitMQInfo struct {
	Url      string
	Port     string
	User     string
	Password string
}

var RabbitMQ *RabbitMQInfo

func init() {
	var err error
	Cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Fatalf("failed to parse 'config.ini: %v'", err)
	}
	loadRabbitMQ()

}

func loadRabbitMQ() {
	sec, err := Cfg.GetSection("RabbitMQ")
	if err != nil {
		log.Fatalf("failed to parse 'RabbitMQ: %v'", err)
	}
	RabbitMQ = &RabbitMQInfo{
		Url:      sec.Key("RabbitMQ_URL").String(),
		Port:     sec.Key("RabbitMQ_Port").MustString("5672"),
		User:     sec.Key("RabbitMQ_User").String(),
		Password: sec.Key("RabbitMQ_Pass").String(),
	}
}
