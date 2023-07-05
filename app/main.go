package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crecentmoon/lazuli-coding-test/cmd/lazuli"
	"github.com/crecentmoon/lazuli-coding-test/internal/infra"
	"github.com/crecentmoon/lazuli-coding-test/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()

	db, err := infra.NewMySQLHandler(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "populate":
			lazuli.PopulateTestData(db)
		default:
			fmt.Println("Command is invalid. Available commands are: populate")
		}
	} else {
		lazuli.InitServer(db)
	}
}
