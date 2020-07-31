// main.go
package main

import (
	"encoding/json"
	"fmt"
	apiHanglers "github.com/filipsladek/einvoice/apiserver/handlers"
	"github.com/filipsladek/einvoice/apiserver/invoice"
	"github.com/filipsladek/einvoice/apiserver/postgres"
	"github.com/filipsladek/einvoice/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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
	Id      string `json:"Id"`
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

func handleRequests(storage storage.Storage, db postgres.DBConnector) {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", returnAllArticles)
	router.HandleFunc("/article", createNewArticle).Methods("POST")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", returnSingleArticle)

	router.PathPrefix("/api/invoice/{id}").Methods("GET").HandlerFunc(apiHanglers.GetInvoiceHandler(storage, db))
	router.PathPrefix("/api/invoice").Methods("POST").HandlerFunc(apiHanglers.CreateInvoiceHandler(storage, db))
	router.PathPrefix("/api/invoices").Methods("GET").HandlerFunc(apiHanglers.GetAllInvoicesHandler(storage, db))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, handlers.CORS(corsOptions...)(router)),
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func main() {
	storage := storage.InitStorage()
	storage.SaveObject("abc", "def")
	fmt.Println("stored")

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("pinged")
	//
	//Articles = []Article{
	//	Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	//	Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	//}

	dbConf := postgres.ConnectionConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "a",
		Database: "einvoice",
	}

	db := postgres.Connect(dbConf)
	defer db.Close()

	dbConnector := postgres.DBConnector{DB: db}

	// dummy data
	if len(dbConnector.GetAllInvoice()) == 0 {
		dbConnector.CreateInvoice(&invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectB"})
		dbConnector.CreateInvoice(&invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectC"})
		dbConnector.CreateInvoice(&invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectD"})
	}

	handleRequests(storage, dbConnector)
}

var corsOptions = []handlers.CORSOption{
	handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Accept", "token"}),
	handlers.AllowedOrigins([]string{"*"}),
	handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
}
