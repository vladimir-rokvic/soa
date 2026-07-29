package service

import (
	"blog_service/dto"
	"blog_service/model"
	"blog_service/repo"
	"fmt"
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

func (service *BlogService) Save(blogDto dto.BlogDTO) error {
	var blog = model.Blog{
		AuthorID: blogDto.AuthorID,
		Title: blogDto.Title,
		Description: blogDto.Description,
	}
	err := service.Repository.Save(blog)
	return err
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

func (service *BlogService) Delete(uuid string) error {
	blog, err := service.GetById(uuid)
	if err != nil {
		fmt.Printf("Blog by uuid: %s, not found", uuid)
		return err
	}

	err = service.Repository.Delete(blog)

	return err
}
