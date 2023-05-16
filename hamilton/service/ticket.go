package service

type Ticket struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Concert struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Date  string `json:"date"`
		Venue string `json:"venue"`
	}
}
