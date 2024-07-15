package controllers

import (
	"bookApp/models"
	"bookApp/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct{}

var BookModel = new(models.Book)

func (b BookController) GetBooks(ctx *gin.Context) {
	fetchedBooks, err := BookModel.GetBooks()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if len(fetchedBooks) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No Books Found!"})
		ctx.Abort()
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"data": fetchedBooks})
	ctx.Abort()
}

func (b BookController) CreateBook(ctx *gin.Context) {
	userRole, ok := ctx.Get("role")
	if ok {
		if userRole != "admin" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
			ctx.Abort()
			return
		}
		var input validations.CreateBookServiceInput
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
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
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
		userRole, ok := ctx.Get("role")
		if ok {
			if userRole != "admin" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
				ctx.Abort()
				return
			}
			var input validations.UpdateBookServiceInput
			if err := ctx.ShouldBindJSON(&input); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			book, err := BookModel.UpdateBook(ctx.Param("id"), &input)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"data": book})
			return
		}
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"massage": "bad request!"})
	ctx.Abort()
}

func (b BookController) DeleteBook(ctx *gin.Context) {
	if ctx.Param("id") != "" {
		userRole, ok := ctx.Get("role")
		if ok {
			if userRole != "admin" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
				ctx.Abort()
				return
			}
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
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"massage": "bad request!"})
	ctx.Abort()
}
