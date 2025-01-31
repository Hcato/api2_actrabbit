package controllers

import (
	"api/src/Products/application"
	"api/src/Products/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	CreateUseCase *application.CreateProductUsecase
}

func NewCreateProductController(CreateUseCase *application.CreateProductUsecase) *CreateProductController {
	return &CreateProductController{CreateUseCase: CreateUseCase}
}

func (createProduct *CreateProductController) Execute(c *gin.Context) {
	var product domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createProduct.CreateUseCase.Execute(product.Name, product.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, gin.H{"message": "Producto registrado"})
}
