package middleware

import (
	"bookApp/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
		ctx.Abort()
		return
	}
	ctx.Set("userID", strconv.Itoa(int(claims.UserId)))
	ctx.Next()
}
