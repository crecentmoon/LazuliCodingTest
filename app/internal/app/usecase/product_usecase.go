package usecase

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/crecentmoon/lazuli-coding-test/internal/app/service"
	"github.com/crecentmoon/lazuli-coding-test/internal/domain"
)

type ProductUseCase struct {
	productService *service.ProductService
}

func NewProductUseCase(productService *service.ProductService) *ProductUseCase {
	return &ProductUseCase{
		productService: productService,
	}
}

func PopulateProductData() {
	db, err := infra.NewSqlHandler()
	if err != nil {
		log.Fatal(err)
	}

	jsonlPath := "app/tests/testdata/"

	files, err := filepath.Glob(jsonlPath + "*.jsonl")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err = importRepositoryDataFromFile(file, db)
		if err != nil {
			log.Println(err)
		}
	}

}

func importRepositoryDataFromFile(file string, db infra.SqlInterface) error {
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
