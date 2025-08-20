package main

import (
    "net/http"
    "log"
    "os"
    "html/template"
    "embed"
    "ond/ond"
)

//go:embed index.html
var index embed.FS

type Page struct {
    Term string
    Url string
    Source string
}

const httpListener = ":8081"

func main() {
    tmpl, err := template.ParseFS(index, "index.html")
    if err != nil {
        log.Printf("failed to parse template: %v\n", err)
        os.Exit(1)
    }

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        p := ond.Rand()
        if err := tmpl.Execute(w, Page{ p.Generate(), p.Url(), p.Source() }); err != nil {
            http.Error(w, "failed to parse template", http.StatusInternalServerError)
        }
    })

    log.Printf("starting http server on %s\n", httpListener)
    
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Printf("failed to serve http: %v\n", err)
        os.Exit(1)
    }
}
