package main

import (
	"blog_service/controller"
	"blog_service/model"
	"blog_service/repo"
	"blog_service/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init_db() *gorm.DB {
	connection_url := 
	"host=database user=postgres password=mypassword dbname=postgres port=5432"

	db, err := gorm.Open(postgres.Open(connection_url), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database")
		fmt.Println(err)
		return nil
	} else {
		fmt.Println("Succsessfully connected to the database!")
	}

	db.AutoMigrate(model.Blog{})
	db.AutoMigrate(model.Comment{})

	return db
}

func main() {
	db := init_db()
	if db == nil {
		return
	}

	router := mux.NewRouter()
	blog_router := router.PathPrefix("/blog").Subrouter()
	comment_router := router.PathPrefix("/blog/comment").Subrouter()

	blog_repo := repo.BlogRepo{Db: db}
	blog_service := service.BlogService{Repository: &blog_repo}
	blog_controller := controller.BlogController{Service: &blog_service}

	comment_repo := repo.CommentRepo{Db: db}
	comment_service := service.CommentService{Repository: &comment_repo}
	comment_controller := controller.CommentController{Service: &comment_service}

	blog_router.HandleFunc("/", blog_controller.GetAll).Methods("GET")
	blog_router.HandleFunc("/", blog_controller.Save).Methods("POST")
	blog_router.HandleFunc("/", blog_controller.UpdateBlog).Methods("PUT")
	blog_router.HandleFunc("/{id}", blog_controller.GetById).Methods("GET")
	blog_router.HandleFunc("/{id}", blog_controller.Delete).Methods("DELETE")
	blog_router.HandleFunc("/user/{id}", blog_controller.GetBlogsByAuthor).Methods("GET")
	blog_router.HandleFunc("/forUser/{id}", blog_controller.GetBlogsForUser).Methods("GET")

	comment_router.HandleFunc("/", comment_controller.GetAll).Methods("GET")
	comment_router.HandleFunc("/", comment_controller.Save).Methods("POST")
	comment_router.HandleFunc("/", comment_controller.Update).Methods("PUT")
	comment_router.HandleFunc("/{id}", comment_controller.GetById).Methods("GET")
	comment_router.HandleFunc("/{id}", comment_controller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
