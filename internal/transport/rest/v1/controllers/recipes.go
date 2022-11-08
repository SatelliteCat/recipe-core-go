package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Recipes struct {
}

func (r *Recipes) GetOneRecipe(ctx *gin.Context) {
	recipeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": recipeId,
	})
}
