package main

import (
	"html/template"
	"os"
)

type User struct {
	Name    string
	Bio     string
	Address Address
}

type Address struct {
	Street   string
	PostCode int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{Name: "Jason Bourne", Bio: "Sharpman things", Address: Address{Street: "1 Logos street", PostCode: 123}}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
