package library

import "time"

type Book struct {
	Id       int        `json:"id"`
	Title    string     `json:"title"`
	Author   string     `json:"author"`
	Pages    int        `json:"pages"`
	IsRead   bool       `json:"is_read"`
	AddTime  time.Time  `json:"add_time"`
	ReadTime *time.Time `json:"read_time"`
}

func NewBook(title, author string, pages int) *Book {
	return &Book{
		Id:       0,
		Title:    title,
		Author:   author,
		Pages:    pages,
		IsRead:   false,
		AddTime:  time.Now(),
		ReadTime: nil,
	}
}
