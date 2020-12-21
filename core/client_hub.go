package core

import (
	"encoding/json"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/seanlan/packages/logging"
	"github.com/seanlan/packages/task_queue"
)

type ClientHub struct {
	// clients Map.
	clients map[string]*hashset.Set
	// websocket收到的消息处理
	queueBuff chan QueueMessage
}

var WSClientHub *ClientHub

// InitHub 初始化Hub
func InitHub() {
	WSClientHub = &ClientHub{
		queueBuff: make(chan QueueMessage),
		clients:   make(map[string]*hashset.Set),
	}
	go WSClientHub.run()
}

// join 加入新的连接
func (hub *ClientHub) join(wsClient *WSClient) {
	clientKey := wsClient.Key
	if _, ok := hub.clients[clientKey]; !ok {
		hub.clients[clientKey] = hashset.New()
	}
	hub.clients[clientKey].Add(wsClient)
}

// drop 断开连接
func (hub *ClientHub) drop(wsClient *WSClient) {
	clientKey := wsClient.Key
	if _, ok := hub.clients[clientKey]; ok {
		hub.clients[clientKey].Remove(wsClient)
		close(wsClient.SendPool)
	}
	logging.Logger.Debugf("hub clients: %v", hub.clients)
}

// distribute 进行消息分发
func (hub *ClientHub) distribute() {
	for {
		queueMsg, _ := <-hub.queueBuff
		clientKey := makeClientKey(queueMsg.AppKey, queueMsg.To)
		clientSet, ok := hub.clients[clientKey]
		message, _ := json.Marshal(queueMsg.WSMessage)
		if ok {
			for _, c := range clientSet.Values() {
				c.(*WSClient).SendMsg(message)
			}
		}
	}
}

// 放入消息
func (hub *ClientHub) putMsg(queueMsg QueueMessage) {
	hub.queueBuff <- queueMsg
	msgBytes, err := json.Marshal(queueMsg)
	if err == nil {
		task_queue.SendTask("SaveMessage",
			[]tasks.Arg{
				{
					Name:  "message",
					Type:  "[]byte",
					Value: msgBytes,
				},
			}, "", 0)
	}

}

// run
func (hub *ClientHub) run() {
	go hub.distribute()
}
