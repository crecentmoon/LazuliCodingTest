package repository

import (
	"context"
	"fmt"
)

type ProductRepository struct {
	SqlHandler
}

func NewProductRepository(sqlHandler SqlHandler) *ProductRepository {
	return &ProductRepository{sqlHandler}
}

func (db *ProductRepository) CreateProduct(ctx context.Context, p *ProductModel) error {
	p.Maker.ID = 1
	p.Brand.ID = 1

	if p.Maker.Name != "" && len(p.Maker.Name) > 0 {
		makerQuery := "INSERT INTO Makers (maker_name) VALUES (?)"
		makerID, err := db.Execute(ctx, makerQuery, p.Maker.Name)
		if err != nil {
			return fmt.Errorf("failed to insert maker: %v", err)
		}
		p.Maker.ID = makerID
	}

	if p.Brand.Name != "" && len(p.Brand.Name) > 0 {
		brandQuery := "INSERT INTO Brands (brand_name) VALUES (?)"
		brandID, err := db.Execute(ctx, brandQuery, p.Brand.Name)
		if err != nil {
			return fmt.Errorf("failed to insert brand: %v", err)
		}
		p.Brand.ID = brandID
	}

	productQuery := "INSERT INTO Products (jan, product_name, maker_id, brand_id) VALUES (?, ?, ?, ?)"
	_, err := db.Execute(ctx, productQuery, p.JAN, p.ProductName, p.Maker.ID, p.Brand.ID)
	if err != nil {
		return fmt.Errorf("failed to insert product: %v", err)
	}

	return nil
}

func (db *ProductRepository) CreateAttribute(ctx context.Context, p *ProductModel) error {
	query := "INSERT INTO Attributes (jan, attribute_data) VALUES (?, ?)"
	_, err := db.Execute(ctx, query, p.JAN, p.Attribute.Value)
	if err != nil {
		return fmt.Errorf("failed to insert attribute: %v", err)
	}

	return nil
}

func (db *ProductRepository) CreateDescriptionTags(ctx context.Context, p *ProductModel) error {
	for _, tag := range p.DescriptionTags {
		query := "INSERT INTO DescriptionTags (jan, tag_from_description) VALUES (?, ?)"
		_, err := db.Execute(ctx, query, p.JAN, tag.Tag)
		if err != nil {
			return fmt.Errorf("failed to insert description tag: %v", err)
		}
	}

	return nil
}

func (db *ProductRepository) CreateReviewTags(ctx context.Context, p *ProductModel) error {
	for _, tag := range p.ReviewTags {
		query := "INSERT INTO ReviewTags (jan, tag_from_review) VALUES (?, ?)"
		_, err := db.Execute(ctx, query, p.JAN, tag.Tag)
		if err != nil {
			return fmt.Errorf("failed to insert review tag: %v", err)
		}
	}

	return nil
}

// func (db *ProductRepository) GetProductByJan(JAN interface{}) (*entity.TrnRecipe, error) {
// 	r := entity.TrnRecipe{}

// 	getRecipe := "SELECT * FROM trn_recipes WHERE id = ?"
// 	if err := db.Query(&r, getRecipe, recipeId); err != nil {
// 		return nil, err
// 	}

// 	getIngredients := "SELECT * FROM trn_ingredients WHERE trn_recipe_id = ?"
// 	if err := db.Query(&r.TrnIngredients, getIngredients, recipeId); err != nil {
// 		return nil, err
// 	}

// 	getCookProcess := "SELECT * FROM trn_cook_processes WHERE trn_recipe_id = ?"
// 	if err := db.Query(&r.TrnCookProcesses, getCookProcess, recipeId); err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }
