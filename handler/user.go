package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int
	Name string
}

func NewUser() User {
	return User{
		ID:   1,
		Name: "user",
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(NewUser())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// postman の変数更新を確認するため意図的にこう書いてます
	if vars["id"] != "1" {
		http.Error(w, "invalid user id", http.StatusBadGateway)
		return
	}

	res, err := json.Marshal(NewUser())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
