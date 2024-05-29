package main

import (
	"database/sql"
	"log"

	goora "github.com/sijms/go-ora/v2"
)

func main() {
	// Oracle 数据库连接信息
	dsn := goora.BuildUrl("10.251.16.185", 1521, "histdb", "UCR_CEN", "123abc", nil)

	// 创建数据库连接
	db, err := sql.Open("goora", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
