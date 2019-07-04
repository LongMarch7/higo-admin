package models

type MicroLink struct {
	Id          int64   `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Status      int     `json:"status" xorm:"not null default 1 comment('状态:1-显示,2-不显示') index TINYINT(3)"`
	Rating      int     `json:"rating" xorm:"not null default 0 comment('友情链接评级') INT(11)"`
	ListOrder   float32 `json:"list_order" xorm:"not null default 10000 comment('排序') FLOAT"`
	CreateTime  int     `json:"create_time" xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime  int     `json:"update_time" xorm:"not null default 0 comment('更新时间') INT(10)"`
	Description string  `json:"description" xorm:"not null default '''' comment('友情链接描述') VARCHAR(255)"`
	Url         string  `json:"url" xorm:"not null default '''' comment('友情链接地址') VARCHAR(255)"`
	Name        string  `json:"name" xorm:"not null default '''' comment('友情链接名称') VARCHAR(30)"`
	Image       string  `json:"image" xorm:"not null default '''' comment('友情链接图标') VARCHAR(100)"`
	Target      string  `json:"target" xorm:"not null default '''' comment('友情链接打开方式') VARCHAR(10)"`
	Rel         string  `json:"rel" xorm:"not null default '''' comment('链接与网站的关系') VARCHAR(50)"`
}

func (c MicroLink) TableName() string {
	return "micro_link"
}
