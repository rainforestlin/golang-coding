package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = ""
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "godemo"
)

func InitDB() *sql.DB {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ := sql.Open("mysql", path)
	//defer DB.Close()
	e := DB.Ping()
	if e != nil {
		panic(e)
	}
	return DB
}
func CreateTable(name string, db *sql.DB) {
	db.Begin()
	_, e := db.Query(strings.Join([]string{"create table if not exists ", name, " (id int, name varchar(50),age int)"}, ""))
	if e != nil {
		panic(e)
	}
}
func InsertToUser(db *sql.DB) {
	db.Begin()
	//_, e := db.Query()
	fmt.Printf("insert into `user`(`id`,`name`,`age`) Values ('%d','%s','%d')\n", 1, "lihua", 18)
	//values中需要带''，不然会被当做列名报错
	_,e := db.Exec(fmt.Sprintf("insert into `user`(`id`,`name`,`age`) Values ('%d','%s','%d')\n", 1, "lihua", 18))
	if e != nil {
		panic(e)
	}
}
func main() {

	db := InitDB()
	defer db.Close()
	CreateTable("user", db)
	InsertToUser(db)
}
