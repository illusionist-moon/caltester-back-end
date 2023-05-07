package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	cfg *ini.File

	RunMode string

	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	User     string
	Password string
	Ip       string
	Port     string
	Database string
	DSN      string
)

func init() {
	var err error
	cfg, err = ini.Load("conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadBase()
	LoadServer()
	LoadMysql()
}

func LoadBase() {
	RunMode = cfg.Section("").Key("RUN_MODE").String()
}

func LoadMysql() {
	User = cfg.Section("mysql").Key("USER").String()
	Password = cfg.Section("mysql").Key("PASSWORD").String()
	Ip = cfg.Section("mysql").Key("IP").String()
	Port = cfg.Section("mysql").Key("PORT").String()
	Database = cfg.Section("mysql").Key("DATABASE").String()

	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		User, Password, Ip, Port, Database)
}

func LoadServer() {
	sec, err := cfg.GetSection("server")
	if err != nil {
		log.Fatalf("fail to get section 'server': %v", err)
	}
	HttpPort = sec.Key("HTTP_PORT").String()
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
