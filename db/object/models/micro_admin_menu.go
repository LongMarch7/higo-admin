package models

type MicroAdminMenu struct {
	Id        int     `json:"id" xorm:"not null pk INT(10)"`
	ParentId  int     `json:"parent_id" xorm:"not null default 0 comment('父菜单id') index INT(10)"`
	Type      int     `json:"type" xorm:"not null default 1 comment('菜单类型:0-只作为菜单,1-有界面可访问菜单,2-无界面可访问菜单,3-有界面非显示菜单') TINYINT(3)"`
	Status    int     `json:"status" xorm:"not null default 0 comment('状态:1-显示,0-不显示') index TINYINT(3)"`
	ListOrder float32 `json:"list_order" xorm:"not null default 10000 comment('排序') FLOAT"`
	Url       string  `json:"url" xorm:"not null default '''' comment('路径') VARCHAR(50)"`
	Func      string  `json:"func" xorm:"not null default '''' comment('控制器函数') VARCHAR(50)"`
	Method    string  `json:"method" xorm:"not null default '''' comment('规则方法(大写)GET、POST、PUT、 PUT | GET') VARCHAR(100)"`
	Param     string  `json:"param" xorm:"not null default '''' comment('额外参数') VARCHAR(50)"`
	Name      string  `json:"name" xorm:"not null default '''' comment('菜单名称') VARCHAR(30)"`
	Icon      string  `json:"icon" xorm:"not null default '''' comment('菜单图标') VARCHAR(20)"`
	Remark    string  `json:"remark" xorm:"not null default '''' comment('备注') VARCHAR(255)"`
}

func (c MicroAdminMenu) TableName() string {
	return "micro_admin_menu"
}
