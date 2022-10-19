package structs

type Producer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Game struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Producer Producer `json:"producer"`
}
