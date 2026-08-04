package controller

import (
	"blog_service/dto"
	"blog_service/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type BlogController struct {
	Service *service.BlogService
}

func (controller *BlogController) GetAll(writer http.ResponseWriter, req *http.Request) {
	blogs, err := controller.Service.GetAll()
	if err != nil {
		fmt.Println("Error fetching all blogs")
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blogs)
}

func (controller *BlogController) Save(writer http.ResponseWriter, req *http.Request) {
	var blogDto dto.BlogDTO

	json.NewDecoder(req.Body).Decode(&blogDto)

	blog, err := controller.Service.Save(blogDto)
	if err != nil {
		fmt.Println("Error saving blog")
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&blog)
}

func (controller *BlogController) GetById(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	uuid, ok := vars["id"]
	if !ok {
		fmt.Println("Error getting id from vars")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	blog, err := controller.Service.GetById(uuid)
	if err != nil {
		fmt.Printf("Error fetching blog by id: %s\n", uuid)
		fmt.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blog)
}

func (controller *BlogController) UpdateBlog(writer http.ResponseWriter, req *http.Request) {
	var blog dto.BlogDTO
	json.NewDecoder(req.Body).Decode(&blog)
	err := controller.Service.Update(blog)

	if err != nil {
		fmt.Println("Error updating blog")
		fmt.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blog)
}

func (controller *BlogController) GetBlogsByAuthor(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	auth_id, ok := vars["id"]
	if !ok {
		fmt.Printf("Error getting vars for author id: %s\n", auth_id);
		writer.WriteHeader(http.StatusBadRequest)
		return;
	}

	blogs, err := controller.Service.GetBlogsByAuthor(auth_id);

	if err != nil {
		fmt.Printf("Error getting blogs by auth: %s\n", auth_id)
		fmt.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-type", "application/json")
	json.NewEncoder(writer).Encode(blogs)
}

func (controller *BlogController) Delete(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	blog, err := controller.Service.Delete(id)
	if err != nil {
		fmt.Println("Error deleting blog")
		fmt.Println(err)
		writer.WriteHeader(http.StatusNotFound)
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-type", "application/json")
	json.NewEncoder(writer).Encode(blog)
}

func (controller *BlogController) GetBlogsForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting id from vars")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	blogs, err := controller.Service.GetBlogsForUser(id)

	if err != nil {
		fmt.Println("Error getting blogs for user")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}
