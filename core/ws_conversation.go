package core

type ChatType int64

// 会话类型
const (
	SingleChat ChatType = 1 //单聊
	GroupChat  ChatType = 2 //群聊
)

type WSConversation struct {
	Key      string   `json:"key,required"`       //会话key
	AppKey   string   `json:"appkey,required"`    //应用ID
	ChatType ChatType `json:"chat_type,required"` //会话类型
}
