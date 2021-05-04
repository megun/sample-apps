package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func checkDb() []string {
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dbconf := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname
	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, _ := db.Query("SHOW DATABASES")
	var databases []string
	var database string
	for res.Next() {
		res.Scan(&database)
		databases = append(databases, database)
	}
	return databases
}

func homePage(w http.ResponseWriter, r *http.Request) {
	databases := checkDb()
	fmt.Fprintf(w, strings.Join(databases, ","))
	fmt.Println(databases)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
