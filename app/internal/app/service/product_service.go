package service

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/crecentmoon/lazuli-coding-test/internal/app/repository"
	"github.com/crecentmoon/lazuli-coding-test/internal/domain"
)

type ProductService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (r *ProductService) ImportProductDataFromJsonl(file string) error {
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

		err = repository.StoreProducts(product, db)
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

func (r *ProductService) convertToProductModel(p *domain.Product) (*repository.ProductModel, error) {
	productModel := &repository.ProductModel{
		JAN:             strconv.Atoi(p.JAN),
		ProductName:     p.ProductName,
		Maker:           p.Maker,
		Brand:           p.Brand,
		Attribute:       p.Attributes,
		DescriptionTags: p.DescriptionTags,
		ReviewTags:      p.ReviewTags,
	}
	return product, nil
}
