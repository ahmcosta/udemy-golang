package main

import (
	"database/sql"

	/*
		"_" doesn’t require to use package math in importing file but init function(s)
		from imported package will be executed anyway (package and it dependencies
		will be initialized). It’s useful if we’re interested only in bootstrapping
		work done by imported package but we don’t reference any exported identifiers from it.

		https://medium.com/golangspec/import-declarations-in-go-8de0fd3ae8ff
	*/
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
