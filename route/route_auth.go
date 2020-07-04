package route

import "net/http"

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
	w.Write([]byte("authenticate"))
}
