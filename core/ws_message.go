package core

type WSMessage struct {
	From           string `json:"from,required"`             //发送者ID
	To             string `json:"to,required"`               //目标ID
	Event          int    `json:"event,required"`            //事件类型 1、普通消息 2、群组消息 3、入群 4、退群
	Content        string `json:"content,required"`          //消息体 自定义
	MsgID          string `json:"msg_id,omitempty"`          //消息ID
	SendTime       int64  `json:"send_time,omitempty"`       //消息产生时间
	ConversationID string `json:"conversation_id,omitempty"` //会话ID
}

type QueueMessage struct {
	WSMessage
	AppKey string `json:"appkey,required"` //应用key
}
