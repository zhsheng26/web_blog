package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"web_blog/service"
	"web_blog/support"
)

const (
	dbName = "web_blog"
	dbUsr  = "root"
	dbPass = "rootpw"
	dbHost = "localhost"
	dbPort = "3306"
)

func main() {
	connectSQL, err := support.ConnectSQL(dbHost, dbPort, dbUsr, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	postService := service.NewPostService(connectSQL)
	router.Route("/", func(r chi.Router) {
		r.Mount("/post", postHandler(postService))
	})
	err = http.ListenAndServe(":8005", router)
	fmt.Println(err)
}

func postHandler(postService *service.PostService) http.Handler {
	r := chi.NewRouter()
	r.Get("/", postService.Fetch)
	r.Get("/{id:[0-9]+}", postService.FindById)
	r.Post("/", postService.Create)
	r.Put("/{id:[0-9]+}", postService.Update)
	r.Delete("/{id:[0-9]+}", postService.Delete)
	return r
}
