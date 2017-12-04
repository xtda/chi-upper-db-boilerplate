package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"upper.io/db.v3/postgresql"

	"github.com/xtda/chi-upper-db-boilerplate/controllers"
	"github.com/xtda/chi-upper-db-boilerplate/models"
)

func main() {
	var (
		config   = LoadConfiguration("config.json")
		settings = postgresql.ConnectionURL{
			Host:     config.Database.Host,
			Database: config.Database.Name, // "remorsev3_dev",
			User:     config.Database.User,
			Password: config.Database.Pasasword,
		}
	)
	db, err := models.OpenDB(settings)
	if err != nil {
		log.Panic(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.WithValue("db", db))
	r.Mount("/", controllers.IndexResource{}.Routes())

	//addr := ":3000"
	s := &http.Server{
		Addr:           config.ListenAddress,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		fmt.Printf("server listening on %s\n", config.ListenAddress)
		errc <- s.ListenAndServe()
	}()
	log.Fatalln("exit", <-errc)
}
