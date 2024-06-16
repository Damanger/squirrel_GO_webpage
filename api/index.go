package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

var templates = template.Must(template.ParseGlob("public/*.html"))

func handler(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/":
        renderTemplate(w, "index.html")
    case "/about.html":
        renderTemplate(w, "about.html")
    case "/contact.html":
        renderTemplate(w, "contact.html")
    default:
        http.Error(w, "Not found", http.StatusNotFound)
    }
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
    err := templates.ExecuteTemplate(w, tmpl, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Println(err)
    }
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server starting on port 3000...")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
