package models

type MicroUserActionLog201904 struct {
	Id            int64  `json:"id" xorm:"pk autoincr BIGINT(20)"`
	UserId        int64  `json:"user_id" xorm:"not null default 0 comment('用户id') index(user_parameter_action) index(user_parameter_action_ip) BIGINT(20)"`
	Count         int    `json:"count" xorm:"not null default 0 comment('访问次数') INT(10)"`
	LastVisitTime int    `json:"last_visit_time" xorm:"not null default 0 comment('最后访问时间') INT(10)"`
	Parameter     string `json:"parameter" xorm:"not null default '''' comment('参数') index(user_parameter_action) index(user_parameter_action_ip) VARCHAR(100)"`
	Action        string `json:"action" xorm:"not null default '''' comment('操作名称:url') index(user_parameter_action) index(user_parameter_action_ip) VARCHAR(50)"`
	Ip            string `json:"ip" xorm:"not null default '''' comment('用户ip') index(user_parameter_action_ip) VARCHAR(15)"`
}

func (c MicroUserActionLog201904) TableName() string {
	return "micro_user_action_log_201904"
}
