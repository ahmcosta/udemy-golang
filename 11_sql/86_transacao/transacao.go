package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:34395198@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin() // initialize the transaction. tx is the base object to work
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(4000, "Bia")
	stmt.Exec(4001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil { // It must be handled. GO doesn't throw a exception
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
