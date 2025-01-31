package main

import (
	product "api/src/Products/infraestructure/dependencies"
	routesProduct "api/src/Products/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	product.Init()

	defer product.CloseDB()

	r := gin.Default()
	routesProduct.Routes(r)
	r.Run()

}
