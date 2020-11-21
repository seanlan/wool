package models

// Application [...]
type Application struct {
	ID          uint64 `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"id"`
	Appkey      string `gorm:"unique;column:appkey;type:varchar(100);not null" json:"appkey"`      // APPKEY
	Appsecret   string `gorm:"column:appsecret;type:varchar(255);not null" json:"appsecret"`       // APPSECRET
	NotifyURL   string `gorm:"column:notify_url;type:varchar(512);not null" json:"notify_url"`     // 消息通知地址
	WhitelistIP string `gorm:"column:whitelist_ip;type:varchar(512);not null" json:"whitelist_ip"` // 白名单ID逗号分隔
	CreateAt    int64  `gorm:"column:create_at;type:bigint;not null" json:"create_at"`             // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *Application) TableName() string {
	return "application"
}
