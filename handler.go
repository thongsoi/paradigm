package main

import (
	"html/template"
	"net/http"
)

type UserHandler struct {
	service UserService
	tmpl    *template.Template
}

func NewUserHandler(service UserService) *UserHandler {
	tmpl := template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/userlist.html",
	))

	return &UserHandler{service: service, tmpl: tmpl}
}

func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users", 500)
		return
	}
	h.tmpl.ExecuteTemplate(w, "index.html", users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	err := h.service.CreateUser(name, email)
	if err != nil {
		http.Error(w, "Unable to create user", 500)
		return
	}

	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users", 500)
		return
	}

	h.tmpl.ExecuteTemplate(w, "userlist.html", users)
}
