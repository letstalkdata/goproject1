package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/letstalkdata/goproject1/entity"
	"github.com/letstalkdata/goproject1/repsitory"
)

var (
	repo repsitory.PostRepository = repsitory.NewPostRepository()
)

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": Error getting the post`))
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}

func addPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": Error Decoding addPost`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
