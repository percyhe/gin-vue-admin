package common

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {

	//articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	articleRouterWithRecord := Router.Group("article")
	var articleApi = v1.ApiGroupApp.CommonApiGroup.ArticleApi
	{
		//CRUD
	}
	{
		articleRouterWithRecord.GET("getArticleList", articleApi.GetArticleList)
	}
}
