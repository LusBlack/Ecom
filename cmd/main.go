package main

import (
	"database/sql"
	"log"

	"github.com/LusBlack/santeria/cmd/api"
	"github.com/LusBlack/santeria/config"
	"github.com/LusBlack/santeria/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	//configure db storage
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	//establish conn with db specs
	initStorage(db)

	//start server
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
