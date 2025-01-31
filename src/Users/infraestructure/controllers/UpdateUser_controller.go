package controllers

import (
	"api/src/Users/application"
	"api/src/Users/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	useCaseUpdate *application.UpdateProduct
}

func NewUpdateUserController(useCaseUpdate *application.UpdateProduct) *UpdateUserController {
	return &UpdateUserController{useCaseUpdate: useCaseUpdate}
}

func (updateUser *UpdateUserController) Execute(ctx *gin.Context) {
	id := ctx.Param("id")

	var User domain.User

	if err := ctx.ShouldBindJSON(&User); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	err = updateUser.useCaseUpdate.Execute(int32(UserId), User.Name, User.Lastname)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message:": "Usuario actualizado"})
}
