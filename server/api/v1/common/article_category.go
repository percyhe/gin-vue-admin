package common

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleCategoryApi struct {
}

var articleCategoryService = service.ServiceGroupApp.CommonServiceGroup.ArticleCategoryService

// CreateArticleCategory 创建ArticleCategory
// @Tags ArticleCategory
// @Summary 创建ArticleCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.ArticleCategory true "创建ArticleCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /articleCategory/createArticleCategory [post]
func (articleCategoryApi *ArticleCategoryApi) CreateArticleCategory(c *gin.Context) {
	var articleCategory common.ArticleCategory
	_ = c.ShouldBindJSON(&articleCategory)
	verify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	if err := utils.Verify(articleCategory, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := articleCategoryService.CreateArticleCategory(articleCategory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArticleCategory 删除ArticleCategory
// @Tags ArticleCategory
// @Summary 删除ArticleCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.ArticleCategory true "删除ArticleCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /articleCategory/deleteArticleCategory [delete]
func (articleCategoryApi *ArticleCategoryApi) DeleteArticleCategory(c *gin.Context) {
	var articleCategory common.ArticleCategory
	_ = c.ShouldBindJSON(&articleCategory)
	if err := articleCategoryService.DeleteArticleCategory(articleCategory); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteArticleCategoryByIds 批量删除ArticleCategory
// @Tags ArticleCategory
// @Summary 批量删除ArticleCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ArticleCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /articleCategory/deleteArticleCategoryByIds [delete]
func (articleCategoryApi *ArticleCategoryApi) DeleteArticleCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := articleCategoryService.DeleteArticleCategoryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArticleCategory 更新ArticleCategory
// @Tags ArticleCategory
// @Summary 更新ArticleCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.ArticleCategory true "更新ArticleCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /articleCategory/updateArticleCategory [put]
func (articleCategoryApi *ArticleCategoryApi) UpdateArticleCategory(c *gin.Context) {
	var articleCategory common.ArticleCategory
	_ = c.ShouldBindJSON(&articleCategory)
	verify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	if err := utils.Verify(articleCategory, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := articleCategoryService.UpdateArticleCategory(articleCategory); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArticleCategory 用id查询ArticleCategory
// @Tags ArticleCategory
// @Summary 用id查询ArticleCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query common.ArticleCategory true "用id查询ArticleCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /articleCategory/findArticleCategory [get]
func (articleCategoryApi *ArticleCategoryApi) FindArticleCategory(c *gin.Context) {
	var articleCategory common.ArticleCategory
	_ = c.ShouldBindQuery(&articleCategory)
	if rearticleCategory, err := articleCategoryService.GetArticleCategory(articleCategory.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rearticleCategory": rearticleCategory}, c)
	}
}

// GetArticleCategoryList 分页获取ArticleCategory列表
// @Tags ArticleCategory
// @Summary 分页获取ArticleCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query commonReq.ArticleCategorySearch true "分页获取ArticleCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /articleCategory/getArticleCategoryList [get]
func (articleCategoryApi *ArticleCategoryApi) GetArticleCategoryList(c *gin.Context) {
	var pageInfo commonReq.ArticleCategorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := articleCategoryService.GetArticleCategoryInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
