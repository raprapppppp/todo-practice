package todos

type Todos struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Task   string `json:"task"`
}
