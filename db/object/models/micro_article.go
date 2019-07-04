package models

type MicroArticle struct {
	Id               int64  `json:"id" xorm:"pk autoincr index(type_status_date) BIGINT(20)"`
	ArticleType      int    `json:"article_type" xorm:"not null default 1 comment('类型:1-文章,2-页面') index(type_status_date) TINYINT(3)"`
	ArticleFormat    int    `json:"article_format" xorm:"not null default 1 comment('内容格式:1-html,2-md') TINYINT(3)"`
	UserId           int64  `json:"user_id" xorm:"not null default 0 comment('发表者用户id') index BIGINT(20)"`
	ArticleStatus    int    `json:"article_status" xorm:"not null default 1 comment('状态:1-已发布,0-未发布') index(type_status_date) TINYINT(3)"`
	CommentStatus    int    `json:"comment_status" xorm:"not null default 1 comment('评论状态:1-允许,0-不允许') TINYINT(3)"`
	IsTop            int    `json:"is_top" xorm:"not null default 0 comment('是否置顶:1-置顶,2-不置顶') TINYINT(3)"`
	Recommended      int    `json:"recommended" xorm:"not null default 0 comment('是否推荐:1-推荐,2-不推荐') TINYINT(3)"`
	ArticleHits      int64  `json:"article_hits" xorm:"not null default 0 comment('查看数') BIGINT(20)"`
	ArticleFavorites int64  `json:"article_favorites" xorm:"not null default 0 comment('收藏数') BIGINT(20)"`
	ArticleLike      int64  `json:"article_like" xorm:"not null default 0 comment('点赞数') BIGINT(20)"`
	CommentCount     int64  `json:"comment_count" xorm:"not null default 0 comment('评论数') BIGINT(20)"`
	CreateTime       int    `json:"create_time" xorm:"not null default 0 comment('创建时间') index index(type_status_date) INT(10)"`
	UpdateTime       int    `json:"update_time" xorm:"not null default 0 comment('更新时间') INT(10)"`
	PublishedTime    int    `json:"published_time" xorm:"not null default 0 comment('发布时间') INT(10)"`
	DeleteTime       int    `json:"delete_time" xorm:"not null default 0 comment('删除时间') INT(10)"`
	ArticleTitle     string `json:"article_title" xorm:"not null default '''' comment('article') VARCHAR(100)"`
	ArticleKeywords  string `json:"article_keywords" xorm:"not null default '''' comment('seo keywords') VARCHAR(150)"`
	ArticleExcerpt   string `json:"article_excerpt" xorm:"not null default '''' comment('article') VARCHAR(500)"`
	Thumbnail        string `json:"thumbnail" xorm:"not null default '''' comment('缩略图') VARCHAR(100)"`
	ArticleContent   string `json:"article_content" xorm:"default 'NULL' comment('文章内容') TEXT"`
	More             string `json:"more" xorm:"default 'NULL' comment('扩展属性:如缩略图,格式为json') TEXT"`
}

func (c MicroArticle) TableName() string {
	return "micro_article"
}
