package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"web_blog/service"
	"web_blog/support"
)

const (
	dbName = "web_blog"
	dbPass = "rootpw"
	dbHost = "localhost"
	dbPort = "3306"
)

func main() {
	//dbName := os.Getenv("DB_NAME")
	//dbPass := os.Getenv("DB_PASS")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	connectSQL, err := support.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//postService := service.NewPostService(connectSQL)
	//router.Route("/", func(r chi.Router) {
	//	r.Mount("/post", postHandler(postService))
	//})
	//router.Route("/", func(r chi.Router) {
	//	r.Mount("/page", pageHandler(pageService))
	//})
	pageService := service.NewPageService(connectSQL)
	http.HandleFunc("/page", pageService.Posts)
	err = http.ListenAndServe(":8005", nil)
	fmt.Println(err)
}

func cross(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8081")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func postHandler(postService *service.PostService) http.Handler {
	r := chi.NewRouter()
	r.Use(cross)
	r.Get("/", postService.Fetch)
	r.Get("/{id:[0-9]+}", postService.FindById)
	r.Post("/", postService.Create)
	r.Put("/{id:[0-9]+}", postService.Update)
	r.Delete("/{id:[0-9]+}", postService.Delete)
	return r
}
