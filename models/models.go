package models

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbm   string  `json:"isbm"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
