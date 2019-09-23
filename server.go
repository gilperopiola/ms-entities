package main

import (
	"flag"
	"log"
	"os"

	"github.com/gilperopiola/ms-entities/config"
	"github.com/gilperopiola/ms-entities/database"
)

var cfg config.MyConfig
var db database.MyDatabase
var rtr MyRouter

func main() {
	env := flag.String("env", "local", "local / dev / prod")
	flag.Parse()

	cfg.Setup(*env)
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()

	log.Println("server started")
	rtr.Run(":" + os.Getenv("PORT"))
}
