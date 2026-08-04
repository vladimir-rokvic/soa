package service

import (
	"blog_service/dto"
	"blog_service/model"
	"blog_service/repo"
	"encoding/json"
	"fmt"
	"net/http"
)

type BlogService struct {
	Repository *repo.BlogRepo
}

func (service *BlogService) GetAll() ([]model.Blog, error) {
	blogs, err := service.Repository.GetAll()
	return blogs, err
}

func (service *BlogService) GetById(uuid string) (model.Blog, error) {
	blog, err := service.Repository.GetById(uuid)
	return blog, err
}

func (service *BlogService) Save(blogDto dto.BlogDTO) (model.Blog, error) {
	var blog = model.Blog{
		AuthorID: blogDto.AuthorID,
		Title: blogDto.Title,
		Description: blogDto.Description,
	}
	blog, err := service.Repository.Save(blog)
	return blog, err
}

func (service *BlogService) Update(blogDto dto.BlogDTO) error {
	blog, err := service.GetById(blogDto.ID.String())
	if err != nil {
		fmt.Printf("Error getting blog of Id: %s", blogDto.ID.String())
		return err
	}

	if blogDto.Title != "" {
		blog.Title = blogDto.Title
	}
	if blogDto.Description != "" {
		blog.Description = blogDto.Description
	}

	err = service.Repository.Update(blog)
	return err
}

func (service *BlogService) Delete(uuid string) (model.Blog, error) {
	blog, err := service.GetById(uuid)
	if err != nil {
		fmt.Printf("Blog by uuid: %s, not found", uuid)
		return blog, err
	}

	err = service.Repository.Delete(blog)

	return blog, err
}

func (service *BlogService) GetBlogsByAuthor(uuid string) ([]model.Blog, error) {
	blogs, err := service.Repository.GetBlogsByAuthor(uuid)
	if err != nil {
		fmt.Printf("Error getting blogs by author: %s\n", uuid)
		return nil, err
	}

	return blogs, err
}

func extractIds(users []dto.UserDTO) []string {
	ids := make([]string, 0, len(users))

	for _, u := range users {
		ids = append(ids, u.ID)
	}

	return ids
}

func (service *BlogService) GetBlogsForUser(id string) ([]model.Blog, error) {
	res, err := http.Get(
		"http://follower-service:8080/followers/users/" + id + "/follows")

	if err != nil {
		return nil, err
	}

	var users []dto.UserDTO
	err = json.NewDecoder(res.Body).Decode(&users)

	if err != nil {
		return nil, err
	}

	ids := extractIds(users)
	blogs, err := service.Repository.GetBlogsByAuthors(ids)

	return blogs, err
}
