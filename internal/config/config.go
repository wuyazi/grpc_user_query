package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DataBase DBConf `json:"DataBase"`
}

// DBConf defines db config
type DBConf struct {
	Driver   string `json:"Driver"`
	Host     string `json:"Host"`
	Port     int64  `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
	Db       string `json:"Db"`
	Schema   string `json:"Schema"`
}
