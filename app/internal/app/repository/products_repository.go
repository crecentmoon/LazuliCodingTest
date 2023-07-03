package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/crecentmoon/lazuli-coding-test/internal/domain"
)

type ProductsRepository struct {
	SqlHandler
}

func (db *ProductsRepository) Store(ctx context.Context, r entity.TrnRecipe) (uint, error) {
	sql := "INSERT INTO trn_recipes(trn_user_id,title,thumbnail,movie,description,citation_flag,view_count,mst_total_time_id,mst_people_number_id,mst_country_id,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	now := time.Now()
	id, err := db.Execute(ctx, sql, r.TrnUserID, r.Title, r.Thumbnail, r.Movie, r.Description, r.CitationFlag, r.ViewCount, r.MstTotalTimeID, r.MstPeopleNumberID, r.MstCountryID, now, now)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *ProductsRepository) GetProductByJan(JAN interface{}) (*entity.TrnRecipe, error) {
	r := entity.TrnRecipe{}

	getRecipe := "SELECT * FROM trn_recipes WHERE id = ?"
	if err := db.Query(&r, getRecipe, recipeId); err != nil {
		return nil, err
	}

	getIngredients := "SELECT * FROM trn_ingredients WHERE trn_recipe_id = ?"
	if err := db.Query(&r.TrnIngredients, getIngredients, recipeId); err != nil {
		return nil, err
	}

	getCookProcess := "SELECT * FROM trn_cook_processes WHERE trn_recipe_id = ?"
	if err := db.Query(&r.TrnCookProcesses, getCookProcess, recipeId); err != nil {
		return nil, err
	}

	return &r, nil
}

func (db *ProductsRepository) StoreProductRelatedData(ctx context.Context, p domain.Product) error {
	makerQuery := "INSERT INTO Makers (maker_name) VALUES (?)"
	makerID, err := db.Execute(ctx, makerQuery, p.Maker)
	if err != nil {
		return fmt.Errorf("failed to insert maker: %v", err)
	}

	brandQuery := "INSERT INTO Brands (brand_name) VALUES (?)"
	brandID, err := db.Execute(ctx, brandQuery, p.Brand)
	if err != nil {
		return fmt.Errorf("failed to insert brand: %v", err)
	}

	productQuery := "INSERT INTO Products (jan, product_name, maker_id, brand_id) VALUES (?, ?, LAST_INSERT_ID(), LAST_INSERT_ID())"
	_, err = db.Execute(ctx, productQuery, p.JAN, p.ProductName, makerID, brandID)
	if err != nil {
		return fmt.Errorf("failed to insert product: %v", err)
	}

	attrQuery := "INSERT INTO Attributes (jan, attribute_data) VALUES (?, ?)"
	_, err = db.Execute(ctx, attrQuery, p.JAN, p.Attributes)
	if err != nil {
		return fmt.Errorf("failed to insert attribute: %v", err)
	}

	for _, tag := range p.TagsFromDescription {
		tagQuery := "INSERT INTO DescriptionTags (jan, tag_from_description) VALUES (?, ?)"
		_, err = db.Execute(ctx, tagQuery, p.JAN, tag)
		if err != nil {
			return fmt.Errorf("failed to insert description tag: %v", err)
		}
	}

	for _, tag := range p.TagsFromReview {
		tagQuery := "INSERT INTO ReviewTags (jan, tag_from_review) VALUES (?, ?)"
		_, err = db.Execute(ctx, tagQuery, p.JAN, tag)
		if err != nil {
			return fmt.Errorf("failed to insert review tag: %v", err)
		}
	}
}
