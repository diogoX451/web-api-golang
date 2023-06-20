package entity

type Book struct {
	TITLE    string `json:"title"`
	SUBTITLE string `json:"subtitle"`
	PRICE    int    `json:"price"`
	AUTHOR   Author `json:"author"`
}
