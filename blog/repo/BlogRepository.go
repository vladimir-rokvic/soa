package repo

import (
	"blog_service/model"

	"gorm.io/gorm"
)

type BlogRepo struct {
	Db *gorm.DB
}

func (repo *BlogRepo) GetAll() ([]model.Blog, error) {
	var blogs []model.Blog
	result := repo.Db.Find(&blogs)

	return blogs, result.Error
}

func (repo *BlogRepo) GetById(uuid string) (model.Blog, error) {
	var blog model.Blog
	result := repo.Db.Find(&blog, "id = ?", uuid)

	return blog, result.Error
}

func (repo *BlogRepo) Save(blog model.Blog) (model.Blog, error) {
	result := repo.Db.Create(&blog)
	return blog, result.Error
}

func (repo *BlogRepo) Update(blog model.Blog) error {
	result := repo.Db.Save(&blog)
	return result.Error
}

func (repo *BlogRepo) Delete(blog model.Blog) error {
	result := repo.Db.Delete(&blog)
	return result.Error
}

func (repo *BlogRepo) GetBlogsByAuthor(uuid string) ([]model.Blog, error) {
	var blogs []model.Blog
	result := repo.Db.Find(&blogs, "author_id = ?", uuid);

	return blogs, result.Error
}
