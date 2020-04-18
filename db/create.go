package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//_ import driver
	_ "github.com/lib/pq"
)

func Create() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	defer db.Close()
}
