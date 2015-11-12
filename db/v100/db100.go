package db100

import (
	"log"
	"time"

	"github.com/rhinoman/couchdb-go"
)

var DB *couchdb.Database

func init() {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("127.0.0.1", 5984, timeout)
	if err != nil {
		log.Fatal(err)
	}
	//auth := couchdb.BasicAuth{Username: "user", Password: "password"}
	DB = conn.SelectDB("gjallarhorn", nil)
}
