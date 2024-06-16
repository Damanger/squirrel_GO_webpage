package main

import (
    "fmt"
    "html/template"
    "net/http"
    "log"
)

func main() {
    // Configurar el manejador de plantillas
    templates := template.Must(template.ParseGlob("../public/*.html"))

    // Manejar la ruta raíz "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Renderizar el archivo index.html con los componentes nav y footer
        err := templates.ExecuteTemplate(w, "index.html", nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            log.Println(err)
        }
    })

    // Manejar la ruta "/about.html"
    http.HandleFunc("/about.html", func(w http.ResponseWriter, r *http.Request) {
        // Renderizar el archivo about.html con los componentes nav y footer
        err := templates.ExecuteTemplate(w, "about.html", nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            log.Println(err)
        }
    })

    // Manejar la ruta "/contact.html"
    http.HandleFunc("/contact.html", func(w http.ResponseWriter, r *http.Request) {
        // Renderizar el archivo contact.html con los componentes nav y footer
        err := templates.ExecuteTemplate(w, "contact.html", nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            log.Println(err)
        }
    })

    // Configurar el servidor para servir archivos estáticos desde el directorio "public"
    fs := http.FileServer(http.Dir("public"))
    http.Handle("/public/", http.StripPrefix("/public/", fs))

    fmt.Println("Server starting on port 3000...")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
