package main

import (
	"fmt"
	"log"
	"net/http"
	"testManabie/cache"
	"testManabie/db"
)

const (
	webPort = "80"
)

type Config struct {
	Db      *db.Db
	CacheDb *cache.CacheDb
}

var app *Config

func main() {
	//connect db
	db := db.New()

	//connect cache
	cachedb := cache.New()

	app = &Config{
		Db:      db,
		CacheDb: cachedb,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Printf("server on port %s", webPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
