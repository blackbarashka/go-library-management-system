package http

import (
	"encoding/json"
	"lib/library"
	"net/http"
	"strconv"
)

type Handler struct {
	lib *library.Library
}

func NewHandler(lib *library.Library) *Handler {
	return &Handler{
		lib: lib,
	}
}

func (h *Handler) HttpAdd(w http.ResponseWriter, r *http.Request) {
	var newbook CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&newbook); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	book := library.NewBook(newbook.Title, newbook.Author, newbook.Pages)
	if err := h.lib.AddBook(book); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	erros := json.NewEncoder(w).Encode(book)
	if erros != nil {
		return
	}
}

func (h *Handler) HttpMarkAsRead(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.lib.MarkAsRead(id_int); err != nil {
		sendError(w, err, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	erros := json.NewEncoder(w).Encode(h.lib.Library[id_int])
	if erros != nil {
		return
	}
}

func (h *Handler) HttpGetBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}
	var book_ *library.Book
	book_, err = h.lib.GetBookInfo(id_int)
	if err != nil {
		sendError(w, err, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	erros := json.NewEncoder(w).Encode(book_)
	if erros != nil {
		return
	}

}

func (h *Handler) HttpFiltredGet(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var author *string
	if v := q.Get("author"); v != "" {
		author = &v
	}
	read, err := parseBoolPtr(q.Get("read"))
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}
	books, err := h.lib.GetListBook(author, read)
	if err != nil {
		sendError(w, err, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	errors := json.NewEncoder(w).Encode(books)
	if errors != nil {
		return
	}

}

func (h *Handler) HttpDeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}
	if err := h.lib.Remove(id_int); err != nil {
		sendError(w, err, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
