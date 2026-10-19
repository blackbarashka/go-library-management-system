package library

import (
	"sync"
	"time"
)

type Library struct {
	Library   map[int]*Book
	currentId int
	mtx       sync.Mutex
}

func NewLibrary() *Library {
	return &Library{
		Library: make(map[int]*Book),
	}
}

func (l *Library) AddBook(book *Book) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if len(book.Author) < 1 || len(book.Title) < 1 {
		return NonCorrentInput
	}

	book.Id = l.currentId
	book.AddTime = time.Now()
	l.Library[l.currentId] = book
	l.currentId++
	return nil
}

func (l *Library) MarkAsRead(id int) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.Library[id]; !ok {
		return Notfound
	}

	t := time.Now()
	l.Library[id].IsRead = true
	l.Library[id].ReadTime = &t
	return nil
}

func (l *Library) GetBookInfo(id int) (*Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.Library[id]; !ok {
		return nil, Notfound
	}
	return l.Library[id], nil
}

func (l *Library) GetListBook(author *string, isRead *bool) ([]Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	books := make([]Book, 0, len(l.Library))

	for _, book := range l.Library {
		if author != nil && book.Author != *author {
			continue
		}

		if isRead != nil && book.IsRead != *isRead {
			continue
		}

		books = append(books, *book)
	}

	return books, nil
}

func (l *Library) Remove(id int) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.Library[id]; !ok {
		return Notfound
	}
	delete(l.Library, id)
	return nil
}
