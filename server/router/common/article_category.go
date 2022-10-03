package common

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArticleCategoryRouter struct {
}

// InitArticleCategoryRouter 初始化 ArticleCategory 路由信息
func (s *ArticleCategoryRouter) InitArticleCategoryRouter(Router *gin.RouterGroup) {
	articleCategoryRouter := Router.Group("articleCategory").Use(middleware.OperationRecord())
	articleCategoryRouterWithoutRecord := Router.Group("articleCategory")
	var articleCategoryApi = v1.ApiGroupApp.CommonApiGroup.ArticleCategoryApi
	{
		articleCategoryRouter.POST("createArticleCategory", articleCategoryApi.CreateArticleCategory)   // 新建ArticleCategory
		articleCategoryRouter.DELETE("deleteArticleCategory", articleCategoryApi.DeleteArticleCategory) // 删除ArticleCategory
		articleCategoryRouter.DELETE("deleteArticleCategoryByIds", articleCategoryApi.DeleteArticleCategoryByIds) // 批量删除ArticleCategory
		articleCategoryRouter.PUT("updateArticleCategory", articleCategoryApi.UpdateArticleCategory)    // 更新ArticleCategory
	}
	{
		articleCategoryRouterWithoutRecord.GET("findArticleCategory", articleCategoryApi.FindArticleCategory)        // 根据ID获取ArticleCategory
		articleCategoryRouterWithoutRecord.GET("getArticleCategoryList", articleCategoryApi.GetArticleCategoryList)  // 获取ArticleCategory列表
	}
}
