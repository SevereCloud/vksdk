package main

import (
	"log"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/api/params"
	"github.com/SevereCloud/vksdk/longpoll-bot"
	"github.com/SevereCloud/vksdk/object"
)

func main() {
	token := "<TOKEN>" // рекомендуется использовать os.Getenv("TOKEN")
	vk := api.Init(token)

	// Получаем информацию о группе
	group, err := vk.GroupsGetByID(api.Params{})
	if err != nil {
		log.Fatal(err)
	}

	// Иницилизируем longpoll
	lp, err := longpoll.Init(vk, group[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	// Событие нового сообщения
	lp.MessageNew(func(obj object.MessageNewObject, groupID int) {
		log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)

		if obj.Message.Text == "ping" {
			b := params.NewMessagesSendBuilder()
			b.Message("pong")
			b.RandomID(0)
			b.PeerID(obj.Message.PeerID)

			_, err := vk.MessagesSend(b.Params)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	// Запускаем Bots Longpoll
	log.Println("Start longpoll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}