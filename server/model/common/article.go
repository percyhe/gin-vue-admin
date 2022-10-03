package common

import (
	"gorm.io/gorm"
	"time"
)

//文章
type Article struct {
	ID             int64           `gorm:"primary_key":"column:id" json:"id"`
	CategoryID     int64           `gorm:"column:category_id;index:idx_category_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"" json:"category_id"`
	Summary        string          `gorm:"column:summary" json:"summary"`
	Title          string          `gorm:"column:title" json:"title"`
	ViewCount      uint32          `gorm:"column:view_count" json:"view_count"`
	CreateTime     int             `gorm:"column:create_time" json:"create_time"`
	UpdateTime     int             `gorm:"column:update_time" json:"update_time"`
	CommentCount   uint32          `gorm:"column:comment_count" json:"comment_count"`
	UserId         int64           `gorm:"column:user_id" json:"user_id"`
	ArticleContent ArticleContent  `gorm:"FOREIGNKEY:ArticleId;has2one:ArticleContent""`
	IsPublish      int             `gorm:"column:is_publish;default:0" json:"is_publish"`
	Status         int             `gorm:"column:status;default:1" json:"status"`
	Category       ArticleCategory `gorm:"foreignkey:CategoryID"`
	Content        string          `gorm:"-" json:"content"`
}

func (Article) TableName() string {
	return "article_v2"
}

//创建钩子
func (article *Article) BeforeCreate(tx *gorm.DB) (err error) {
	article.CreateTime = int(time.Now().Unix())
	article.UpdateTime = int(time.Now().Unix())
	return
}

//更新hook
func (article *Article) BeforeUpdate(tx *gorm.DB) (err error) {
	article.UpdateTime = int(time.Now().Unix())
	return
}
