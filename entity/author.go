package entity

type Author struct {
	NAME  string `json:"name"`
	EMAIL string `json:"email"`
	BOOK  []Book `json:"book"`
}
