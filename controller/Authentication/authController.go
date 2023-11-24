package authController

import (
	"html/template"
	"net/http"
	"tempat-kita/entities"
	authModel "tempat-kita/model/Authentication"
	"time"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("view/login-register/register.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var user entities.User

		user.Email = r.FormValue("email")
		user.Name = r.FormValue("fullname")
		user.Password = r.FormValue("password")
		user.Created_At = time.Now()
		user.Updated_At = time.Now()

		if success := authModel.Register(user); !success {
			temp, err := template.ParseFiles("view/login-register/register.html")
			if err != nil {
				http.Error(w, "Gagal mendaftarkan user", http.StatusInternalServerError)
			}
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func CheckRole(role int, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Contoh: role dari sesi
		sessionRole := 1

		// Periksa apakah pengguna memiliki role yang sesuai
		if sessionRole != role {
			http.Error(w, "Tidak memiliki izin", http.StatusForbidden)
			return
		}

		// Lanjutkan ke handler berikutnya jika role sesuai
		next.ServeHTTP(w, r)
	}
}

func RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionRole := 1
		if sessionRole == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Lanjutkan ke handler berikutnya jika pengguna telah login
		next.ServeHTTP(w, r)
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("view/login-register/login.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var user entities.User

		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := authModel.GetEmailLogin(email)
		if err != nil {
			http.Error(w, "Gagal login, Email atau Password Salah!", http.StatusInternalServerError)
			return
		}

		err = authModel.VerifyPassword(user.Password, password)
		if err != nil {
			http.Error(w, "Gagal login, Password Salah!", http.StatusInternalServerError)
			return
		}

		if user.Role == 1 {
			http.Redirect(w, r, "/course", http.StatusSeeOther)
		}
	}
}
