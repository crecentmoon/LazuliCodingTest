package lazuli

import (
	"github.com/crecentmoon/lazuli-coding-test/internal/app/repository"
)

func PopulateTestData(db repository.SqlHandler) {
	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productService)

}
