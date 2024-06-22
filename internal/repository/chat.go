package repository

import (
	"Megic-core/internal/model"
	"database/sql"
)

type ChatRepository interface {
	Get(session string) ([]model.ChatModel, error)
	Create(chat model.ChatModel) error
}

type chatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) Get(session string) ([]model.ChatModel, error) {
	rows, err := r.db.Query("SELECT id, session, message, role FROM chats WHERE session = ?", session)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []model.ChatModel
	for rows.Next() {
		var chat model.ChatModel
		if err := rows.Scan(&chat.ID, &chat.Session, &chat.Message, &chat.Role); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *chatRepository) Create(chat model.ChatModel) error {
	_, err := r.db.Exec("INSERT INTO chats (session, message, role) VALUES (?, ?, ?)", chat.Session, chat.Message, chat.Role)
	return err
}
