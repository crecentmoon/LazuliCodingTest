package rdb

import (
	"context"
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

func (db *ProductsRepository) GetRecipeById(recipeId interface{}) (*entity.TrnRecipe, error) {
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

func (db *ProductsRepository) StoreProducts(ctx context.Context, p domain.Product) {
	makerQuery := "INSERT INTO Makers (maker_name) VALUES (?)"
	_, err := db.Execute(ctx, makerQuery, p.Maker)
}
