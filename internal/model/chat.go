package model

type ChatModel struct {
	ID      int    `db:"id" json:"id"`
	Session string `db:"session" json:"session"`
	Message string `db:"message" json:"message"`
	Role    string `db:"role" json:"role"`
}
