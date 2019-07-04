package models

type MicroUserLike201904 struct {
	Id          int64  `json:"id" xorm:"pk autoincr BIGINT(20)"`
	UserId      int64  `json:"user_id" xorm:"not null default 0 comment('用户 id') index BIGINT(20)"`
	ObjectId    int    `json:"object_id" xorm:"not null default 0 comment('内容原来的主键id') INT(10)"`
	CreateTime  int    `json:"create_time" xorm:"not null default 0 comment('创建时间') INT(10)"`
	Url         string `json:"url" xorm:"not null default '''' comment('内容的地址') VARCHAR(255)"`
	Title       string `json:"title" xorm:"not null default '''' comment('内容的标题') VARCHAR(100)"`
	Thumbnail   string `json:"thumbnail" xorm:"not null default '''' comment('缩略图') VARCHAR(100)"`
	Description string `json:"description" xorm:"default 'NULL' comment('内容的描述') TEXT"`
}

func (c MicroUserLike201904) TableName() string {
	return "micro_user_like_201904"
}
