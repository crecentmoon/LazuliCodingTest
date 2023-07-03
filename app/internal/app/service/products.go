package service

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/crecentmoon/lazuli-coding-test/internal/domain"
	"github.com/crecentmoon/lazuli-coding-test/internal/infra"
)

func ImportRepositoryDataFromFile(file string, db infra.SqlInterface) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()

		var product domain.Product
		err := json.Unmarshal(line, &product)
		if err != nil {
			log.Println(err)
			continue
		}

		err = repository.StoreProductRelatedData(product, db)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
