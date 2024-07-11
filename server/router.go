package server

import (
	"book/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	health := new(controllers.HealthController)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	})
	router.GET("/health", health.Status)
	v1 := router.Group("/v1")
	{
		bookGroup := v1.Group("/book")
		{
			book := new(controllers.BookController)
			bookGroup.GET("/", book.GetBooks)
			bookGroup.POST("/", book.CreateBook)
			bookGroup.GET("/:id", book.FindBook)
			bookGroup.PATCH("/:id", book.UpdateBook)
			bookGroup.DELETE("/:id", book.DeleteBook)
		}
	}
	return router
}
