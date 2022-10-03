package common

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	commonReq "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ArticleService struct {
}

func (articleService *ArticleService) GetArticleInfoList(info commonReq.ArticleSearch) (list []common.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&common.Article{})
	var articles []common.Article
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("create_time BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	//err = db.Debug().Model(&common.Article{}).Limit(limit).Offset(offset).Preload("Category").Preload("ArticleContent").Find(&articles).Error
	err = db.Debug().Model(&common.Article{}).Limit(limit).Offset(offset).Preload("Category").Find(&articles).Error

	return articles, total, err
}
