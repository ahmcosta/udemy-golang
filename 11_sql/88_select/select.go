package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:34395198@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 4004)
	defer rows.Close() // It should be closed to close the result set

	for rows.Next() { // indeterminated loop
		var u usuario
		rows.Scan(&u.id, &u.nome) // It scans the rs and then maps to a struct
		fmt.Println(u)
	}
}
