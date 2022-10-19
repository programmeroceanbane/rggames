package methods

import (
	"favgames/structs"
)

func GetAllProducers() map[string]structs.Producer {
	producers := map[string]structs.Producer{
		"rockstar_games": {ID: 1, Name: "Rockstar Games", Description: "lorem ipsum"},
		"techland":       {ID: 2, Name: "Techland", Description: "lorem ipsum"},
	}

	return producers
}

func GetAllGames() []structs.Game {
	games := []structs.Game{
		{ID: 1, Name: "Red Dead Redemption 2", Price: 250, Producer: GetAllProducers()["rockstar_games"]},
		{ID: 2, Name: "Grand Theft Auto V", Price: 120, Producer: GetAllProducers()["rockstar_games"]},
		{ID: 3, Name: "Dying Light", Price: 100, Producer: GetAllProducers()["techland"]},
	}

	return games
}

func GetGame(ID int) structs.Game {
	for _, value := range GetAllGames() {
		if value.ID == ID {
			return value
		}
	}

	return structs.Game{}
}

func GetGamesByPrice(from, to int) []structs.Game {
	games := []structs.Game{}

	for _, value := range GetAllGames() {
		if value.Price >= from && value.Price <= to {
			games = append(games, value)
		}
	}

	return games
}
