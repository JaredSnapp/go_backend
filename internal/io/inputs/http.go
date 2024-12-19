package HTTPHandler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTTPHandler interface {
	serve()
}

// type Http struct {

// }

func Serve() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homePage)

	http.ListenAndServe(":3000", r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
