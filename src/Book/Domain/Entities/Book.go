package entities

type Book struct{
	ID          int64    `json:"id"`
	Title       string `json:"title"`
	Year        int      `json:"year"`
	
}