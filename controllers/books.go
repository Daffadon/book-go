package controllers

import (
	"book/db"
	"book/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct{}

var BookModel = new(models.Book)

func (b BookController) GetBooks(ctx *gin.Context) {
	var books []models.Book
	db.GetDB().Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func (b BookController) CreateBook(ctx *gin.Context) {
	var input models.CreateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := BookModel.CreateBook(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create Book!"})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": book})
	ctx.Abort()
}

func (b BookController) FindBook(ctx *gin.Context) {
	if ctx.Param("id") != "" {
		book, err := BookModel.GetBookByID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": book})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"massage": "bad request!"})
	ctx.Abort()
}

func (b BookController) UpdateBook(ctx *gin.Context) {
	if ctx.Param("id") != "" {
		var input models.UpdateBookInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		book, err := BookModel.UpdateBook(ctx.Param("id"), &input)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": book})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"massage": "bad request!"})
	ctx.Abort()
}

func (b BookController) DeleteBook(ctx *gin.Context) {
	if ctx.Param("id") != "" {
		deleted, err := BookModel.DeleteBook(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Book Failed to Delete!"})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": deleted})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"massage": "bad request!"})
	ctx.Abort()
}
