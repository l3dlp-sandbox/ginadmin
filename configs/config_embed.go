//+build embed

/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-29 19:47:42
 */

package configs

import (
	"bytes"
	_ "embed"
	"fmt"

	"gopkg.in/ini.v1"
)

type AppConf struct {
	MysqlConf   `ini:"mysql"`
	RedisConf   `ini:"redis"`
	SessionConf `ini:"session"`
	BaseConf    `ini:"base"`
}

type MysqlConf struct {
	Host        string `ini:"host"`
	Port        string `ini:"port"`
	UserName    string `ini:"username"`
	Password    string `ini:"password"`
	Database    string `ini:"database"`
	MaxOpenConn int    `ini:"max_open_conn"`
	MaxIdleConn int    `ini:"max_idle_conn"`
}

type RedisConf struct {
	Addr     string `ini:"addr"`
	Db       int    `ini:"db"`
	Password string `ini:"password"`
}

type SessionConf struct {
	SessionName string `ini:"session_name"`
}

type BaseConf struct {
	Port         string `ini:"port"`
	Host         string `ini:"host"`
	FillData     bool   `ini:"fill_data"`
	MigrateTable bool   `ini:"migrate_table"`
}

var App = new(AppConf)

//go:embed config.ini
var iniStr string

//初始化配置文件
func init() {
	err := ini.MapTo(App, bytes.NewReader([]byte(iniStr)))
	if err != nil {
		fmt.Printf("load ini err:%v", err)
	}
}
