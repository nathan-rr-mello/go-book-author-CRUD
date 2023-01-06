package models

type Author struct {
	ID 		uint 	`json:"id"`
	Name 	string 	`json:"name"`
	Age 	uint8 	`json:"age"`
	Gender	string	`json:"gender"`
	Books 	[]Book	`json:"books,omitempty"`
}