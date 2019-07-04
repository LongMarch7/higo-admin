package models

type MicroCasbinRule struct {
	PType string `json:"p_type" xorm:"not null default '''' comment('规则类型') unique(rule_key) VARCHAR(100)"`
	V0    string `json:"v0" xorm:"default 'NULL' comment('规则0') unique(rule_key) VARCHAR(100)"`
	V1    string `json:"v1" xorm:"default 'NULL' comment('规则1') unique(rule_key) VARCHAR(100)"`
	V2    string `json:"v2" xorm:"default 'NULL' comment('规则2') unique(rule_key) VARCHAR(100)"`
	V3    string `json:"v3" xorm:"default 'NULL' comment('规则3') unique(rule_key) VARCHAR(100)"`
	V4    string `json:"v4" xorm:"default 'NULL' comment('规则4') unique(rule_key) VARCHAR(100)"`
	V5    string `json:"v5" xorm:"default 'NULL' comment('规则4') unique(rule_key) VARCHAR(100)"`
}

func (c MicroCasbinRule) TableName() string {
	return "micro_casbin_rule"
}
