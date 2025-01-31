package routes

import (
	"api/src/Users/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/users")
	createUser := dependencies.GetCreateUserController().Execute
	getAllUser := dependencies.GetGetAllUserController().Execute
	deleteUser := dependencies.GetDeleteUserController().Execute
	updateUser := dependencies.GetUpdateUserController().Execute

	routes.POST("/", createUser)
	routes.GET("/", getAllUser)
	routes.DELETE("/:id", deleteUser)
	routes.PUT("/:id", updateUser)
}
