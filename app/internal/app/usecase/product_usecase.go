package usecase

import (
	"log"
	"path/filepath"

	"github.com/crecentmoon/lazuli-coding-test/internal/app/service"
)

type ProductUseCase struct {
	productService *service.ProductService
}

func NewProductUseCase(productService *service.ProductService) *ProductUseCase {
	return &ProductUseCase{
		productService: productService,
	}
}

func (p *ProductUseCase) PopulateProductData() {

	jsonlPath := "app/tests/testdata/"

	files, err := filepath.Glob(jsonlPath + "*.jsonl")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err = p.productService.ImportProductDataFromJsonl(file)
		if err != nil {
			log.Println(err)
		}
	}
}

// func (p *ProductUseCase) SearchProductByJan(jan string) (*domain.Product, error) {
// 	return p.productService.SearchProductByJan(jan)
// }

// func (p *ProductUseCase) CalculateProductAdequacyRate(product *domain.Product) (float64, error) {
// 	return p.productService.CalculateProductAdequacyRate(product)
// }
