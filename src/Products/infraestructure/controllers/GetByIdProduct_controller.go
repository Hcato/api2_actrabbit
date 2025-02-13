package controllers

import (
	"api/src/Products/application"
	"api/src/Products/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GetByIdProductController struct {
	useCaseGetById *application.GetByIdProduct
}

func NewGetByIdProductController(useCaseGetById *application.GetByIdProduct) *GetByIdProductController {
	return &GetByIdProductController{useCaseGetById: useCaseGetById}
}

func (getByIdUser *GetByIdProductController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}

	//long polling
	initialProduct, err := getByIdUser.useCaseGetById.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateChan := make(chan domain.Product)
	timeout := time.After(1 * time.Second)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			updatedUser, err := getByIdUser.useCaseGetById.Execute(int32(id))
			if err != nil {
				continue
			}
			if updatedUser != initialProduct {
				updateChan <- updatedUser
				return
			}
		}
	}()

	select {
	case updatedProduct := <-updateChan:
		ctx.JSON(http.StatusOK, gin.H{"product": updatedProduct})
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"product": initialProduct})
	}
}
