package main

import (
	"database/sql"
	"log"

	"github.com/iam-benjamen/simple_bank/api"
	db "github.com/iam-benjamen/simple_bank/db/sqlc"
	"github.com/iam-benjamen/simple_bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
