package main
import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Post struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

var posts []Post

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the post list to a json file
	postListBytes, err := json.Marshal(posts)

	// print error to console if there is any
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//write to post list if successful
	w.Write(postListBytes)
}
func createPostHandler(w http.ResponseWriter, r *http.Request) {
	//new post object
	post := Post{}

	//parse html form into object
	err := r.ParseForm()

	//error handling again
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse data from form into post object
	post.Title = r.Form.Get("title")
	post.Content = r.Form.Get("content")

	//append new post object into our post list
	posts = append(posts, post)

	//redirect user to post list
	http.Redirect(w, r, "/posts/", http.StatusFound)
}