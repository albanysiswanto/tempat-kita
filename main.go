package main

import (
	"log"
	"net/http"
	"tempat-kita/config"
	authController "tempat-kita/controller/Authentication"
	courseController "tempat-kita/controller/Courses"
	homeController "tempat-kita/controller/Home"
)

func main() {
	config.ConnectDB()

	// Memanggil assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("vendors"))))

	// 1. Home
	http.HandleFunc("/", homeController.Index)

	// 2. Register & Login
	http.Handle("/course", authController.RequireLogin(authController.CheckRole(1, http.HandlerFunc(courseController.Index))))
	http.HandleFunc("/register", authController.RegisterHandler)
	http.HandleFunc("/login", authController.LoginHandler)

	// 3. User -> Course
	// http.HandleFunc("/course", courseController.Index)

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
