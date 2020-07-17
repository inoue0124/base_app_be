package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inoue0124/salon-be/forms"
	"github.com/inoue0124/salon-be/services"
)

type ArticleController struct{}

func (u ArticleController) GetAllArticles(c *gin.Context) {
	ArticleService := new(services.ArticleService)
	Articles, err := ArticleService.GetAllArticles()
	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve Article", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article founded!", "Articles": Articles})
	return
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

func (u ArticleController) PostArticle(c *gin.Context) {
	ArticlePayload := forms.Article{}
	err := c.ShouldBindJSON(&ArticlePayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to post Article", "error": err.Error()})
		c.Abort()
		return
	}
	ArticleService := new(services.ArticleService)
	Article, err := ArticleService.PostArticle(ArticlePayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to post Article", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article registered!", "Article": Article})
	return
}
