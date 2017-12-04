package controllers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/xtda/chi-upper-db-boilerplate/models"
)

type IndexResource struct{}

func (rs IndexResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.Home) // GET /
	r.NotFound(rs.Home)
	return r
}

// Home grabs the home page
func (rs IndexResource) Home(w http.ResponseWriter, r *http.Request) {
	var testdata, err = models.AllTest(r.Context())
	if err != nil {
		panic(err)
	}
	log.Printf("%s", testdata)
	w.Write([]byte("Hello World"))
}
