package models

type MicroRole struct {
	Id         int     `json:"id" xorm:"not null pk autoincr INT(10)"`
	ParentId   int     `json:"parent_id" xorm:"not null default 0 comment('父角色ID') index INT(10)"`
	RoleStatus int     `json:"role_status" xorm:"not null default 0 comment('状态:1-正常,2-禁用') index TINYINT(3)"`
	CreateTime int     `json:"create_time" xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int     `json:"update_time" xorm:"not null default 0 comment('更新时间') INT(10)"`
	ListOrder  float32 `json:"list_order" xorm:"not null default 0 comment('排序') FLOAT"`
	RoleName   string  `json:"role_name" xorm:"not null default '''' comment('角色名称') unique VARCHAR(20)"`
	Remark     string  `json:"remark" xorm:"not null default '''' comment('备注') VARCHAR(255)"`
}

func (c MicroRole) TableName() string {
	return "micro_role"
}
