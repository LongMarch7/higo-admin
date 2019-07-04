package models

type MicroUserBalanceLog201904 struct {
	Id          int64  `json:"id" xorm:"pk autoincr BIGINT(20)"`
	UserId      int64  `json:"user_id" xorm:"not null default 0 comment('用户 id') BIGINT(20)"`
	CreateTime  int    `json:"create_time" xorm:"not null default 0 comment('创建时间') INT(10)"`
	Change      string `json:"change" xorm:"not null default 0.00 comment('更改余额') DECIMAL(16,2)"`
	Balance     string `json:"balance" xorm:"not null default 0.00 comment('更改后余额') DECIMAL(16,2)"`
	Fee         string `json:"fee" xorm:"not null default 0.00 comment('手续费') DECIMAL(16,2)"`
	Description string `json:"description" xorm:"not null default '''' comment('描述') VARCHAR(255)"`
	Remark      string `json:"remark" xorm:"not null default '''' comment('备注') VARCHAR(255)"`
}

func (c MicroUserBalanceLog201904) TableName() string {
	return "micro_user_balance_log_201904"
}
