package v1

import (
	"github.com/gin-gonic/gin"
	"recipe.core/internal/transport/rest/v1/controllers"
)

func SetupRoutes(router *gin.Engine) {
	recipesRoutes := controllers.Recipes{}

	superGroup := router.Group("/api/v1")
	{
		recipesGroup := superGroup.Group("/recipe")
		{
			recipesGroup.GET("/:id", recipesRoutes.GetOneRecipe)
		}
	}
}
