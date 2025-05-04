package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Chat struct {
	ChatID int64  `json:"idChat"`
	Subject  string `json:"subject"`
	Email  string `json:"email"`
}

type Message struct {
	MessageID int64  `json:"idMessage"`
	Content    string `json:"content"`
	ChatID    int64  `json:"idChat"`
}
