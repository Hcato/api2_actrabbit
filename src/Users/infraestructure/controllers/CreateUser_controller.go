package controllers

import (
	"api/src/Users/application"
	"api/src/Users/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	CreateUseCase *application.CreateUserUsecase
}

func NewCreateProductController(CreateUseCase *application.CreateUserUsecase) *CreateUserController {
	return &CreateUserController{CreateUseCase: CreateUseCase}
}

func (createUser *CreateUserController) Execute(c *gin.Context) {
	var User domain.User

	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createUser.CreateUseCase.Execute(User.Name, User.Lastname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registrado"})
}
