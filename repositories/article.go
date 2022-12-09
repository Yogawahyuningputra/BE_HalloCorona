package repositories

import (
	"be_corona/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	AddArticle(article models.Article) (models.Article, error)
	GetArticle(ID int) (models.Article, error)

	FindArticles() ([]models.Article, error)
	UpdateArticle(article models.Article) (models.Article, error)
	DeleteArticle(article models.Article) (models.Article, error)
}

func RepositoryArticle(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) AddArticle(article models.Article) (models.Article, error) {
	err := r.db.Create(&article).Error

	return article, err
}

func (r *repository) FindArticles() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Find(&articles).Error

	return articles, err
}

func (r *repository) GetArticle(ID int) (models.Article, error) {
	var article models.Article
	err := r.db.First(&article, ID).Error

	return article, err
}

func (r *repository) UpdateArticle(article models.Article) (models.Article, error) {
	err := r.db.Save(&article).Error

	return article, err
}

func (r *repository) DeleteArticle(article models.Article) (models.Article, error) {
	err := r.db.Delete(&article).Error

	return article, err
}
