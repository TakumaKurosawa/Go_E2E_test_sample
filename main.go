package main

import (
	"e2e/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// 接続テスト
	r.HandleFunc("/ping", handler.Ping)
	// ユーザ作成
	r.HandleFunc("/user", handler.CreateUser).Methods("POST")
	// ユーザ詳細
	r.HandleFunc("/user/{id}", handler.ShowUser).Methods("GET")

	fmt.Println("start a server on http://localhost:8011")
	log.Fatal(http.ListenAndServe(":8011", r))
}
