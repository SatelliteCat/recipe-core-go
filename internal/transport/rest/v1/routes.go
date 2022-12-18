package v1

import (
	"recipe.core/internal/config"
	"recipe.core/internal/transport/rest/v1/controllers"
)

func SetupRoutes(config *config.Config) {
	recipesController := controllers.Recipes{
		Storage: config.Store,
	}

	superGroup := config.App.Group("/api/v1")
	{
		recipesGroup := superGroup.Group("/recipe")
		{
			recipesGroup.Get("/:id<int>", recipesController.GetOneRecipe)
			recipesGroup.Post("/", recipesController.CreateOneRecipe)
		}
	}
}
