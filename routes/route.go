package routes

import (
	"github.com/karimbaggari/ecommerce-cart-go/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup",controllers.Signup())
	incomingRoutes.POST("users/login",controllers.login())
	incomingRoutes.POST("admin/addproduct",controllers.productview())
	incomingRoutes.GET("users/productview",controllers.SearchProduct())
	incomingRoutes.GET("users/search",controllers.SearchProductByQuery())
}