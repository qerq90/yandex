package client

import (
	"errors"
	"os"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
)

type VkClient struct {
	api *api.VK
}

func MakeVkClient() (*VkClient, error) {
	VK_TOKEN, exists := os.LookupEnv("VK_TOKEN")
	if !exists {
		return nil, errors.New("no VK_TOKEN found in .env file")
	}

	vk := api.NewVK(VK_TOKEN)

	return &VkClient{api: vk}, nil
}

func (vkClient *VkClient) SendMessage(text string, userId int, keyboard *object.MessagesKeyboard) error {
	p := params.NewMessagesSendBuilder()
	p.UserID(userId)
	p.RandomID(0)
	p.Message(text)
	if keyboard != nil {
		p.Keyboard(keyboard)
	}

	_, err := vkClient.api.MessagesSend(p.Params)
	return err
}
