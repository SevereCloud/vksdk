# VK API for Golang 5.92

Внимание - этот репозиторий в **очень ранней разработке**.

## Установка

```shell
go get -u github.com/severecloud/vkapi
```

## Модули

- [API](https://github.com/SevereCloud/vksdk/tree/master/5.92/api)
- [Callback](https://github.com/SevereCloud/vksdk/tree/master/5.92/callback)

## Пример

```go
package main

import (
	"fmt"
	"log"

	"github.com/severecloud/vksdk/5.92/api/"
)

func main() {
	vk := vksdk.Init("<TOKEN>")

	params := make(map[string]string)
	params["user_ids"] = "1"

	users, vkErr := vk.UsersGet(params)
	if vkErr.Code != 0 {
		log.Fatal(vkErr.Message)
	}

	for _, user := range users {
		fmt.Printf("Пользователя с id%d зовут %s %s\n", user.ID, user.FirstName, user.LastName)
	}
}
```

## Известные проблемы

- [VK API JSON Schema](https://github.com/VKCOM/vk-api-schema) кишит ошибками и не обновляется
- Документация VK имеет ошибки и не обновляется 
- На некоторые методы, API возвращает динамический JSON

### Костыли

[AccountGetInfo](https://vk.com/dev/account.getInfo) вместо поля `2fa_required`, вернет `TwoFactorRequired`

[Execute](https://vk.com/dev/execute) возвращает Response JSON в `[]byte`.

```go
var S struct {
	Text string `json:"text"`
}

rawResponse, vkErr := vk.Execute(`return {text: "hello"};`)
if vkErr.Code != 0 {
	log.Fatal(vkErr.Message)
}

err := json.Unmarshal(rawResponse, &S)
if err != nil {
	panic(err)
}

fmt.Print(S.Text)
```

[StorageGet](https://vk.com/dev/storage.get) если нет параметра `keys`, вернет массив из одного объекта.

Еще не реализованные: 

[AdsUpdateAds](https://vk.com/dev/ads.updateAds) поле `AdPlatform` вернет **строку**, поле `AdPlatformNoWall` вернет **число**.

[MessagesSend](https://vk.com/dev/messages.send) если нет параметра `user_ids`, вернет массив из одного объекта.

[LikesGetList](https://vk.com/dev/likes.getList) параметр `extended` всегда будет равен **1**

...

## TODO

- [ ] Все методы API
- [ ] Ограничитель на запросы
- [x] Callback
- [ ] LongPoll bot
- [ ] LongPoll user
- [ ] Streaming API
- [ ] Получение токена
- [ ] Создание develop ветки
- [ ] Поддержка go 1.11
- [ ] Тесты
- [ ] Англоязычный README
- [ ] Поддержка следующих версий API

### Список методов

- [x] Account 19/19
- [ ] Ads 0/0
- [ ] Apps 7/8
- [ ] AppWidgets 1/8
- [x] Auth 2/2
- [ ] Board 0/13
- [x] Database 11/12 (1 баговый метод)
- [ ] Docs 0/11
- [x] Execute
- [ ] Fave 0/12
- [ ] Friends 0/18
- [x] Gifts 1/1
- [ ] Groups 1/45
- [ ] leadForms 0/7
- [ ] Leads 0/0
- [ ] Likes 3/4
- [ ] Market 0/24
- [ ] Messages 35/39
- [ ] Newsfeed 0/16
- [ ] Notes 0/10
- [ ] Notifications 0/3
- [ ] Orders 0/0
- [ ] Pages 0/8
- [ ] Photos 0/46
- [ ] Places 0/0
- [ ] Polls 0/9
- [ ] Prettycards 0/6
- [ ] Search 0/1
- [ ] Secure 0/0
- [ ] Stats 0/3
- [x] Status 2/2
- [x] Storage 3/3
- [ ] Stories 0/13
- [x] Streaming 5/5
- [ ] Users 0/5
- [ ] Utils 1/7
- [ ] Video 0/24
- [ ] Podcasts 0/0
- [ ] Wall 0/23
- [ ] Widgets 0/2