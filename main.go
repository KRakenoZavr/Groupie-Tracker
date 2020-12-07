package main

import (
	"fmt"
	"html/template"
	"net/http"

	parse "./parse"
)

type Er struct {
	Status  int
	Message string
}

var er Er

func indexHandler(w http.ResponseWriter, r *http.Request) {
	URL1 := "https://groupietrackers.herokuapp.com/api/artists"
	URL2 := "https://groupietrackers.herokuapp.com/api/relation"
	t, err := template.ParseFiles("index.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	group, err := parse.GetData(URL1, URL2)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "index.html", group)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	if status == http.StatusNotFound {
		er = Er{Status: 404, Message: "Page Not Found"}
	} else if status == http.StatusInternalServerError {
		er = Er{Status: 500, Message: "Server Error"}
	}

	w.WriteHeader(status)
	temp := template.Must(template.ParseGlob("templates/error.html"))
	temp.ExecuteTemplate(w, "error", er)
}

func main() {
	http.HandleFunc("/", indexHandler)

	css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	js := http.FileServer(http.Dir("js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	// fonts := http.FileServer(http.Dir("fonts"))
	// http.Handle("/fonts/", http.StripPrefix("/fonts/", fonts))

	fmt.Println("Listening port :3030")
	http.ListenAndServe(":3030", nil)
}
