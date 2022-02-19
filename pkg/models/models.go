package models

type Payment struct {
	ID       	int    	`json:"id"`
	Title    	string 	`json:"title"`
	Date	 	string	`json:"date"`
	Type	 	string 	`json:"type"`
	Comments 	string	`json:"comments"`
	CategoryId 	int 	`json:"category_id"`
}

type Category struct {
	ID       	int    	`json:"id"`
	Name    	string 	`json:"name"`
}
