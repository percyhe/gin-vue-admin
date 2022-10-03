// 自动生成模板ArticleCategory
package common

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ArticleCategory 结构体
type ArticleCategory struct {
	global.GVA_MODEL
	Title string `json:"title" form:"title" gorm:"column:title;comment:标题;size:64;"`
	Sort  *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:3;default:1"`
	Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:100;"`
}

// TableName ArticleCategory 表名
func (ArticleCategory) TableName() string {
	return "article_category"
}
