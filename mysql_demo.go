package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/kin_demo?charset=utf8")
	if err != nil {
		//panic(err)
		fmt.Printf("%v\n", err)
		return
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		//panic(err)
		fmt.Printf("%v\n", err)
		return
	}

	result, err := db.Query("SELECT * FROM person")
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		for result.Next() {
			var id int
			var age int
			var name string
			result.Scan(&id, &age, &name)
			fmt.Printf("%v-%v-%v\n", id, age, name)
		}
		result.Close()
	}
}
