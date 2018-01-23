package main

import (
	"./config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

func main() {
	db := gorose.Open(config.DbConfig, "mysql_dev")
	// close DB
	defer db.Close()

	data := map[string]interface{}{
		"age":  17,
		"job":  "it3",
		"name": "fizz5",
	}
	where := map[string]interface{}{
		"id": 17,
	}

	res := db.Table("users").Data(data).Where(where).Update()
	fmt.Println(db.LastSql())
	fmt.Println(res)

}
