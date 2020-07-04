package route

import (
	"chitchat/infrastructures/repository"
	"chitchat/services"
	"net/http"
)

var userService = services.NewUserService(repository.NewUserRepository())

func Login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("login"))
}

func Signup(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("signup"))
}

func Logout(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("logout"))
}

func Authenticate(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	email := req.PostFormValue("email")
	password := req.PostFormValue("password")

	if userService.Authenticate(email, password) {
		w.Write([]byte("success"))
	} else {
		w.Write([]byte("failed"))
	}
}
