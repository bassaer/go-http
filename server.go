package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	w.Write(dump)
	fmt.Println(string(dump))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("serving..")
	http.ListenAndServe(":8080", nil)
}
