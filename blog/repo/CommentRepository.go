package repo

import (
	"blog_service/model"

	"gorm.io/gorm"
)


type CommentRepo struct {
	Db *gorm.DB
}

func (repo *CommentRepo) Save(c *model.Comment) (*model.Comment, error) {
	result := repo.Db.Create(c)

	return c, result.Error
}

func (repo *CommentRepo) Delete(c *model.Comment) error {
	result := repo.Db.Delete(c)

	return result.Error
}

func (repo *CommentRepo) Update(c *model.Comment) (*model.Comment, error) {
	result := repo.Db.Save(c)

	return c, result.Error
}

func (repo *CommentRepo) GetAll() ([]model.Comment, error) {
	var comments []model.Comment

	result := repo.Db.Find(&comments)

	return comments, result.Error
}

func (repo *CommentRepo) GetByBlogId(id string) ([]model.Comment, error) {
	var comments []model.Comment

	result := repo.Db.Find(&comments, "blog_id = ?", id)

	return comments, result.Error
}

func (repo *CommentRepo) GetById(id string) (model.Comment, error) {
	var comment model.Comment

	result := repo.Db.First(&comment, "id = ?", id)

	return comment, result.Error
}

func (repo *CommentRepo) GetByAuthorId(id string) ([]model.Comment, error) {
	var comments []model.Comment

	result := repo.Db.Find(&comments, "author_id = ?", id)

	return comments, result.Error
}
