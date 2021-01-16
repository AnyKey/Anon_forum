package main

import (
	"awesomeProject/handlers"
	"awesomeProject/model"
	"awesomeProject/repository"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func MustDBConn() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "non-root:123@tcp(127.0.0.1:3306)/BBD")
	if err != nil {
		log.Fatalln(err)
	}
	if db.Ping() != nil {
		log.Fatalln(err)
	}
	return db
}



func main() {
	conn := MustDBConn()

	l, _ := zap.NewDevelopment()
	logger := l.Sugar()

	router := mux.NewRouter()
	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			h.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/api/message/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		var err error
		defer func() {
			if err != nil {
				_ = handlers.WriteJsonToResponse(writer,map[string]string{
					"error": err.Error(),
				})
			}
			writer.WriteHeader(http.StatusOK)
		}()
		bytes, _ := ioutil.ReadAll(request.Body)
		var post model.AddMessagePostRequest
		err = json.Unmarshal(bytes, &post)
		if err != nil {
			logger.Error("Error with data")
			return
		}

		err = handlers.MessageHandler{Conn: conn, Logger: logger, Repo: repository.MessageRepository{Conn: conn}}.AddMessage(post.CategoryID,post.ParentID,post.Text)
		if err != nil {
			return
		}

		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodPost)

	router.HandleFunc("/api/category/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		var err error
		defer func() {
			if err != nil {
				_ = handlers.WriteJsonToResponse(writer,map[string]string{
					"error": err.Error(),
				})
			}
			writer.WriteHeader(http.StatusOK)
		}()

		messages, err := handlers.CategoryHandler{Conn: conn, Logger: logger,Repo:repository.CategoryRepository{Conn: conn}}.GetCategories()
		if err != nil {
			return
		}

		_ = handlers.WriteJsonToResponse(writer, messages)
	}).Methods(http.MethodGet)
	router.HandleFunc("/api/{category_id}/messages/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		var err error
		defer func() {
			if err != nil {
				_ = handlers.WriteJsonToResponse(writer,map[string]string{
					"error": err.Error(),
				})
			}
			writer.WriteHeader(http.StatusOK)
		}()
		vars := mux.Vars(request)
		categoryID, err := strconv.Atoi(vars["category_id"])
		if err != nil {
			return
		}

		messages, err := handlers.MessageHandler{Conn: conn, Logger: logger,Repo: repository.MessageRepository{Conn: conn}}.GetThreads(categoryID)
		if err != nil {
			return
		}

		_ = handlers.WriteJsonToResponse(writer, messages)
	}).Methods(http.MethodGet)
	router.HandleFunc("/api/{category_id}/messages/{parent_id}/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		var err error
		defer func() {
			if err != nil {
				_ = handlers.WriteJsonToResponse(writer,map[string]string{
					"error": err.Error(),
				})
			}
			writer.WriteHeader(http.StatusOK)
		}()
		vars := mux.Vars(request)
		categoryID, err := strconv.Atoi(vars["category_id"])
		if err != nil {
			return
		}

		parentID, err := strconv.Atoi(vars["parent_id"])
		if err != nil {
			return
		}

		messages, err := handlers.MessageHandler{Conn: conn, Logger: logger, Repo: repository.MessageRepository{Conn: conn}}.GetMessagesByThread(categoryID,int64(parentID))
		if err != nil {
			return
		}

		_ = handlers.WriteJsonToResponse(writer, messages)
	}).Methods(http.MethodGet)

	srv := &http.Server{
		Handler: router,
		Addr:    ":8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
