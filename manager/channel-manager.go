package manager

import (
	"sync"
)

// ChannelManager .
type ChannelManager struct {
	ChannelMap map[string]chan string
}

var instance *ChannelManager
var once sync.Once

// GetChannelManageInstance .
func GetChannelManageInstance() *ChannelManager {
	once.Do(func() {
		instance = &ChannelManager{}
		instance.ChannelMap = make(map[string]chan string)
	})
	return instance
}

// StopChannel .
func (channelManager ChannelManager) StopChannel(key string) bool {
	channel := channelManager.ChannelMap[key]
	if channel == nil {
		return false
	}
	<-channel
	delete(channelManager.ChannelMap, key)
	return true
}

// AddChannel .
func (channelManager ChannelManager) AddChannel(key string, channel chan string) {
	channelManager.ChannelMap[key] = channel
}
