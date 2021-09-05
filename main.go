package main

import (
	"github.com/michaelwp/dragon/dragon"
	"github.com/michaelwp/dragon/handlers"
	"github.com/michaelwp/dragon/middlewares"
	"log"
)

func main() {
	r := dragon.NewRouter()

	// api
	api := r.Group("/api/v1")

	// home
	home := api.Group("/home")
	home.GET("", handlers.Home)

	// user
	user := api.Group("/users")
	user.Use(middlewares.Authorize())
	user.GET("", handlers.ListUser)
	user.POST("", handlers.CreateUser)
	user.GET("/:id", handlers.GetUser)

	// product
	product := api.Group("/products")
	product.GET("", handlers.ListProduct)
	product.POST("", handlers.CreateProduct)
	product.GET("/:type/:name", handlers.GetProduct)

	log.Fatal(r.Run(":8090"))
}
