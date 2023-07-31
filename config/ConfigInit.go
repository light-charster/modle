package config

import "gorm.io/gorm"

var MysqlConnect *gorm.DB

func Init() {
	load()
	MysqlConnect = MysqlConn()
}
