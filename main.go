package main

import (
	"github.com/karimbaggari/ecommerce-cart-go/controllers"
	"github.com/karimbaggari/ecommerce-cart-go/database"
	"github.com/karimbaggari/ecommerce-cart-go/middleware"
	"github.com/karimbaggari/ecommerce-cart-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port: os.Getenv("PORT")
	if port == "" {
        port = "8080"
    }

	app:= controllers.NewApplication(database.ProductData(database.client,"Products"), database.UserData(database.client,"Users"))

	router := gin.New()
	router.Use(gin.logger())

	routes.UserRoutes(router)
	router.User(middleware.Authentication)

	router.GET("/addtocart",app.AddToCart())
	router.GET("/removeitem", app.Removeitem())
	router.GET("/cartcheckout",app.BuyFromCart())
	router.GET("instantBuy",app.instantBuy())

	log.Fatal(router.run(":", port))

}
