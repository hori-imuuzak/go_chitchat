package route

import "net/http"

func PostThread(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("postThread"))
}

func ReadThread(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("readThread"))
}
