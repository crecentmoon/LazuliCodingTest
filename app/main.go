package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/crecentmoon/lazuli-coding-test/cmd/lazuli"
	"github.com/joho/godotenv"
)

func main() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprintf(os.Stderr, "ERROR: Unable to identify current directory (needed to load .env)")
		os.Exit(1)
	}
	basepath := filepath.Dir(file)
	log.Println(basepath)
	err := godotenv.Load(filepath.Join(basepath, ".env"))
	if err != nil {
		log.Fatal("ERROR: Failed to load .env file")
	}

	if os.Args[1] == "populate" {
		lazuli.PopulateTestData()
		log.Println("OK: Finished populating test data")
		os.Exit(1)
	}

	lazuli.InitServer()
}
