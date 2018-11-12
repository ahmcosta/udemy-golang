package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Usuario :)
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/") // This line trims all content before /usuarios/. eg: http://localhost:3000/v3/usuarios/1795 will return just 1795
	id, _ := strconv.Atoi(sid)                          // It converts string to int. Ignore if error

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET": // if id is nil
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:34395198@/cursogo")
	if err != nil {
		log.Fatal(err) // TODO implement http 500 error
	}
	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome) // It's different from db.Query which returns a result set with multiple lines. db.QueryRow returns only one line

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json") // set the content type
	fmt.Fprint(w, string(json))                        // write on ResponseWriter
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:34395198@/cursogo")
	if err != nil {
		log.Fatal(err) // TODO implement http 500 error
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	var usuarios []Usuario // usuarios slice
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
