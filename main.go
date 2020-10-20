package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {

	//router handle http requests

	//hello world handler
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	//homepage handler
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.PathPrefix("/assets").Handler(staticFileHandler).Methods("GET")

	//postlist handler
	staticFileDirectoryPosts := http.Dir("./posts/")
	staticFileHandlerPosts := http.StripPrefix("/posts", http.FileServer(staticFileDirectoryPosts))
	r.PathPrefix("/posts/").Handler(staticFileHandlerPosts).Methods("GET")
	r.PathPrefix("/posts").Handler(staticFileHandlerPosts).Methods("GET")

	//createPost handler
	staticFileDirectoryCreatePost := http.Dir("./createPost/")
	staticFileHandlerCreatePost := http.StripPrefix("/createPost", http.FileServer(staticFileDirectoryCreatePost))
	r.PathPrefix("/createPost/").Handler(staticFileHandlerCreatePost).Methods("GET")
	r.PathPrefix("/createPost").Handler(staticFileHandlerCreatePost).Methods("GET")

	//for getting and posting
	r.HandleFunc("/post", getPostHandler).Methods("GET")
	r.HandleFunc("/post", createPostHandler).Methods("POST")
	return r
}

func main() {
	//initialize server
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//function for hellow world
	fmt.Fprintf(w, "Hello World!")
}