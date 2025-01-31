package controllers

import (
	"api/src/Products/application"
	"api/src/Products/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	useCaseUpdate *application.UpdateProduct
}

func NewUpdateProductController(useCaseUpdate *application.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{useCaseUpdate: useCaseUpdate}
}

func (updateProduct *UpdateProductController) Execute(ctx *gin.Context) {
	id := ctx.Param("id")

	var product domain.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	err = updateProduct.useCaseUpdate.Execute(int32(productId), product.Name, product.Price)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message:": "Product actualizado"})
}
