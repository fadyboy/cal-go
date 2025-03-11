package controllers

import (
	"fmt"
	"net/http"

	"github.com/fadyboy/lenslocked/models"
)

type Users struct {
	Templates struct {
		New Template
		SignIn Template
	}
	UserService *models.UserService
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
		fmt.Printf("error :%v", err)
		http.Error(w, "Error parsing request form", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Printf("error :%v", err)
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "user created:%+v", user)
	
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")

	u.Templates.SignIn.Execute(w, data)

}
