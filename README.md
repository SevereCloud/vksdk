# VK SDK for Golang

[![Build Status](https://travis-ci.com/SevereCloud/vksdk.svg?branch=master)](https://travis-ci.com/SevereCloud/vksdk)
[![Documentation](https://godoc.org/github.com/SevereCloud/vksdk?status.svg)](https://pkg.go.dev/github.com/SevereCloud/vksdk)
[![VK Developers](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.com/dev/)
[![codecov](https://codecov.io/gh/SevereCloud/vksdk/branch/master/graph/badge.svg)](https://codecov.io/gh/SevereCloud/vksdk)
[![VK chat](https://img.shields.io/badge/chat-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.me/join/AJQ1d6Or8Q00Y_CSOESfbqGt)
[![release](https://img.shields.io/github/v/tag/SevereCloud/vksdk?label=release)](https://github.com/SevereCloud/vksdk/releases)
[![license](https://img.shields.io/github/license/SevereCloud/vksdk.svg?maxAge=2592000)](https://github.com/SevereCloud/vksdk/blob/master/LICENSE)

**VK SDK for Golang** готовая реализация основных функций VK API для языка Go.

[English documentation](https://pkg.go.dev/github.com/SevereCloud/vksdk)

### Возможности

- [API](https://github.com/SevereCloud/vksdk/tree/master/api#api)
  - Возвращает готовые структуры
  - Реализовано 400+ методов
	- Возможность изменять обработчик запросов
  - Возможность изменять HTTP клиент
  - Ограничитель запросов
  - Загрузка файлов
- [Callback API](https://github.com/SevereCloud/vksdk/tree/master/callback#callback-api)
  - Поддерживает все события
  - Возвращает готовые структуры
- [Bots Long Poll API](https://github.com/SevereCloud/vksdk/tree/master/longpoll-bot#bots-long-poll-api)
  - Поддерживает все события
  - Возвращает готовые структуры
  - Возможность изменять HTTP клиент
- [User Long Poll API](https://github.com/SevereCloud/vksdk/tree/master/longpoll-user#user-long-poll-api)
  - ~~Возвращает готовые структуры~~ [#44](https://github.com/SevereCloud/vksdk/issues/44)
  - Возможность изменять HTTP клиент
- ~~Streaming api~~ [#3](https://github.com/SevereCloud/vksdk/issues/3)
- [FOAF](https://github.com/SevereCloud/vksdk/tree/master/foaf#foaf)
  - Работает с пользователями и группами
  - Возвращает готовые структуры
- [VK Mini Apps](https://github.com/SevereCloud/vksdk/tree/master/vkapps#vk-mini-apps)
  - Проверка параметров запуска
  - Промежуточный http обработчик

### Установка

```shell
# go mod init mymodulename
go get github.com/SevereCloud/vksdk@latest
```

### Пример

```go
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
```

### Лицензия

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSevereCloud%2Fvksdk.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FSevereCloud%2Fvksdk?ref=badge_large)