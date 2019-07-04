package models

type MicroUserLoginAttempt struct {
	Id            int64  `json:"id" xorm:"pk autoincr BIGINT(20)"`
	UserId        int64  `json:"user_id" xorm:"not null default 0 comment('用户 id') BIGINT(20)"`
	LoginAttempts int    `json:"login_attempts" xorm:"not null default 0 comment('尝试次数') INT(10)"`
	AttemptTime   int    `json:"attempt_time" xorm:"not null default 0 comment('尝试登录时间') INT(10)"`
	LockedTime    int    `json:"locked_time" xorm:"not null default 0 comment('锁定时间') INT(10)"`
	Ip            string `json:"ip" xorm:"not null default '''' comment('用户 ip') VARCHAR(15)"`
}

func (c MicroUserLoginAttempt) TableName() string {
	return "micro_user_login_attempt"
}
