package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"time"
)

type ArticleSearch struct {
	common.Article
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	PageInfo
}
