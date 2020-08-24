package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DbConnection *sql.DB

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"string,omitempty"`
}

const (
	DATABASE_DRIVER = "mysql"
	DATABASE_CONFIG = "root:password@tcp(golang-mysql:3306)/mysql"
)

func init() {
	DbConnection, err := sql.Open(DATABASE_DRIVER, DATABASE_CONFIG)
	log.Println(DbConnection)
	log.Println("connected to mysql.")
	// 接続エラー時処理
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS person(
							name STRING,
							age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err.Error())
	}

	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err := DbConnection.Exec(cmd, "motty", 93)
	if err != nil {
		log.Fatalln(err.Error())
	}

	rows, err := DbConnection.Query("SELECT * FROM person")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(rows)
}
