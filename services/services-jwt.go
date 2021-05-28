package services

import (
	"strconv"
	"time"

	"github.com/JuanDiegoE/api-gin/models"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "Secret"

func GenerateToken(user models.User) string {
	
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})
	
	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		panic(err)
	}
	return token
}

func ValidateToken(encodedToken string)(*jwt.Token, error){
	token, err := jwt.ParseWithClaims(encodedToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil{
		return nil,err
	}

	return token,nil
}