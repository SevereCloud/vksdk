# VK SDK for Golang

[![Build Status](https://travis-ci.com/SevereCloud/vksdk.svg?branch=master)](https://travis-ci.com/SevereCloud/vksdk)
[![GolangCI](https://golangci.com/badges/github.com/SevereCloud/vksdk.svg)](https://golangci.com/r/github.com/SevereCloud/vksdk)
[![Documentation](https://godoc.org/github.com/SevereCloud/vksdk?status.svg)](http://godoc.org/github.com/SevereCloud/vksdk)
[![codecov](https://codecov.io/gh/SevereCloud/vksdk/branch/master/graph/badge.svg)](https://codecov.io/gh/SevereCloud/vksdk)
[![VK](https://img.shields.io/badge/chat-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.me/join/AJQ1d6Or8Q00Y_CSOESfbqGt)
[![license](https://img.shields.io/github/license/SevereCloud/vksdk.svg?maxAge=2592000)](https://github.com/SevereCloud/vksdk/blob/master/LICENSE)

**VK SDK for Golang** готовая реализация основных функций VK API для языка Go.

### Возможности

- [API](https://github.com/SevereCloud/vksdk/tree/master/5.92/api#api)
  - Возвращает готовые структуры
  - Реализовано 400+ методов
  - Возможность изменять HTTP клиент
  - Ограничитель запросов
  - Загрузка файлов
- [Callback API](https://github.com/SevereCloud/vksdk/tree/master/5.92/callback#callback-api)
  - Поддерживает все события
  - Возвращает готовые структуры
- [Bots Long Poll API](https://github.com/SevereCloud/vksdk/tree/master/5.92/longpoll-bot#bots-long-poll-api)
  - Поддерживает все события
  - Возвращает готовые структуры
  - Возможность изменять HTTP клиент
- [User Long Poll API](https://github.com/SevereCloud/vksdk/tree/master/5.92/longpoll-user#user-long-poll-api)
  - ~~Возвращает готовые структуры~~ [#44](https://github.com/SevereCloud/vksdk/issues/44)
  - Возможность изменять HTTP клиент
- ~~Streaming api~~ [#3](https://github.com/SevereCloud/vksdk/issues/3)

### Статус

Внимание - этот репозиторий в **очень ранней разработке**. Возможны серьезные изменения в коде - cмотри [#40](https://github.com/SevereCloud/vksdk/issues/40)

### Установка

```shell
go get -u github.com/SevereCloud/vksdk
```

### Пример

```go
package main

import (
    "log"

    vkapi "github.com/SevereCloud/vksdk/5.92/api"
)

func main() {
    vk := vkapi.Init("<TOKEN>") // рекомендуется использовать os.Getenv("TOKEN")
    
    params := map[string]string{
        "user_ids": "1"
    }

    users, vkErr := vk.UsersGet(params)
    if vkErr.Code != 0 {
        log.Fatal(vkErr.Message)
    }

    for _, user := range users {
        log.Printf("Пользователя с id%d зовут %s %s\n", user.ID, user.FirstName, user.LastName)
    }
}
```

### Лицензия

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSevereCloud%2Fvksdk.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FSevereCloud%2Fvksdk?ref=badge_large)