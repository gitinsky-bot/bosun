package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var TSDBHttp string
var templates *template.Template
var router = mux.NewRouter()

func Listen(addr, dir, tsdbhttp string) error {
	TSDBHttp = tsdbhttp
	var err error
	templates, err = template.New("").ParseFiles(
		dir + "/templates/chart.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	router.HandleFunc("/", Index)
	router.HandleFunc("/api/chart", Chart)
	http.Handle("/", router)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	log.Println("web listening on", addr)
	return http.ListenAndServe(addr, nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "chart.html", nil)
}
