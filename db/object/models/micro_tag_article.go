package models

type MicroTagArticle struct {
	Id        int64 `json:"id" xorm:"pk autoincr BIGINT(20)"`
	TagId     int64 `json:"tag_id" xorm:"not null default 0 comment('标签 id') BIGINT(20)"`
	ArticleId int64 `json:"article_id" xorm:"not null default 0 comment('文章 id') index BIGINT(20)"`
}

func (c MicroTagArticle) TableName() string {
	return "micro_tag_article"
}
