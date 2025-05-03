package models

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Chat struct {
	ChatID int64  `json:"chat_id"`
	Topic  string `json:"topic"`
	Email  string `json:"email"`
}

type Message struct {
	MessageID int64  `json:"message_id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	ChatID    int64  `json:"chat_id"`
}
