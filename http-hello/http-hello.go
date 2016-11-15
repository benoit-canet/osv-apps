package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "C"
)

func main() {
	GoMain();
}

//export GoMain
func GoMain() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}