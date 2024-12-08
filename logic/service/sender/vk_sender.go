package sender

import (
	"qerq90/yandex/logic/client"
	"qerq90/yandex/logic/db"
)

type VkSender struct {
	vkClient *client.VkClient
	dao      *db.Dao
}

func MakeVkSender(vk *client.VkClient, dao *db.Dao) *VkSender {
	return &VkSender{vk, dao}
}

func (sender *VkSender) Send(id int, message string) {
	// vkId := sender.dao.GetVkId(id) //TODO use dao
	sender.vkClient.SendMessage(message, id, nil)
}
