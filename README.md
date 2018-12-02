# VK API for Golang

Внимание - этот репозиторий в **очень ранней разработке**.

Версия API 5.92

## Установка

```shell
go get -u github.com/severecloud/vkapi
```

## Пример

```go
package main

import (
	"fmt"
	"github.com/severecloud/vksdk"
)

func main() {
	vk := vksdk.Init("<TOKEN>")

	params := make(map[string]string)
	params["user_ids"] = "1"

	users, err := vk.UsersGet(params)
	if err != nil{
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("Пользователя с id%d зовут %s %s", user.ID, user.FirstName, user.LastName)
	}
}
```

## Известные проблемы

- [VK API JSON Schema](https://github.com/VKCOM/vk-api-schema) кишит ошибками и не обновляется
- На некоторые методы, API возвращает динамический json

### Костыли

[StorageGet](https://vk.com/dev/storage.get) даже если нет параметра `keys`, вернет массив из одного объекта.

Еще не реализованные: 

[AccountGetInfo](https://vk.com/dev/account.getInfo) вместо поля `2fa_required`, вернет `TwoFactorRequired`

[AdsUpdateAds](https://vk.com/dev/ads.updateAds) поле `AdPlatform` вернет **строку**, поле `AdPlatformNoWall` вернет **число**

[MessagesSend](https://vk.com/dev/messages.send) даже если нет параметра `user_ids`, вернет массив из одного объекта.

[LikesGetList](https://vk.com/dev/likes.getList) параметр `extended` всегда будет равен 1

...

## TODO

- [ ] Все методы API
- [ ] Callback
- [ ] LongPoll bot
- [ ] LongPoll user
- [ ] Streaming API
- [ ] Получение токена
- [ ] Тесты(чуть чуть)
- [ ] Поддержка go 1.11
- [ ] Англоязычный README
- [ ] Поддержка следующих версий API

### Список методов

- [ ] Account 0/19
- [ ] Ads 0/0
- [ ] Apps 0/8
- [ ] Board 0/13
- [ ] Database 0/10
- [ ] Docs 0/11
- [ ] Fave 0/12
- [ ] Friends 0/18
- [x] Gifts 1/1
- [ ] Groups 0/42
- [ ] leadForms 0/7
- [ ] Leads 0/0
- [ ] Likes 3/4
- [ ] Market 0/24
- [ ] Messages 0/29
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
- [ ] Streaming 0/5
- [ ] Users 0/5
- [ ] Utils 1/7
- [ ] Video 0/24
- [ ] Podcasts 0/0
- [ ] Wall 0/23
- [ ] Widgets 0/2