package main

import (
	product "api/src/Products/infraestructure/dependencies"
	routesProduct "api/src/Products/infraestructure/routes"
	user "api/src/Users/infraestructure/dependencies"
	routesUser "api/src/Users/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	product.Init()
	user.Init()

	defer user.CloseDB()
	defer product.CloseDB()

	r := gin.Default()
	routesProduct.Routes(r)
	routesUser.Routes(r)
	r.Run()

}
