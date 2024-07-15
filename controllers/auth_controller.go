package controllers

import (
	"bookApp/models"
	"bookApp/utils"
	"bookApp/validations"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

var UserModel = new(models.User)

func (auth AuthController) Login(ctx *gin.Context) {
	var input validations.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	user, err := UserModel.GetUser(&input)
	if err == nil {
		errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
		if errCompare != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
			ctx.Abort()
			return
		}
		token, errJWT := utils.GenerateToken(user.ID, user.Role)
		if errJWT != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token!"})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"jwt": token})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
	ctx.Abort()
}

func (auth AuthController) Register(ctx *gin.Context) {
	var input validations.RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if input.Password != input.ConfirmPasword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password and Confirm Password not match!"})
		ctx.Abort()
		return
	}

	if hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), 14); errHash != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password!"})
		ctx.Abort()
		return
	} else {
		input.Password = string(hashedPassword)
	}

	createdUser := validations.CreateUserInput{
		Fullname: input.Fullname,
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}
	_, err := UserModel.CreateUser(&createdUser)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"massage": "success"})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	ctx.Abort()
}
