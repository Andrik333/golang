package main

import (
	"fmt"
	"log"
	"net/http"
)

type ShortLink struct {
	code string
	link string
}

var arrLink = []ShortLink{
	{"adwsegse", "http:/esefs/test"},
	{"jdrgjdfg", "http:/esefs/api/request"},
	{"45ujzdf2", "http:/esefs/knowledgebase"},
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if r.Method == http.MethodGet {
		code := r.URL.Path[1:]

		if code == "" {
			w.WriteHeader(400)
			fmt.Fprintln(w, "Код не указан")
			return
		}

		for _, value := range arrLink {
			if value.code == code {
				w.WriteHeader(307)
				w.Header().Set("Location", value.link)
				return
			}
		}

		w.WriteHeader(404)
		fmt.Fprintln(w, "Код не найден")
		return
	}

	if r.Method == http.MethodPost {
		link := r.FormValue("link")

		if link == "" {
			w.WriteHeader(400)
			fmt.Fprintln(w, "URL не указан")
			return
		}

		for _, value := range arrLink {
			if value.link == link {
				w.WriteHeader(201)
				fmt.Fprintln(w, value.code)
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/", GetHandler)

	err := http.ListenAndServe(":8282", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
