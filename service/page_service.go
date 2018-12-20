package service

import (
	"html/template"
	"net/http"
	"web_blog/model"
	"web_blog/repository"
	"web_blog/support"
)

type TodoPageData struct {
	Posts []*model.Post
}
type PageService struct {
	repo repository.PostRepository
}

func NewPageService(db *support.DB) *PageService {
	return &PageService{
		repo: repository.NewMysqlPostRepo(db.SQL),
	}
}

func (page *PageService) Posts(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/posts.html"))
	posts, _ := page.repo.Fetch(r.Context(), 10)
	data := TodoPageData{
		Posts: posts,
	}
	tmpl.Execute(w, data)
}
