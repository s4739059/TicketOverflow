package service

type Concert struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Date  string `json:"date"`
	Venue string `json:"venue"`
	Seats struct {
		Max       int `json:"max"`
		Purchased int `json:"purchased"`
	}
}
