package tasks

import (
	"encoding/json"
	"github.com/seanlan/packages/db"
	"github.com/seanlan/packages/logging"
	"github.com/seanlan/wool/core"
	"github.com/seanlan/wool/models"
	"gorm.io/gorm"
)

func SaveChatMessage(message []byte) error {
	logging.Logger.Debug(message)
	var queueMsg core.QueueMessage
	var err error
	err = json.Unmarshal(message, &queueMsg)
	if err == nil {
		conversationID := core.MakeConversationID(queueMsg.AppKey, queueMsg.From, queueMsg.To, queueMsg.Event)
		chatMessage := models.ImChatHistory{
			Appkey:         queueMsg.AppKey,
			MsgID:          queueMsg.MsgID,
			ConversationID: conversationID,
			FromTag:        queueMsg.From,
			ToTag:          queueMsg.To,
			Event:          queueMsg.Event,
			Content:        queueMsg.Content,
			SendTime:       queueMsg.SendTime,
		}
		// 保存消息记录
		db.DB.Save(&chatMessage)
		// 更新未读消息条目
		db.DB.Model(&models.ImUserConversation{}).Where(
			"conversation_id = ? and user_tag!= ?",
			conversationID, queueMsg.From).Update(
			"unread_count", gorm.Expr("unread_count + ?", 1))
	}
	return nil
}
