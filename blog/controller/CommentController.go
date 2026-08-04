package controller

import (
	"blog_service/dto"
	"blog_service/model"
	"blog_service/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


type CommentController struct {
	Service *service.CommentService
}

func (controller *CommentController) GetAll(w http.ResponseWriter, r *http.Request) {
	blogs, err := controller.Service.GetAll()

	if err != nil {
		fmt.Println("Error getting all blogs")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func (controller *CommentController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting id from vars")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	blog, err := controller.Service.GetById(id)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return 
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func (controller *CommentController) Save(w http.ResponseWriter, r *http.Request) {
	var comment model.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	c, err := controller.Service.Save(&comment)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func (controller *CommentController) Update(w http.ResponseWriter, r *http.Request) {
	var u dto.CommentDTO

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	c, err := controller.Service.Update(&u)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func (controller *CommentController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting id from vars")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c, err := controller.Service.Delete(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}
