package models

type MicroTaskList201904 struct {
	Id           int64   `json:"id" xorm:"pk autoincr comment('id') BIGINT(20)"`
	UserId       int64   `json:"user_id" xorm:"not null default 1 comment('发布者id') BIGINT(20)"`
	ClassId      int64   `json:"class_id" xorm:"not null default 1 comment('分类id') BIGINT(20)"`
	Count        int     `json:"count" xorm:"not null default 1 comment('任务数量') INT(15)"`
	ConsumeCount int     `json:"consume_count" xorm:"not null default 0 comment('已接单任务数量') INT(15)"`
	CheckCount   int     `json:"check_count" xorm:"not null default 0 comment('提交审核任务数量') INT(15)"`
	FinishCount  int     `json:"finish_count" xorm:"not null default 0 comment('完成任务数量') INT(15)"`
	Balance      string  `json:"balance" xorm:"not null default 0.00 comment('任务单价') DECIMAL(16,2)"`
	CreateTime   int     `json:"create_time" xorm:"not null default 0 comment('发布日期') INT(11)"`
	EndTime      int     `json:"end_time" xorm:"not null default 0 comment('结束日期') INT(11)"`
	ListOrder    float32 `json:"list_order" xorm:"not null default 10000 comment('排序') FLOAT"`
	RegionRule   int64   `json:"region_rule" xorm:"not null default 0 comment('地域id限制') BIGINT(10)"`
	ClassRule    int     `json:"class_rule" xorm:"not null default 0 comment('类别限制:0-不限制,1-同类别限制') TINYINT(2)"`
	Status       int     `json:"status" xorm:"not null default 0 comment('发布状态:0-未发布,1-审核中,2-已发布') TINYINT(3)"`
	TaskTitle    string  `json:"task_title" xorm:"not null default '''' comment('任务标题') VARCHAR(100)"`
	TaskDescribe string  `json:"task_describe" xorm:"not null default '''' comment('任务简述') VARCHAR(200)"`
	Thumbnail    string  `json:"thumbnail" xorm:"not null default '''' comment('缩略图') VARCHAR(100)"`
	TaskContent  string  `json:"task_content" xorm:"default 'NULL' comment('任务内容') TEXT"`
}

func (c MicroTaskList201904) TableName() string {
	return "micro_task_list_201904"
}
