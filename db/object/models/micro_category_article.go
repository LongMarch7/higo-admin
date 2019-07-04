package models

type MicroCategoryArticle struct {
	Id         int64   `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ArticleId  int64   `json:"article_id" xorm:"not null default 0 comment('文章id') BIGINT(20)"`
	CategoryId int64   `json:"category_id" xorm:"not null default 0 comment('分类id') index BIGINT(20)"`
	ListOrder  float32 `json:"list_order" xorm:"not null default 10000 comment('排序') FLOAT"`
}

func (c MicroCategoryArticle) TableName() string {
	return "micro_category_article"
}
