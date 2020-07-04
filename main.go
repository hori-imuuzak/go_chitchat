package main

import (
	"chitchat/route"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/login", route.Login)
	mux.HandleFunc("/signup", route.Signup)
	mux.HandleFunc("/logout", route.Logout)
	mux.HandleFunc("/authenticate", route.Authenticate)

	mux.HandleFunc("/thread/post", route.PostThread)
	mux.HandleFunc("/thread/read", route.ReadThread)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
