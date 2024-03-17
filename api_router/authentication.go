package api_router

import (
	"api/helper"
	model "api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostUser             godoc
// @Summary      Register a new user
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         auth
// @Produce      json
// @Param        book  body      model.AuthenticationInput  true  "Book JSON"
// @Success      200   {object}  model.AuthenticationInput
// @Router       /auth/register [post]
func Register(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	findedUser, findedUserErr := model.FindUserByUsername(user.Username)
	fmt.Printf("%+v\n", findedUser)

	if findedUser.Username != "" {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "user already exist"})
	} else {
		savedUser, err := user.Save()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		context.JSON(http.StatusCreated, gin.H{"user": savedUser})
	}

	if findedUserErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": findedUserErr.Error()})
	}

}

func AddEntry(context *gin.Context) {
	var input model.Entry

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = user.ID
	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})

}

func Login(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
