package lazuli

import (
	"github.com/crecentmoon/lazuli-coding-test/internal/app/repository"
	"github.com/crecentmoon/lazuli-coding-test/internal/app/service"
	"github.com/crecentmoon/lazuli-coding-test/internal/app/usecase"
)

func PopulateTestData(db repository.SqlHandler) {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(*productRepository)
	productUsecase := usecase.NewProductUseCase(productService)

	productUsecase.PopulateProductData()
}
