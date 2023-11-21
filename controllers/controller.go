package controllers

import (
	"github.com/gin-gonic/gin"
    "net/http"
)

func HashPassword(password string) string {

}


func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}


func Signup(password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err.Error()}}
			return
		}
		validationError := Validate.Struct(user)
		if validationError != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": validationError}}	
			return
		}		
		count, err := UserCollection.CountDocuments(ctx,bson.M{"email": user.Email})
		if err != nil {
		}
		}
}


func Login() gin.HandlerFunc {

}


func ProductViewerAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}

func SearchProductByQuery() gin.HandlerFunc {

}
