
// main.go
package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "a"
    dbname   = "einvoice"
)

// Article - Our struct for all articles
type Article struct {
    Id      string    `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}


func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("pinged");

    Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}