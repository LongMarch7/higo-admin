package models

type MicroCategory struct {
	Id             int64   `json:"id" xorm:"pk autoincr comment('分类id') BIGINT(20)"`
	ParentId       int64   `json:"parent_id" xorm:"not null default 0 comment('分类父id') BIGINT(20)"`
	Level          int     `json:"level" xorm:"not null default 1 comment('级别:1-一级,2-二级,3-三级') TINYINT(3)"`
	PostCount      int64   `json:"post_count" xorm:"not null default 0 comment('分类文章数') BIGINT(20)"`
	Status         int     `json:"status" xorm:"not null default 1 comment('状态:1-发布,2-不发布') TINYINT(3)"`
	DeleteTime     int     `json:"delete_time" xorm:"not null default 0 comment('删除时间') INT(10)"`
	ListOrder      float32 `json:"list_order" xorm:"not null default 10000 comment('排序') FLOAT"`
	Name           string  `json:"name" xorm:"not null default '''' comment('分类名称') VARCHAR(200)"`
	Description    string  `json:"description" xorm:"not null default '''' comment('分类描述') VARCHAR(255)"`
	SeoTitle       string  `json:"seo_title" xorm:"not null default '''' VARCHAR(100)"`
	SeoKeywords    string  `json:"seo_keywords" xorm:"not null default '''' VARCHAR(255)"`
	SeoDescription string  `json:"seo_description" xorm:"not null default '''' VARCHAR(255)"`
	CreateTime     int     `json:"create_time" xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime     int     `json:"update_time" xorm:"not null default 0 comment('更新时间') INT(10)"`
	OneTpl         string  `json:"one_tpl" xorm:"not null default '''' comment('分类文章页模板') VARCHAR(50)"`
	More           string  `json:"more" xorm:"default 'NULL' comment('扩展属性') TEXT"`
}

func (c MicroCategory) TableName() string {
	return "micro_category"
}
