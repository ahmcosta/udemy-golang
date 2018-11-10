package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // implicit import
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	// mysql arguments uses the import at line 6
	db, err := sql.Open("mysql", "root:34395198@/") // root database. There isn't database yet. It will be created in the next lines
	if err != nil {
		panic(err)
	}
	defer db.Close() // postpone the execution

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	// backtic for multiple lines (`)
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
