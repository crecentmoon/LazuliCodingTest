package lazuli

import (
	"github.com/crecentmoon/lazuli-coding-test/internal/app/repository"
)

func PopulateTestData(db repository.SqlHandler) {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productUsecase := usecase.NewProductUsecase(productService)
	productHandler := handler.NewProductHandler(productUsecase)

}
