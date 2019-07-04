package models

type MicroTaskCheckLog201904 struct {
	Id             int64  `json:"id" xorm:"pk autoincr comment('id') BIGINT(20)"`
	UserId         int64  `json:"user_id" xorm:"not null default 1 comment('用户id') BIGINT(20)"`
	TaskId         int64  `json:"task_id" xorm:"not null default 1 comment('任务id') BIGINT(20)"`
	ReasonDescribe string `json:"reason_describe" xorm:"not null default '''' comment('原因描述') VARCHAR(200)"`
}

func (c MicroTaskCheckLog201904) TableName() string {
	return "micro_task_check_log_201904"
}
