package models

type Message struct {
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}
