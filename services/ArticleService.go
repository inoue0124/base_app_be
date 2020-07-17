package services

import (
	"errors"

	"github.com/inoue0124/salon-be/db"
	"github.com/inoue0124/salon-be/forms"
	"github.com/inoue0124/salon-be/models"
)

type ArticleService struct{}

type Article models.Article

func (s ArticleService) PostArticle(ArticlePayload forms.Article) (Article, error) {
	db := db.GetDB()
	db.AutoMigrate(&Article{})
	var article Article
	article = Article{
		Title:   ArticlePayload.Title,
		Content: ArticlePayload.Content,
	}
	db.Create(&article)
	return article, nil
}

func (s ArticleService) GetAllArticles() ([]Article, error) {
	db := db.GetDB()
	var articles []Article
	db.Order("id desc").Limit(10).Find(&articles)
	if len(articles) == 0 {
		return nil, errors.New("No articles found!")
	}
	return articles, nil
}
