package controllers

import (
	"api/src/Users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllUserController struct {
	GetAllUseCase *application.GetAllUser
}

func NewGetAllUserController(GetAllUseCase application.GetAllUser) *GetAllUserController {
	return &GetAllUserController{GetAllUseCase: &GetAllUseCase}
}

func (getalluser *GetAllUserController) Execute(ctx *gin.Context) {
	products, err := getalluser.GetAllUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"Users": products})
}
