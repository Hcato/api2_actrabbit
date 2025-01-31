package controllers

import (
	"api/src/Products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCaseDelete *application.DeleteProductUsecase
}

func NewDeleteProductController(useCaseDelete *application.DeleteProductUsecase) *DeleteProductController {
	return &DeleteProductController{useCaseDelete: useCaseDelete}
}

func (deleteTicket *DeleteProductController) Execute(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, gin.H{"message": "Producto Eliminado"})
}
