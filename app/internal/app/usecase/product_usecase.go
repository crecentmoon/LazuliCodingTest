package usecase

import (
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

func (p *ProductUseCase) PopulateProductData() error {

	jsonlPath := "tests/testdata/"

	files, err := filepath.Glob(jsonlPath + "*.jsonl")
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := p.productService.ImportProductDataFromJsonlFile(file); err != nil {
			return err
		}
	}

	return nil
}

// func (p *ProductUseCase) SearchProductByJan(jan string) (*domain.Product, error) {
// 	return p.productService.SearchProductByJan(jan)
// }

// func (p *ProductUseCase) CalculateProductAdequacyRate(product *domain.Product) (float64, error) {
// 	return p.productService.CalculateProductAdequacyRate(product)
// }
