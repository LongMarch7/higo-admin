package models

type MicroTaskAction201904 struct {
	Id           int64  `json:"id" xorm:"pk autoincr comment('id') BIGINT(20)"`
	TaskId       int64  `json:"task_id" xorm:"not null default 1 comment('任务id') BIGINT(20)"`
	UserId       int64  `json:"user_id" xorm:"not null default 1 comment('用户id') BIGINT(20)"`
	RegionId     int    `json:"region_id" xorm:"not null default 1 comment('用户地域') INT(10)"`
	ClassId      int    `json:"class_id" xorm:"not null default 1 comment('用户分类') INT(10)"`
	Balance      string `json:"balance" xorm:"not null default 0.00 comment('任务单价') DECIMAL(16,2)"`
	CreateTime   int    `json:"create_time" xorm:"not null default 0 comment('接单日期') INT(11)"`
	EndTime      int    `json:"end_time" xorm:"not null default 0 comment('结束日期') INT(11)"`
	Status       int    `json:"status" xorm:"not null default 0 comment('状态:0-未完成,1-自动审核中,2-已完成,3-复审,4-无效') TINYINT(3)"`
	CheckCount   int    `json:"check_count" xorm:"not null default 0 comment('复审次数') TINYINT(10)"`
	QrUrl        string `json:"qr_url" xorm:"not null default ''noset'' comment('二维码路径') VARCHAR(100)"`
	CommentLevel int    `json:"comment_level" xorm:"not null default 4 comment('评论等级') TINYINT(3)"`
	CommentText  string `json:"comment_text" xorm:"default 'NULL' comment('评论内容') TEXT"`
}

func (c MicroTaskAction201904) TableName() string {
	return "micro_task_action_201904"
}
