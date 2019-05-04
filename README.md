# VK SDK for Golang

VK SDK for Golang это реализация библиотеки для работы с VK API.

[![Go Report Card](https://goreportcard.com/badge/github.com/severecloud/vksdk)](https://goreportcard.com/report/github.com/severecloud/vksdk)

### Документация

- [Версия 5.92](https://github.com/SevereCloud/vksdk/tree/master/5.92)
- [API](https://github.com/SevereCloud/vksdk/tree/master/5.92/api)
- [Callback API](https://github.com/SevereCloud/vksdk/tree/master/5.92/callback)
- [Bots Long Poll API](https://github.com/SevereCloud/vksdk/tree/master/5.92/longpoll-bot)
- ~~Longpoll user~~
- ~~Streaming api~~

### Статус

Внимание - этот репозиторий в **очень ранней разработке**. Возможны серьезные изменения в коде.

### Установка

```shell
go get -u github.com/SevereCloud/vksdk
```

### Пример

```go
package main

import (
	"fmt"
	"log"

	vksdk "github.com/SevereCloud/vksdk/5.92/api"
)

func main() {
	vk := vksdk.Init("<TOKEN>")
	// vk.ProxyAddress = "127.0.0.1:9050"

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

### License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSevereCloud%2Fvksdk.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FSevereCloud%2Fvksdk?ref=badge_large)