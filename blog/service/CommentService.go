package service

import (
	"blog_service/dto"
	"blog_service/model"
	"blog_service/repo"
)


type CommentService struct {
	Repository *repo.CommentRepo
}

func (service *CommentService) GetById(id string) (model.Comment, error) {
	c, err := service.Repository.GetById(id)

	return c, err
}

func (service *CommentService) GetAll() ([]model.Comment, error) {
	c, err := service.Repository.GetAll()

	return c, err
}

func (service *CommentService) GetByBlogId(id string) ([]model.Comment, error) {
	c, err := service.Repository.GetByBlogId(id)

	return c, err
}

func (service *CommentService) Save(comment *model.Comment) (*model.Comment, error) {
	c, err := service.Repository.Save(comment)

	return c, err
}

func (service *CommentService) Update(newC *dto.CommentDTO) (*model.Comment, error) {
	id := newC.ID
	
	c, err := service.Repository.GetById(id)
	if err != nil {
		return &c, err
	}
	
	c.Text = newC.Text

	ret, err := service.Repository.Update(&c)

	return ret, err
}

func (service *CommentService) Delete(id string) (model.Comment, error) {
	c, err := service.Repository.GetById(id)
	if err != nil {
		return c, err
	}

	err = service.Repository.Delete(&c)
	return c, err
}
