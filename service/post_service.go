package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	models "web_blog/model"
	"web_blog/repository"
	"web_blog/support"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(db *support.DB) *PostService {
	return &PostService{
		repo: repository.NewMysqlPostRepo(db.SQL),
	}
}

func (ps *PostService) Fetch(w http.ResponseWriter, r *http.Request) {
	posts, _ := ps.repo.Fetch(r.Context(), 5)
	responseWithJson(w, http.StatusOK, posts)
}

func (ps *PostService) Create(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	json.NewDecoder(r.Body).Decode(&post)
	newId, err := ps.repo.Create(r.Context(), &post)
	fmt.Println(newId)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "server error")
	}
	responseWithJson(w, http.StatusOK, map[string]string{"message": "Created Successfully"})
}

func (ps *PostService) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	post := models.Post{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&post)
	payload, err := ps.repo.Update(r.Context(), &post)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Server Error")
	}
	responseWithJson(w, http.StatusOK, payload)
}

func (ps *PostService) FindById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	post, err := ps.repo.FindById(r.Context(), int64(id))
	if err != nil {
		responseWithError(w, http.StatusNoContent, "Content not found")
	}
	responseWithJson(w, http.StatusOK, post)
}
func (ps *PostService) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := ps.repo.Delete(r.Context(), int64(id))
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Server Error")
	}
	responseWithJson(w, http.StatusOK, map[string]string{"message": "Delete Successfully"})
}

func responseWithJson(writer http.ResponseWriter, code int, resp interface{}) {
	response, _ := json.Marshal(resp)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(response)
}

func responseWithError(writer http.ResponseWriter, code int, msg string) {
	responseWithJson(writer, code, map[string]string{"message": msg})
}
