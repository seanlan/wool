package models

// ImApplication [...]
type ImApplication struct {
	ID          uint64 `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"id"`
	Appkey      string `gorm:"unique;column:appkey;type:varchar(128);not null" json:"appkey"`      // APPKEY
	Appsecret   string `gorm:"column:appsecret;type:varchar(255);not null" json:"appsecret"`       // APPSECRET
	NotifyURL   string `gorm:"column:notify_url;type:varchar(512);not null" json:"notify_url"`     // 消息通知地址
	WhitelistIP string `gorm:"column:whitelist_ip;type:varchar(512);not null" json:"whitelist_ip"` // 白名单ID逗号分隔
	CreateAt    int64  `gorm:"column:create_at;type:bigint;not null" json:"create_at"`             // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *ImApplication) TableName() string {
	return "im_application"
}

// ImChatHistory [...]
type ImChatHistory struct {
	ID             uint   `gorm:"primary_key;column:id;type:int unsigned;not null" json:"id"`
	Appkey         string `gorm:"index:appkey;column:appkey;type:varchar(128);not null" json:"appkey"`                            // 应用appkey
	MsgID          string `gorm:"unique;column:msg_id;type:varchar(128);not null" json:"msg_id"`                                  // 消息ID
	ConversationID string `gorm:"index:conversation_id;column:conversation_id;type:varchar(128);not null" json:"conversation_id"` // 会话ID
	FromTag        string `gorm:"index:from_tag;column:from_tag;type:varchar(128);not null" json:"from_tag"`                      // 发送方标示
	ToTag          string `gorm:"column:to_tag;type:varchar(128);not null" json:"to_tag"`                                         // 接收方标示
	Event          int    `gorm:"column:event;type:int;not null" json:"event"`                                                    // 消息类型 1、普通消息 2、群组消息 3、入群 4、退群
	Content        string `gorm:"column:content;type:varchar(512);not null" json:"content"`                                       // 消息内容json自定义
	SendTime       int64  `gorm:"column:send_time;type:bigint;not null" json:"send_time"`                                         // 消息发送时间
}

// TableName get sql table name.获取数据库表名
func (m *ImChatHistory) TableName() string {
	return "im_chat_history"
}

// ImConversation [...]
type ImConversation struct {
	ID             uint   `gorm:"primary_key;column:id;type:int unsigned;not null" json:"id"`
	Appkey         string `gorm:"index:appkey;column:appkey;type:varchar(128);not null" json:"appkey"`             // 应用appkey
	ConversationID string `gorm:"unique;column:conversation_id;type:varchar(128);not null" json:"conversation_id"` // 会话ID
	ChatType       int    `gorm:"column:chat_type;type:int;not null" json:"chat_type"`                             // 会话类型 1单聊 2群聊
	TargetTag      string `gorm:"column:target_tag;type:varchar(128);not null" json:"target_tag"`                  // 目标标示
	CreateUserTag  string `gorm:"column:create_user_tag;type:varchar(128);not null" json:"create_user_tag"`        // 创建会话的用户标示
	Icon           string `gorm:"column:icon;type:varchar(128);not null" json:"icon"`                              // 会话的图标
	Title          string `gorm:"column:title;type:varchar(128);not null" json:"title"`                            // 会话的title
	CreateAt       int64  `gorm:"column:create_at;type:bigint;not null" json:"create_at"`                          // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *ImConversation) TableName() string {
	return "im_conversation"
}

// ImUserConversation [...]
type ImUserConversation struct {
	ID             uint   `gorm:"primary_key;column:id;type:int unsigned;not null" json:"id"`
	Appkey         string `gorm:"unique;column:appkey;type:varchar(128);not null" json:"appkey"`            // 应用appkey
	ConversationID string `gorm:"column:conversation_id;type:varchar(128);not null" json:"conversation_id"` // 会话ID
	UserTag        string `gorm:"column:user_tag;type:varchar(128);not null" json:"user_tag"`               // 用户标示
	UniqueCode     string `gorm:"column:unique_code;type:varchar(128);not null" json:"unique_code"`         // 唯一标示
	UnreadCount    int    `gorm:"column:unread_count;type:int;not null" json:"unread_count"`                // 未读消息数量
	UpdateAt       int64  `gorm:"column:update_at;type:bigint;not null" json:"update_at"`                   // 最好更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ImUserConversation) TableName() string {
	return "im_user_conversation"
}
