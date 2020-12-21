package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/seanlan/packages/db"
	"github.com/seanlan/wool/core"
	"github.com/seanlan/wool/models"
	"time"
)

// CreateSingleConversation 创建单聊会话
func CreateSingleConversation(appKey,from,to,icon,title string) string {
	conversationID := core.MakeConversationID(appKey, from, to, 1)
	now := time.Now().Unix()
	tx := db.DB.Begin()
	conversation := models.ImConversation{
		Appkey:         appKey,
		ConversationID: conversationID,
		ChatType:       1,
		TargetTag:      to,
		CreateUserTag:  from,
		Icon:           icon,
		Title:          title,
		CreateAt:       now,
	}
	tx.Save(&conversation)
	fromUniqueCode := hex.EncodeToString(md5.New().Sum([]byte(fmt.Sprintf("%s:%s",conversationID,from))))
	fromConversationRecord := &models.ImUserConversation{
		Appkey:         appKey,
		ConversationID: conversationID,
		UserTag:        from,
		UniqueCode:		fromUniqueCode,
		UnreadCount:    0,
		UpdateAt:       now,
	}
	tx.Save(&fromConversationRecord)
	toUniqueCode := hex.EncodeToString(md5.New().Sum([]byte(fmt.Sprintf("%s:%s",conversationID,to))))
	toConversationRecord := &models.ImUserConversation{
		Appkey:         appKey,
		ConversationID: conversationID,
		UserTag:        to,
		UniqueCode:		toUniqueCode,
		UnreadCount:    0,
		UpdateAt:       now,
	}
	tx.Save(&toConversationRecord)
	tx.Commit()
	return conversationID
}