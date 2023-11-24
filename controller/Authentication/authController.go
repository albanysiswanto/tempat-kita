package authController

import (
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/login-register/register.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/login-register/login.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
