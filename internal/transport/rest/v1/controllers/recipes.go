package controllers

import (
	"github.com/gofiber/fiber/v2"
	"recipe.core/internal/storage"
	"recipe.core/internal/storage/psql"
)

type Recipes struct {
	*storage.Storage
}

type responseError struct {
	Error string `json:"error"`
}

type responseRecipeOne struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Recipe      string `json:"recipe"`
	Ingredients string `json:"ingredients"`
}

func (r *Recipes) GetOneRecipe(ctx *fiber.Ctx) error {
	recipeId, err := ctx.ParamsInt("id")
	if err != nil {
		ctx.Response().SetStatusCode(400)

		return ctx.JSON(responseError{Error: "Invalid recipeModel ID"})
	}

	recipeRepository := psql.RecipeRepository{
		Pool: r.Storage.GetPgxConn(),
	}

	recipeModel, err := recipeRepository.FindRecipeOneById(recipeId)
	if err != nil {
		ctx.Response().SetStatusCode(404)

		return ctx.JSON(responseError{Error: "Recipe not found"})
	}

	return ctx.JSON(responseRecipeOne{
		Id:          recipeModel.Id,
		Name:        recipeModel.Name,
		Url:         recipeModel.Url,
		Recipe:      recipeModel.Recipe,
		Ingredients: recipeModel.Ingredients,
	})
}

func (r *Recipes) CreateOneRecipe(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "created",
	})
}
