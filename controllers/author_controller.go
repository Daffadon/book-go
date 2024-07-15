package controllers

import (
	"bookApp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorController struct{}

var AuthorModel = new(models.Author)

func (a AuthorController) GetAuthors(ctx *gin.Context) {
	fetchedAutors, err := AuthorModel.GetAuthors()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if len(fetchedAutors) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Authors Found!"})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": fetchedAutors})
	ctx.Abort()
}
