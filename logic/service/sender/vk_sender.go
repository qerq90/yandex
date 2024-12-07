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

func (sender *VkSender) send(id int, message string) {
	vkId := sender.dao.GetVkId(id)
	sender.vkClient.SendMessage(message, vkId, nil)
}
