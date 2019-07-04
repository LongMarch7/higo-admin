package models

type MicroComment struct {
	Id           int64  `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ParentId     int64  `json:"parent_id" xorm:"not null default 0 comment('被回复的评论id') index BIGINT(20)"`
	UserId       int    `json:"user_id" xorm:"not null default 0 comment('发表评论的用户id') INT(10)"`
	ToUserId     int    `json:"to_user_id" xorm:"not null default 0 comment('被评论的用户id') INT(10)"`
	ObjectId     int    `json:"object_id" xorm:"not null default 0 comment('评论内容 id') index index(table_id_status) INT(10)"`
	LikeCount    int    `json:"like_count" xorm:"not null default 0 comment('点赞数') INT(10)"`
	DislikeCount int    `json:"dislike_count" xorm:"not null default 0 comment('吐槽数') INT(10)"`
	Floor        int    `json:"floor" xorm:"not null default 0 comment('楼层数') INT(10)"`
	CreateTime   int    `json:"create_time" xorm:"not null default 0 comment('评论时间') index INT(10)"`
	DeleteTime   int    `json:"delete_time" xorm:"not null default 0 comment('删除时间') INT(10)"`
	Status       int    `json:"status" xorm:"not null default 1 comment('状态:1-已审核,0-未审核') index index(table_id_status) TINYINT(3)"`
	Type         int    `json:"type" xorm:"not null default 1 comment('评论类型:1-实名评论') TINYINT(3)"`
	TabName      string `json:"tab_name" xorm:"not null default '''' comment('评论内容所在表,不带表前缀') index(table_id_status) VARCHAR(64)"`
	FullName     string `json:"full_name" xorm:"not null default '''' comment('评论者昵称') VARCHAR(50)"`
	Url          string `json:"url" xorm:"default 'NULL' comment('原文地址') TEXT"`
	Content      string `json:"content" xorm:"default 'NULL' comment('评论内容') TEXT"`
	More         string `json:"more" xorm:"default 'NULL' comment('扩展属性') TEXT"`
}

func (c MicroComment) TableName() string {
	return "micro_comment"
}
