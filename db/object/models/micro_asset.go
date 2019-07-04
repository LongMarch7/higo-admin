package models

type MicroAsset struct {
	Id            int64  `json:"id" xorm:"pk autoincr BIGINT(20)"`
	UserId        int64  `json:"user_id" xorm:"not null default 0 comment('用户id') BIGINT(20)"`
	FileSize      int64  `json:"file_size" xorm:"not null default 0 comment('文件大小,单位B') BIGINT(20)"`
	CreateTime    int    `json:"create_time" xorm:"not null default 0 comment('上传时间') INT(10)"`
	Status        int    `json:"status" xorm:"not null default 1 comment('状态:1-可用,0-不可用') TINYINT(3)"`
	DownloadTimes int    `json:"download_times" xorm:"not null default 0 comment('下载次数') INT(10)"`
	FileKey       string `json:"file_key" xorm:"not null default '''' comment('文件惟一码') VARCHAR(64)"`
	Filename      string `json:"filename" xorm:"not null default '''' comment('文件名') VARCHAR(100)"`
	FilePath      string `json:"file_path" xorm:"not null default '''' comment('文件路径,相对于upload目录,可以为url') VARCHAR(100)"`
	FileMd5       string `json:"file_md5" xorm:"not null default '''' comment('文件md5值') VARCHAR(32)"`
	FileSha1      string `json:"file_sha1" xorm:"not null default '''' VARCHAR(40)"`
	Suffix        string `json:"suffix" xorm:"not null default '''' comment('文件后缀名,不包括点') VARCHAR(10)"`
	More          string `json:"more" xorm:"default 'NULL' comment('其它详细信息,JSON格式') TEXT"`
}

func (c MicroAsset) TableName() string {
	return "micro_asset"
}
