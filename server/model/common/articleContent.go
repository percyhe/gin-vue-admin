package common

type ArticleContent struct {
	Id        int64  `gorm:"primary_key;column:id" json:"id"`
	ArticleId int64  `gorm:"column:article_id;index:idx_article_id" json:"article_id"`
	Content   string `gorm:"column:content;type:longtext" json:"contnet"`
}

func (ArticleContent) TableName() string {
	return "article_content"
}
