package models

type MicroOption struct {
	Id          int64  `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Autoload    int    `json:"autoload" xorm:"not null default 1 comment('是否自动加载:1-自动加载,0-不自动加载') TINYINT(3)"`
	OptionName  string `json:"option_name" xorm:"not null default '''' comment('配置名') unique VARCHAR(64)"`
	OptionValue string `json:"option_value" xorm:"default 'NULL' comment('配置值') LONGTEXT"`
}

func (c MicroOption) TableName() string {
	return "micro_option"
}
