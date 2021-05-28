package middlewares

import (
	"log"
	"net/http"

	"github.com/JuanDiegoE/api-gin/models"
	"github.com/JuanDiegoE/api-gin/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Saludo(c *gin.Context) {
	c.JSON(200, gin.H{
		"Saludo": "Bienvenido",
	})
}

func CreateUser(c *gin.Context) {
	var (
		user models.User
		exist bool
	)
	err := c.BindJSON(&user); if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	exist,err = models.UserExists(user)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if exist{
		c.JSON(http.StatusNotAcceptable, gin.H{"Message": "Email exists"})
	} else {
		user, err = models.NewUser(user)

		if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"User":user})
	}
}

func LoginUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user); if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userNew, err := models.LoginUser(user)

	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":"Email or password incorrect",
			"error": err.Error(),
		})
	} else {

		token := services.GenerateToken(userNew)

		cookie, err := c.Cookie("jwt")

        if err != nil {
            cookie = "NotSet"
            c.SetCookie("jwt", token, 3000, "/", "localhost", false, true)
        }

        log.Print("Cookie value: "+ cookie)

		c.JSON(http.StatusOK, gin.H{
			"Message": "Valid credentials!",
			"User" : userNew,
			"Token" : token,	
		})		
	}
}

func UserLogged(c *gin.Context) {
	
	cookie, err := c.Cookie("jwt")

	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	if len(cookie) == 0{

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated",
		})
			
	} else {

		token, err := services.ValidateToken(cookie)
		
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}

		claims := token.Claims.(*jwt.StandardClaims)

		c.JSON(http.StatusOK,claims)
	}

}

func Logout(c *gin.Context){
	cookie, err := c.Cookie("jwt")

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

    c.SetCookie("jwt","nil", -1, "/", "localhost", false, true)

	if len(cookie) > 0 {
		c.JSON(http.StatusOK,gin.H{
		"message": "success",
	})
	}

}
