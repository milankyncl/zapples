package main

import (
	"github.com/milankyncl/feature-toggles/internal/http"
	"github.com/milankyncl/feature-toggles/internal/storage/adapters"
)

func main() {
	err, sqlite := adapters.NewSQLite()
	if err != nil {
		panic(err)
	}
	defer sqlite.Close()

	server := http.NewServer(sqlite)
	server.ListenAndServe(3000)
}
