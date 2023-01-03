package main

import (
	"github.com/milankyncl/feature-toggles/internal/http"
	"github.com/milankyncl/feature-toggles/internal/storage/adapters"
	"log"
	"os"
)

var basePath string

func main() {
	basePath = os.Getenv("ZAPPLES_BASE_PATH")
	if basePath == "" {
		log.Println("zapples base path is not defined, using working directory")
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal("could not get current working directory")
		}
		basePath = wd
	}

	err, sqlite := adapters.NewSQLite(basePath)
	if err != nil {
		panic(err)
	}
	defer sqlite.Close()

	server := http.NewServer(basePath, sqlite)

	log.Println("Starting webserver on 0.0.0.0:3000")
	_ = server.ListenAndServe(3000)
}
