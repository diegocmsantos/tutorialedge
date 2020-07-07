package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getArticle")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	vars := mux.Vars(r)
	id := vars["id"]

	var updatedArticle Article
	json.Unmarshal(reqBody, &updatedArticle)
	fmt.Println(Articles)
	for i, article := range Articles {
		if article.Id == id {
			foundArticle := &Articles[i]
			foundArticle.Title = updatedArticle.Title
			foundArticle.Desc = updatedArticle.Desc
			foundArticle.Content = updatedArticle.Content
			updatedArticle.Id = id
		}
	}
	fmt.Println(Articles)
	json.NewEncoder(w).Encode(updatedArticle)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for index, article := range Articles {
		if key == article.Id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", createArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}", getArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}

	handleRequests()
}
