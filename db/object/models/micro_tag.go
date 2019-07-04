package models

type MicroTag struct {
	Id          int64  `json:"id" xorm:"pk autoincr comment('标签id') BIGINT(20)"`
	Status      int    `json:"status" xorm:"not null default 1 comment('状态:1-发布,2-不发布') TINYINT(3)"`
	Recommended int    `json:"recommended" xorm:"not null default 0 comment('是否推荐:1-推荐,2-不推荐') TINYINT(3)"`
	CreateTime  int    `json:"create_time" xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime  int    `json:"update_time" xorm:"not null default 0 comment('更新时间') INT(10)"`
	Name        string `json:"name" xorm:"not null default '''' comment('标签名称') unique VARCHAR(20)"`
}

func (c MicroTag) TableName() string {
	return "micro_tag"
}
