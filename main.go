package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

const (
	dbName = "web_blog"
	dbPass = "rootpw"
	dbHost = "localhost"
	dbPort = "3306"
)

var router *chi.Mux

var db *sql.DB

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dbSource)
	fmt.Println(db.Stats())
	catch(err)
}

func main() {
	routers()
	_ = http.ListenAndServe(":8005", Logger())
}

func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r)
	})
}
func catch(e error) {
	if e != nil {
		panic(e)

	}
}
func routers() *chi.Mux {
	router.Get("/posts", AllPosts)
	router.Get("/posts/{id}", DetailPost)
	router.Post("/posts", CreatePost)
	router.Put("/posts/{id}", UpdatePost)
	router.Delete("/post/{id}", DeletePost)
	return router
}

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func DeletePost(w http.ResponseWriter, r *http.Request) {

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post = Post{}
	json.NewDecoder(r.Body).Decode(&post)
	stmt, _ := db.Prepare("insert posts set title=?, content=?")
	_, err := stmt.Exec(post.Title, post.Content)
	catch(err)
	defer stmt.Close()
	responseWithJson(w, http.StatusCreated, map[string]string{"message": "successfully create"})

}

func DetailPost(w http.ResponseWriter, r *http.Request) {

}

func AllPosts(w http.ResponseWriter, r *http.Request) {

}

func responseWithJson(writer http.ResponseWriter, code int, resp map[string]string) {
	response, _ := json.Marshal(resp)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(response)

}
