package usecase

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/crecentmoon/lazuli-coding-test/internal/infra"
)

type Products struct {
	JAN                 string       `json:"jan"`
	ProductName         string       `json:"product_name"`
	Attributes          []Attributes `json:"attributes"`
	Maker               string       `json:"maker"`
	Brand               string       `json:"brand"`
	TagsFromDescription []string     `json:"tags_from_description"`
	TagsFromReview      []string     `json:"tags_from_review"`
}

type Attributes struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

func PopulateProductRelatedData() {
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
		err = importDataFromFile(file, db)
		if err != nil {
			log.Println(err)
		}
	}

}

func importDataFromFile(file string, db infra.SqlInterface) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()

		var product Products
		err := json.Unmarshal(line, &product)
		if err != nil {
			log.Println(err)
			continue
		}

		err = saveProduct(product, db)
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

func saveProduct(product Products, db infra.SqlInterface) error {
	makerQuery := "INSERT INTO Makers (maker_name) VALUES (?)"
	_, err := db.Execute(makerQuery, product.Maker)
	if err != nil {
		return fmt.Errorf("failed to insert maker: %v", err)
	}

	brandQuery := "INSERT INTO Brands (brand_name) VALUES (?)"
	_, err = db.Exec(brandQuery, product.Brand)
	if err != nil {
		return fmt.Errorf("failed to insert brand: %v", err)
	}

	productQuery := "INSERT INTO Products (jan, product_name, maker_id, brand_id) VALUES (?, ?, LAST_INSERT_ID(), LAST_INSERT_ID())"
	_, err = db.Exec(productQuery, product.JAN, product.ProductName)
	if err != nil {
		return fmt.Errorf("failed to insert product: %v", err)
	}

	for _, attr := range product.Attributes {
		attrQuery := "INSERT INTO Attributes (jan, attribute_data) VALUES (?, ?)"
		_, err = db.Exec(attrQuery, product.JAN, attr)
		if err != nil {
			return fmt.Errorf("failed to insert attribute: %v", err)
		}
	}

	for _, tag := range product.TagsFromDescription {
		tagQuery := "INSERT INTO DescriptionTags (jan, tag_from_description) VALUES (?, ?)"
		_, err = db.Exec(tagQuery, product.JAN, tag)
		if err != nil {
			return fmt.Errorf("failed to insert description tag: %v", err)
		}
	}

	for _, tag := range product.TagsFromReview {
		tagQuery := "INSERT INTO ReviewTags (jan, tag_from_review) VALUES (?, ?)"
		_, err = db.Exec(tagQuery, product.JAN, tag)
		if err != nil {
			return fmt.Errorf("failed to insert review tag: %v", err)
		}
	}

	return nil
}
