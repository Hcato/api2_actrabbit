package routes

import (
	"api/src/Users/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/users")
	createUser := dependencies.GetCreateUserController().Execute
	getAllUser := dependencies.GetGetAllUserController().Execute
	getIdUser := dependencies.GetGetByIdUserController().Execute
	deleteUser := dependencies.GetDeleteUserController().Execute
	updateUser := dependencies.GetUpdateUserController().Execute

	routes.POST("/", createUser)
	routes.GET("/", getAllUser)
	routes.GET("/:id", getIdUser)
	routes.DELETE("/:id", deleteUser)
	routes.PUT("/:id", updateUser)
}
