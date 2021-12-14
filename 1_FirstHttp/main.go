package main

import "net/http"

func aboutHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("About Page"))
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World!!"))
	})

	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)
}
