package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"time"
)

type ArticleCategorySearch struct {
	common.ArticleCategory
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	PageInfo
}
