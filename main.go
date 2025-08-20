package main

import (
    "net/http"
    "log"
    "os"
    "html/template"
    "embed"
    "github.com/frooplexp/ytgy/ytgy"
)

//go:embed index.html
var index embed.FS

type Page struct {
    Term string
    Url string
    Source string
}

const httpListener = ":8080"

func main() {
    tmpl, err := template.ParseFS(index, "index.html")
    if err != nil {
        log.Printf("failed to parse template: %v\n", err)
        os.Exit(1)
    }

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        p := ytgy.Rand()
        if err := tmpl.Execute(w, Page{ p.Generate(), p.Url(), p.Source() }); err != nil {
            http.Error(w, "failed to parse template", http.StatusInternalServerError)
        }
    })

    log.Printf("starting http server on %s\n", httpListener)
    
    if err := http.ListenAndServe(httpListener, nil); err != nil {
        log.Printf("failed to serve http: %v\n", err)
        os.Exit(1)
    }
}
