package models

type MicroUser struct {
	Id                int64  `json:"id" xorm:"pk autoincr BIGINT(20)"`
	BindId            string `json:"bind_id" xorm:"not null default '''' comment('绑定第三方账号ID') VARCHAR(32)"`
	UserType          int    `json:"user_type" xorm:"not null default 2 comment('用户类型:1-admin,2-star,3-ad') TINYINT(3)"`
	Sex               int    `json:"sex" xorm:"not null default 1 comment('性别:1-保密,2-男,3-女') TINYINT(2)"`
	Birthday          int    `json:"birthday" xorm:"not null default 0 comment('生日') INT(11)"`
	LastLoginTime     int    `json:"last_login_time" xorm:"not null default 0 comment('最后登录时间') INT(11)"`
	Score             int    `json:"score" xorm:"not null default 0 comment('用户等级') INT(11)"`
	Coin              int    `json:"coin" xorm:"not null default 0 comment('金币') INT(10)"`
	Balance           string `json:"balance" xorm:"not null default 0.00 comment('余额') DECIMAL(16,2)"`
	CreateTime        int    `json:"create_time" xorm:"not null default 0 comment('注册时间') INT(11)"`
	UpdateTime        int    `json:"update_time" xorm:"not null default 0 comment('更新时间') INT(11)"`
	UserStatus        int    `json:"user_status" xorm:"not null default 3 comment('用户状态:1-正常,2-冻结,3-禁用') TINYINT(3)"`
	UserLogin         string `json:"user_login" xorm:"not null default '''' comment('用户名') unique VARCHAR(32)"`
	UserPass          string `json:"user_pass" xorm:"not null default '''' comment('登录密码:micro_password加密') VARCHAR(64)"`
	PayPass           string `json:"pay_pass" xorm:"not null default '''' comment('支付密码:micro_password加密') VARCHAR(64)"`
	FreezeTime        int    `json:"freeze_time" xorm:"not null default 0 comment('冻结时间') INT(11)"`
	UserNickname      string `json:"user_nickname" xorm:"not null default '''' comment('用户昵称') VARCHAR(50)"`
	UserEmail         string `json:"user_email" xorm:"not null default '''' comment('用户登录邮箱') VARCHAR(100)"`
	Avatar            string `json:"avatar" xorm:"not null default '''' comment('用户头像') VARCHAR(255)"`
	Signature         string `json:"signature" xorm:"not null default '''' comment('个性签名') VARCHAR(255)"`
	LastLoginIp       string `json:"last_login_ip" xorm:"not null default '''' comment('最后登录ip') VARCHAR(15)"`
	UserActivationKey string `json:"user_activation_key" xorm:"not null default '''' comment('激活码') VARCHAR(60)"`
	Mobile            string `json:"mobile" xorm:"not null default '''' comment('中国手机不带国家代码.国际手机号格式为：国家代码-手机号') VARCHAR(20)"`
	More              string `json:"more" xorm:"default 'NULL' comment('扩展属性') TEXT"`
}

func (c MicroUser) TableName() string {
	return "micro_user"
}
