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
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	routesProduct.Routes(r)
	routesUser.Routes(r)
	r.Run()

}
