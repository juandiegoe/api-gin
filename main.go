package main

import (
	"fmt"
	"os"

	"github.com/JuanDiegoE/api-gin/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger())

	routes.Setup(server)

	godotenv.Load()
	port := os.Getenv("PORT")
	address := fmt.Sprintf(":%s", port)

	server.Run(address)
}
