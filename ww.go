package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pingHandler)

	userHandler := &UserHandler{
		userDB: make([]string, 0),
	}
	http.Handle("/user", userHandler)

	http.ListenAndServe(":8080", nil)
}

func pingHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("pong"))
}

type UserHandler struct {
	userDB []string
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// TODO
		h.createUser(w, r)
	default:
		w.Write([]byte("unhandled method"))
	}
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("invalid request"))
		return
	}

	h.userDB = append(h.userDB, string(body))
}
