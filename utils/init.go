package utils

import (
	"database/sql"
)

var db *sql.DB

func Init() (err error) {
	// 数据库连接配置， 以Mysql为例
	dbn := "higginslee:79b75c8bedcad1c9@tcp(mysql.sqlpub.com:3306)/test_proxy_pool?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dbn)
	if err != nil {
		return nil
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil

}
