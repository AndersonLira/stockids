package model

//Family represents one group
type Family struct {
	ID          string   `json:"id"`
	CreatedAt   int64    `json:"created_at"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Avatar      string   `json:"avatar"`
	UserIDS     []string `json:"user_ids"`
}
