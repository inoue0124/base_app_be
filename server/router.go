package server

import (
	"github.com/gin-gonic/gin"
	"github.com/inoue0124/salon-be/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	userGroup := router.Group("users")
	{
		user := new(controllers.UserController)
		userGroup.POST("signup", user.SignupUser)
		userGroup.POST("signin", user.SigninUser)
		userGroup.POST("signout", user.SignoutUser)
		userGroup.GET("/:id", user.RetrieveUser)
	}

	articleGroup := router.Group("article")
	{
		article := new(controllers.ArticleController)
		articleGroup.POST("post", article.PostArticle)
		articleGroup.GET("", article.GetAllArticles)
	}

	return router
}
