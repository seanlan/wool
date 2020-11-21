package core

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/seanlan/packages/logging"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second    //心跳60秒一次
	pingPeriod     = (pongWait * 9) / 10 //定时发送ping消息的间隔时间
	maxMessageSize = 1024 * 2            //允许消息最大长度
)

type WSClient struct {
	Key      string          //连接标识
	AppKey   string          //所属应用
	Tag      string          //所属应用的标识
	conn     *websocket.Conn //websocket连接
	SendPool chan []byte     //发送的消息缓存池
	hub      *ClientHub      //client管理对象
}

//开始读取消息
func (wc *WSClient) readPump() {
	//关闭链接
	defer func() {
		wc.hub.draw(wc)
		_ = wc.conn.Close()
	}()
	//设置消息最大长度
	wc.conn.SetReadLimit(maxMessageSize)
	//设置读取失败时间
	_ = wc.conn.SetReadDeadline(time.Now().Add(pongWait))
	//设置心跳
	wc.conn.SetPongHandler(
		func(string) error {
			_ = wc.conn.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		})
	for {
		_, message, err := wc.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logging.Logger.Errorf("error: %v", err)
			}
			break
		}
		logging.Logger.Debugf("received messages: %s", string(message))
		//消息重新封装
		var wsMsg WSMessage
		err = json.Unmarshal(message, &wsMsg)
		if err != nil {
			// 消息格式解析失败
			logging.Logger.Debugf("消息解析失败:%s", message)
			continue
		}
		wsMsg.From = wc.Tag
		var qMsg = QueueMessage{
			WSMessage: wsMsg,
			AppKey:    wc.AppKey,
		}
		//将消息发送到hub进行分发
		wc.hub.queueBuff <- qMsg
	}
	logging.Logger.Debug("ws client readPump stop !!")
}

//开始发送消息
func (wc *WSClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = wc.conn.Close()
	}()
	for {
		select {
		case message, ok := <-wc.SendPool:
			_ = wc.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// channel关闭
				logging.Logger.Debug("SendPool closed!!")
				_ = wc.conn.WriteMessage(websocket.CloseMessage, []byte{})
				goto ForEnd
			}
			w, err := wc.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				goto ForEnd
			}
			_, _ = w.Write(message)
			n := len(wc.SendPool)
			for i := 0; i < n; i++ {
				_, err = w.Write(<-wc.SendPool)
			}
			if err := w.Close(); err != nil {
				goto ForEnd
			}
		case <-ticker.C:
			_ = wc.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := wc.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				goto ForEnd
			}
		}
	}
ForEnd:
	logging.Logger.Debug("ws client writePump stop!!")
}

func makeClientKey(appKey, tag string) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s:%s", appKey, tag)))
	return hex.EncodeToString(h.Sum(nil))
}

func newWSClient(appKey, tag string, conn *websocket.Conn) *WSClient {
	wsClient := &WSClient{
		Key:      makeClientKey(appKey, tag),
		AppKey:   appKey,
		Tag:      tag,
		conn:     conn,
		SendPool: make(chan []byte),
		hub:      WSClientHub,
	}
	go wsClient.readPump()
	go wsClient.writePump()
	return wsClient
}
