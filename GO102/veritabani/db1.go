package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := sql.Open("sql-driver", "filmler.sqbpro")
	defer db.Close()
	checkErr(err)
	db.Exec("CREATE TABLE notlar(id INT PRIMARY KEY , notdurumu TEXT,isim TEXT)")

}
