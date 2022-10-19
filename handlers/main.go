package handlers

import (
	"favgames/methods"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllGames(context *gin.Context) {
	context.JSON(http.StatusOK, methods.GetAllGames())
}

func GetGame(context *gin.Context) {
	ID, err := strconv.Atoi(context.Param("ID"))
	game := methods.GetGame(ID)

	if err != nil {
		panic(err)
	}

	if game.Name == "" {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Game does not exist.",
		})
	} else {
		context.JSON(http.StatusOK, game)
	}
}

func GetGamesByPrice(context *gin.Context) {
	from, err := strconv.Atoi(context.Query("from"))

	if err != nil {
		panic(err)
	}

	to, err := strconv.Atoi(context.Query("to"))

	if err != nil {
		panic(err)
	}

	if from < 0 || to < 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "`from` or `to` parameter is less than 0.",
		})
	} else {
		games := methods.GetGamesByPrice(from, to)

		context.JSON(http.StatusOK, games)
	}

}
