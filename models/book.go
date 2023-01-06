package models

type Book struct {
	ID 			uint	`json:"id"`
	Name		string	`json:"name"`
	AuthorID	uint	`json:"author_id"`
	Pages		int		`json:"pages"`	 
	Rating 		float64	`json:"rating"`  
}