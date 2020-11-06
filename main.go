package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var db_user = "root"
var db_pass = "root"
var db_host = "localhost"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHello)
	r.HandleFunc("/db", indexDb)
	http.Handle("/", r)
	fmt.Println("Starting server ...")
	http.ListenAndServe(":8888", nil)
}

func indexHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func indexDb(w http.ResponseWriter, req *http.Request) {
	user, exists := os.LookupEnv("MYSQL_ROOT_USER")
	if exists {
		fmt.Println("Found MYSQL_ROOT_USER ", user)
		db_user = user
	}
	pass, exists := os.LookupEnv("MYSQL_ROOT_PASSWORD")
	if exists {
		fmt.Println("Found MYSQL_ROOT_PASSWORD ", pass)
		db_pass = pass
	}
	host, exists := os.LookupEnv("MYSQL_HOST")
	if exists {
		fmt.Println("Found MYSQL_HOST ", host)
		db_host = host
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", db_user, db_pass, db_host))
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Fprintln(w, "Database OK!")
	}
	defer db.Close()
}
