package routes

import (
	"api/src/Products/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/products")
	createProduct := dependencies.GetCreateProductController().Execute
	getAllProduct := dependencies.GetGetAllProductController().Execute
	deleteProducts := dependencies.GetDeleteProductController().Execute
	updateProducts := dependencies.GetUpdateProductController().Execute
	getByIdProducts := dependencies.GetByIdProductController().Execute

	routes.POST("/", createProduct)
	routes.GET("/", getAllProduct)
	routes.DELETE("/:id", deleteProducts)
	routes.PUT("/:id", updateProducts)
	routes.GET("/:id", getByIdProducts)
}
