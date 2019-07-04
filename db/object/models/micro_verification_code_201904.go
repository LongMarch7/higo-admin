package models

type MicroVerificationCode201904 struct {
	Id         int64  `json:"id" xorm:"pk autoincr comment('表id') BIGINT(20)"`
	Count      int    `json:"count" xorm:"not null default 0 comment('当天已经发送成功的次数') INT(10)"`
	SendTime   int    `json:"send_time" xorm:"not null default 0 comment('最后发送成功时间') INT(10)"`
	ExpireTime int    `json:"expire_time" xorm:"not null default 0 comment('验证码过期时间') INT(10)"`
	Code       string `json:"code" xorm:"not null default '''' comment('最后发送成功的验证码') VARCHAR(8)"`
	Account    string `json:"account" xorm:"not null default '''' comment('手机号或者邮箱') VARCHAR(100)"`
}

func (c MicroVerificationCode201904) TableName() string {
	return "micro_verification_code_201904"
}
