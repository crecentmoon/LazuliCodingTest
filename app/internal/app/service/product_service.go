package service

import (
	"bufio"
	"context"
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

func (r *ProductService) ImportProductDataFromJsonlFile(file string) error {
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

		productModel := r.mapToProductModel(&product)
		err = r.productRepo.CreateProduct(context.Background(), &productModel)
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

func (r *ProductService) mapToProductModel(p *domain.Product) repository.ProductModel {
	productModel := &repository.ProductModel{
		JAN:             convertJAN(p.JAN),
		ProductName:     p.ProductName,
		Maker:           r.mapToMakerModel(p),
		Brand:           r.mapToBrandModel(p),
		Attribute:       r.mapToAttributeModel(p),
		DescriptionTags: r.mapToDescriptionTagsModel(p),
		ReviewTags:      r.mapToReviewTagsModel(p),
	}

	return *productModel
}

func (r *ProductService) mapToMakerModel(p *domain.Product) repository.MakerModel {
	makerModel := &repository.MakerModel{
		Name: p.Maker,
	}

	return *makerModel
}

func (r *ProductService) mapToBrandModel(p *domain.Product) repository.BrandModel {
	brandModel := &repository.BrandModel{
		Name: p.Brand,
	}

	return *brandModel
}

func (r *ProductService) mapToAttributeModel(p *domain.Product) repository.AttributeModel {
	attributeModel := &repository.AttributeModel{
		JAN:   convertJAN(p.JAN),
		Value: p.Attributes,
	}

	return *attributeModel
}

func (r *ProductService) mapToDescriptionTagsModel(p *domain.Product) []repository.DescriptionTagModel {
	var descriptionTagsModel []repository.DescriptionTagModel

	for _, tag := range p.TagsFromDescription {
		descriptionTagModel := &repository.DescriptionTagModel{
			JAN: convertJAN(p.JAN),
			Tag: tag,
		}

		descriptionTagsModel = append(descriptionTagsModel, *descriptionTagModel)
	}

	return descriptionTagsModel
}

func (r *ProductService) mapToReviewTagsModel(p *domain.Product) []repository.ReviewTagModel {
	var reviewTagsModel []repository.ReviewTagModel

	for _, tag := range p.TagsFromReview {
		reviewTagModel := &repository.ReviewTagModel{
			JAN: convertJAN(p.JAN),
			Tag: tag,
		}

		reviewTagsModel = append(reviewTagsModel, *reviewTagModel)
	}

	return reviewTagsModel
}

func convertJAN(jan string) int64 {
	janInt, _ := strconv.ParseInt(jan, 10, 64)

	return janInt
}
