package controller

import (
	"blog_service/dto"
	"blog_service/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type BlogController struct {
	Service *service.BlogService
}

func (controller *BlogController) GetAll(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	blogs, err := controller.Service.GetAll()
	if err != nil {
		fmt.Println("Error fetching all blogs")
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blogs)
}

func (controller *BlogController) Save(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var blogDto dto.BlogDTO

	json.NewDecoder(req.Body).Decode(&blogDto)

	err := controller.Service.Save(blogDto)
	if err != nil {
		fmt.Println("Error saving blog")
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&blogDto)
}

func (controller *BlogController) GetById(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	uuid := strings.TrimPrefix(req.URL.Path, "/blog/")
	blog, err := controller.Service.GetById(uuid)

	if err != nil {
		fmt.Println("Error fetching blog by id")
		fmt.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blog)
}

func (controller *BlogController) UpdateBlog(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var blog dto.BlogDTO
	json.NewDecoder(req.Body).Decode(&blog)
	err := controller.Service.Update(blog)

	if err != nil {
		fmt.Println("Error updating blog")
		fmt.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blog)
}
