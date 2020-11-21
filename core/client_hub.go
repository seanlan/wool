package core

import (
	"encoding/json"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/seanlan/packages/logging"
)

type ClientHub struct {
	// clients Map.
	clients map[string]*hashset.Set
	// websocket收到的消息处理
	queueBuff chan QueueMessage
}

var WSClientHub *ClientHub

func InitHub() {
	WSClientHub = &ClientHub{
		queueBuff: make(chan QueueMessage),
		clients:            make(map[string]*hashset.Set),
	}
	go WSClientHub.run()
}

// 加入新的连接
func (hub *ClientHub) join(wsClient *WSClient) {
	clientKey := wsClient.Key
	if _, ok := hub.clients[clientKey]; !ok {
		hub.clients[clientKey] = hashset.New()
	}
	hub.clients[clientKey].Add(wsClient)
}

// 端开连接
func (hub *ClientHub) draw(wsClient *WSClient) {
	clientKey := wsClient.Key
	if _, ok := hub.clients[clientKey]; ok {
		hub.clients[clientKey].Remove(wsClient)
		close(wsClient.SendPool)
	}
	logging.Logger.Debugf("hub clients: %v", hub.clients)
}

func (hub *ClientHub) run() {
	// 进行消息分发
	for {
		queueMsg, _ := <-hub.queueBuff
		clientKey := makeClientKey(queueMsg.AppKey, queueMsg.To)
		clientSet, ok := hub.clients[clientKey]
		message,_ := json.Marshal(queueMsg.WSMessage)
		if ok {
			for _, c := range clientSet.Values() {
				c.(*WSClient).SendPool <- message
			}
		}
	}
}
