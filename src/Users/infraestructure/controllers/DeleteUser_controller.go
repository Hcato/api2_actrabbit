package controllers

import (
	"api/src/Users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	useCaseDelete *application.DeleteUserUsecase
}

func NewDeleteUserController(useCaseDelete *application.DeleteUserUsecase) *DeleteUserController {
	return &DeleteUserController{useCaseDelete: useCaseDelete}
}

func (deleteTicket *DeleteUserController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
		return
	}
	err = deleteTicket.useCaseDelete.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User Eliminado"})
}
