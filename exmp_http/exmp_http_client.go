package exmphttp

import (
	"fmt"
	"net/http"
)

func StartServer() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("testing"))
	})
	http.ListenAndServe(":8080", nil)

}

func hello(w http.ResponseWriter, req *http.Request) {
	if value, ok := req.Header["Accept"]; ok {
		fmt.Println("header =", value)
	}
	w.Write([]byte("Hello, world"))
}
