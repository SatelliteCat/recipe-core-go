package psql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"recipe.core/internal/model"
)

type RecipeRepository struct {
	*pgxpool.Pool
}

func (rr *RecipeRepository) FindRecipeOneById(recipeId int) (*model.Recipe, error) {
	var recipe model.Recipe

	row := rr.Pool.QueryRow(
		context.Background(),
		`
select id
     , name
     , url
     , recipe
     , ingredients
from recipe
where id = $1
`,
		recipeId,
	)

	err := row.Scan(&recipe.Id, &recipe.Name, &recipe.Url, &recipe.Recipe, &recipe.Ingredients)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}
