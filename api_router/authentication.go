package api_router

import (
	"api/helper"
	model "api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostUser             godoc
// @Summary      Register a new user
// @Description  Takes a user JSON and store in DB. Return saved JSON.
// @Tags         auth
// @Produce      json
// @Param        user  body      model.RegisterInput  true  "User JSON"
// @Success      201   {object}  model.UserResp
// @Router       /auth/register [post]
func Register(context *gin.Context) {
	var input model.RegisterInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	findedUser, findedUserErr := model.FindUserByUsername(user.Username)
	findedEmail, emailErr := model.FindUserByEmail(user.Email)

	if findedUser.Username != "" {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "user already exist"})
	} else if findedEmail.Email != "" {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "email already exist"})
	} else {
		savedUser, err := user.Save()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		respUser := model.UserResp{
			UserId:   int(savedUser.ID),
			UserName: savedUser.Username,
			Email:    savedUser.Email,
		}
		context.JSON(http.StatusCreated, respUser)
	}

	if findedUserErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": findedUserErr.Error()})
	}

	if emailErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": emailErr.Error()})
	}

}

// PostUser             godoc
// @Summary      Login
// @Description  Takes a user JSON and store in DB. Return saved JSON.
// @Tags         auth
// @Produce      json
// @Param        user  body      model.AuthenticationInput  true  "User JSON"
// @Success      200   {object}  model.JwtResp
// @Router       /auth/login [post]
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
