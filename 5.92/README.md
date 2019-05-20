# VK API for Golang 5.92

Внимание - этот репозиторий в **очень ранней разработке**.

## Установка

```shell
go get -u github.com/SevereCloud/vksdk
```

## Модули

- [API](https://github.com/SevereCloud/vksdk/tree/master/5.92/api)
- [Callback](https://github.com/SevereCloud/vksdk/tree/master/5.92/callback)
- [Bots Long Poll API](https://github.com/SevereCloud/vksdk/tree/master/5.92/longpoll-bot)

## Пример

```go
package main

import (
	"fmt"
	"log"

	vksdk "github.com/SevereCloud/vksdk/5.92/api"
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

[StorageGet](https://vk.com/dev/storage.get) если нет параметра `keys`, вернет массив из одного объекта.
...

## TODO

- [ ] Большенство методов API
- [x] Ограничитель на запросы
- [x] Callback
- [x] LongPoll bots
- [ ] LongPoll users
- [ ] Streaming API
- [x] HTTP client
- [ ] Получение токена
- [ ] Загрузка файлов
- [ ] Тесты
- [ ] Англоязычный README
- [ ] Версионирование
- [ ] Поддержка следующих версий API

### Список методов

- [ ] Ads 0/44
- [ ] Leads 0/6
- [ ] Newsfeed 0/18
- [ ] Orders 0/8
- [ ] Photos 0/51
- [ ] Prettycards 0/6
- [ ] Secure 0/10
- [x] Account 19/19
- [x] Apps 9/9
- [x] AppWidgets 8/8
- [x] Auth 2/2
- [x] Board 15/15
- [x] Database 12/12
- [x] Docs 11/11
- [x] Execute
- [x] Fave 15/15
- [x] Friends 19/19
- [x] Gifts 1/1
- [x] Groups 50/50
- [x] LeadForms 7/7
- [x] Likes 5/5
- [x] Market 0/27
- [x] Messages 43/43
- [x] Notes 10/10
- [x] Notifications 3/3
- [x] Pages 8/8
- [x] Polls 10/10
- [x] Search 1/1
- [x] Stats 3/3
- [x] Status 2/2
- [x] Storage 3/3
- [x] Stories 18/18
- [x] Streaming 5/5
- [x] Users 7/7
- [x] Utils 8/8
- [x] Video 0/29
- [x] Wall 28/28
- [x] Widgets 2/2