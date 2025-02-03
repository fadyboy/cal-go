package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing request form", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Email: %s", r.Form.Get("email"))
	fmt.Fprintf(w, "Password: %s", r.Form.Get("password"))
}
