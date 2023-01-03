package main

import (
	"github.com/milankyncl/feature-toggles/internal/http"
	"github.com/milankyncl/feature-toggles/internal/storage/adapters"
	"log"
	"os"
)

func main() {
	dbpath := os.Getenv("ZAPPLES_DB_PATH")
	if dbpath == "" {
		log.Fatal("dbpath is not defined")
	}
	migrationsPath := os.Getenv("ZAPPLES_MIGRATIONS_PATH")
	if dbpath == "" {
		log.Fatal("migrationsPath is not defined")
	}
	uiPath := os.Getenv("ZAPPLES_UI_PATH")
	if dbpath == "" {
		log.Fatal("migrationsDir is not defined")
	}

	err, sqlite := adapters.NewSQLite(dbpath, migrationsPath)
	if err != nil {
		panic(err)
	}
	defer sqlite.Close()

	server := http.NewServer(uiPath, sqlite)

	log.Println("Starting webserver on 0.0.0.0:3000")
	_ = server.ListenAndServe(3000)
}
