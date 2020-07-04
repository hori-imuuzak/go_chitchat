package route

import (
	"chitchat/domains/model"
	"chitchat/infrastructures/repository"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("login"))
}

func Signup(w http.ResponseWriter, req *http.Request) {
	// test
	repo := repository.NewUserRepository()
	res := repo.Create(model.User{
		Uuid:  "test-uuid-1234",
		Name:  "hori",
		Email: "hogehoge@gmail.com",
	})
	w.Write([]byte(fmt.Sprintf("%v", res)))
}

func Logout(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("logout"))
}

func Authenticate(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("authenticate"))
}
