package main

import (
	"github.com/gohouse/gorose"
	"fmt"
)

var dbConfig = map[string]map[string]string {
	"mysql": {
		"host":     "localhost",
		"username": "root",
		"password": "",
		"port":     "3306",
		"database": "test",
		"charset":  "utf8",
		"protocol": "tcp",
	},
	"mysql_dev": {
		"host":     "localhost",
		"username": "root",
		"password": "",
		"port":     "3306",
		"database": "gorose",
		"charset":  "utf8",
		"protocol": "tcp",
	},
}

func main() {
	gorose.Open(dbConfig, "mysql")

	var db gorose.Database
	//
	res := db.Table("users").First()
	//
	fmt.Println(res)
}