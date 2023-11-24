package homeController

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/home/home.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
