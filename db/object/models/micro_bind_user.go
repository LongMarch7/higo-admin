package models

type MicroBindUser struct {
	Id         int64  `json:"id" xorm:"pk autoincr comment('id') BIGINT(20)"`
	UserId     string `json:"user_id" xorm:"not null comment('绑定账号') VARCHAR(20)"`
	Start      int    `json:"start" xorm:"not null default 2 comment('粉丝数') INT(11)"`
	RegionId   int    `json:"region_id" xorm:"not null default 0 comment('地域') INT(10)"`
	ClassId    int    `json:"class_id" xorm:"not null default 0 comment('分类') INT(10)"`
	BindStatus int    `json:"bind_status" xorm:"not null default 0 comment('绑定状态:1-已绑定,2-未绑定,3-审核中') TINYINT(3)"`
	PassTime   int    `json:"pass_time" xorm:"not null default 0 comment('自动审核通过日期') INT(11)"`
}

func (c MicroBindUser) TableName() string {
	return "micro_bind_user"
}
