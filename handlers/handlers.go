package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ShortLink struct {
	code string
	link string
}

var arrLink = []ShortLink{
	{"adwsegse", "https://ya.ru"},
	{"jdrgjdfg", "https://www.google.com"},
	{"45ujzdf2", "https://mail.ru"},
}

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/", PostHandler)
		r.Get("/*", GetHandler)
	})

	return r
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	code := chi.URLParam(r, "*")

	if code == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "Код не указан")
		return
	}

	for _, value := range arrLink {
		if value.code == code {
			http.Redirect(w, r, value.link, 307)
			return
		}
	}

	w.WriteHeader(404)
	fmt.Fprint(w, "Код не найден")
	return
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	link := r.FormValue("link")

	if link == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "URL не указан")
		return
	}

	for _, value := range arrLink {
		if value.link == link {
			w.WriteHeader(201)
			fmt.Fprint(w, value.code)
			return
		}
	}

	w.WriteHeader(500)
	fmt.Fprint(w, "Ошибка преобразования")
	return
}
