package common

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ArticleCategoryService struct {
}

// CreateArticleCategory 创建ArticleCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleCategoryService *ArticleCategoryService) CreateArticleCategory(articleCategory common.ArticleCategory) (err error) {
	err = global.GVA_DB.Create(&articleCategory).Error
	return err
}

// DeleteArticleCategory 删除ArticleCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleCategoryService *ArticleCategoryService) DeleteArticleCategory(articleCategory common.ArticleCategory) (err error) {
	err = global.GVA_DB.Delete(&articleCategory).Error
	return err
}

// DeleteArticleCategoryByIds 批量删除ArticleCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleCategoryService *ArticleCategoryService) DeleteArticleCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]common.ArticleCategory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateArticleCategory 更新ArticleCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleCategoryService *ArticleCategoryService) UpdateArticleCategory(articleCategory common.ArticleCategory) (err error) {
	err = global.GVA_DB.Save(&articleCategory).Error
	return err
}

// GetArticleCategory 根据id获取ArticleCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleCategoryService *ArticleCategoryService) GetArticleCategory(id uint) (articleCategory common.ArticleCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&articleCategory).Error
	return
}

// GetArticleCategoryInfoList 分页获取ArticleCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleCategoryService *ArticleCategoryService) GetArticleCategoryInfoList(info commonReq.ArticleCategorySearch) (list []common.ArticleCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&common.ArticleCategory{})
	var articleCategorys []common.ArticleCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&articleCategorys).Error
	return articleCategorys, total, err
}
