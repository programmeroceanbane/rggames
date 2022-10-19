package main

import (
	"favgames/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/games", handlers.GetAllGames)
	router.GET("/games/price", handlers.GetGamesByPrice)
	router.GET("/games/:ID", handlers.GetGame)

	router.Run()
}
