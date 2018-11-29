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

Еще не реализованные: 

[MessagesSend](https://vk.com/dev/messages.send) даже если нет параметра `user_ids`, вернет массив объектов, каждый из которых содержит поля: `PeerId` `NessageId` `Error`

[AccountGetInfo](https://vk.com/dev/account.getInfo) вместо поля `2fa_required`, вернет `TwoFactorRequired`

[AdsUpdateAds](https://vk.com/dev/ads.updateAds) поле `AdPlatform` вернет **строку**, поле `AdPlatformNoWall` вернет **число*

...

## TODO

- [ ] 407 методов
- [ ] Callback
- [ ] LongPoll bot
- [ ] LongPoll user
- [ ] Streaming API
- [ ] Получение токена
- [ ] Тесты
- [ ] Поддержка go 1.11
- [ ] Англоязычный README
- [ ] Поддержка следующих версий API
