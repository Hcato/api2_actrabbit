package controllers

import (
	"api/src/Users/application"
	"api/src/Users/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GetByIdUserController struct {
	useCaseGetById *application.GetByIdUser
}

func NewGetByIdUserController(useCaseGetById *application.GetByIdUser) *GetByIdUserController {
	return &GetByIdUserController{useCaseGetById: useCaseGetById}
}

func (getByIdUser *GetByIdUserController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}

	//long polling
	initialUser, err := getByIdUser.useCaseGetById.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateChan := make(chan domain.User)
	timeout := time.After(30 * time.Second)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			updatedUser, err := getByIdUser.useCaseGetById.Execute(int32(id))
			if err != nil {
				continue
			}
			if updatedUser != initialUser {
				updateChan <- updatedUser
				return
			}
		}
	}()

	select {
	case updatedUser := <-updateChan:
		ctx.JSON(http.StatusOK, gin.H{"user": updatedUser})
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"user": initialUser})
	}
}
