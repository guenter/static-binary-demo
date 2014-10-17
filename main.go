package main

import (
    "net/http"
    "fmt"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s %s from %s\n", r.Method, r.RequestURI, r.RemoteAddr)
    fmt.Fprintf(w, "Hi there, I love %s!\n\n", r.URL.Path[1:])
    for _, e := range os.Environ() {
        fmt.Fprintln(w, e)
    }
}

func main() {
    listen := fmt.Sprintf(":%s", os.Getenv("PORT"))
    fmt.Printf("Listening on %s\n", listen)
    http.HandleFunc("/", handler)
    http.ListenAndServe(listen, nil)
}
