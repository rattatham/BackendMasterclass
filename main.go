package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/techschool/api"
	db "github.com/techschool/db/sqlc"
	"github.com/techschool/util"
)

func main() {
	config, err := util.LoadConfig(".") //path
	if err != nil {
		log.Fatal("cannot load config file", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
