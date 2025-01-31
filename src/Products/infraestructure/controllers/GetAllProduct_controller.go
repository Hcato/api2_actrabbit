package controllers

import (
	"api/src/Products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProductController struct {
	GetAllUseCase *application.GetAllProduct
}

func NewGetAllProductController(GetAllUseCase application.GetAllProduct) *GetAllProductController {
	return &GetAllProductController{GetAllUseCase: &GetAllUseCase}
}

func (getallproduct *GetAllProductController) Execute(ctx *gin.Context) {
	products, err := getallproduct.GetAllUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"Productos": products})
}
