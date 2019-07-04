package models

type MicroUserFavorite struct {
	Id          int64  `json:"id" xorm:"pk autoincr BIGINT(20)"`
	UserId      int64  `json:"user_id" xorm:"not null default 0 comment('用户 id') index BIGINT(20)"`
	Title       string `json:"title" xorm:"not null default '''' comment('收藏内容的标题') VARCHAR(100)"`
	Thumbnail   string `json:"thumbnail" xorm:"not null default '''' comment('缩略图') VARCHAR(100)"`
	Url         string `json:"url" xorm:"default 'NULL' comment('收藏内容的地址，JSON格式') VARCHAR(255)"`
	Description string `json:"description" xorm:"default 'NULL' comment('收藏内容的描述') TEXT"`
	CreateTime  int    `json:"create_time" xorm:"default 0 comment('收藏时间') INT(10)"`
}

func (c MicroUserFavorite) TableName() string {
	return "micro_user_favorite"
}
