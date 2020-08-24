package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func init() {
}

func main() {
	DbConnection, err := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}

	cmd := `CREATE TABLE IF NOT EXISTS person(
						name STRING,
						age INT)`
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = DbConnection.Exec(cmd, "nancy", 20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 30, "nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// multiple select
	// cmd = "SELECT * FROM person"
	// rows, _ := DbConnection.Query(cmd)
	// defer rows.Close()
	//
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	// 一つ一つエラーをck
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// // 上記エラーをまとめてck
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// single select
	// cmd = "SELECT * FROM person WHERE age <= ?"
	// row := DbConnection.QueryRow(cmd, 100)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// delete
	// cmd = "DELETE FROM person WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, "nancy")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }

	// table nameに対しては?じゃなく%sを使う
	// tableName := "person; INSERT INTO person (name, age) VALUES ('Mr.X', 100);" // sql injection
	tableName := "person"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()

	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		// 一つ一つエラーをck
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	// 上記エラーをまとめてck
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
