package main

import (
	"log"
	"os"

	"github.com/crecentmoon/lazuli-coding-test/cmd/lazuli"
	"github.com/crecentmoon/lazuli-coding-test/internal/infra"
	"github.com/crecentmoon/lazuli-coding-test/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("\033[1;31mERROR: Failed to load config\033[0m")
	} else {
		log.Println("\033[1;32mSUCCESS: Loaded config\033[0m")
	}

	db, err := infra.NewMySQLHandler(cfg)
	if err != nil {
		log.Fatalln("\033[1;31mERROR: Failed to connect to database\033[0m")
	} else {
		log.Println("\033[1;32mSUCCESS: Connected to database\033[0m")
	}

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "populate":
			lazuli.PopulateTestData(db)
		default:
			log.Println("\033[1;33mNOTICE: Command is invalid. \n Available commands are: populate\033[0m")
		}
	} else {
		lazuli.InitServer(db)
	}
}
