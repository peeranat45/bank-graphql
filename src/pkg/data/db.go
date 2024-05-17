package data

import (
	"database/sql"
	"log"
	"os"

	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var once sync.Once
var instance *DB

type DB struct {
	Connection *sql.DB
}

func NewDB() *DB {
	once.Do(func () {
		db, err := sql.Open("sqlite3", os.Getenv("DB"))
		if err != nil {
		log.Fatal(err)
		}
		instance = &DB{Connection: db}
	})
	
	return instance
}



