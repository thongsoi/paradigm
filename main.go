package main

import (
	"log"
	"net/http"
)

func main() {
	db := InitDB()
	defer db.Close()

	service := NewPGUserService(db)
	handler := NewUserHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/register", handler.CreateUser)

	log.Println("Server running on http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}
