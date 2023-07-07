package lazuli

import (
	"fmt"
	"log"

	"github.com/crecentmoon/lazuli-coding-test/internal/app/repository"
	"github.com/crecentmoon/lazuli-coding-test/internal/app/service"
	"github.com/crecentmoon/lazuli-coding-test/internal/app/usecase"
)

func PopulateTestData(db repository.SqlHandler) {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(*productRepository)
	productUsecase := usecase.NewProductUseCase(productService)

	if err := productUsecase.PopulateProductData(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully populated test data")
}
