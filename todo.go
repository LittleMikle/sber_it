package todo

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Date        string `json:"date" db:"date"`
	Status      string `json:"status" db:"status"`
}

type TodoParams struct {
	Date   string
	Status string
}
