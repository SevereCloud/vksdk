# User Long Poll API

[![Documentation](https://godoc.org/github.com/SevereCloud/vksdk/longpoll-user?status.svg)](https://pkg.go.dev/github.com/SevereCloud/vksdk/longpoll-user)
[![VK](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.com/dev/using_longpoll)

## Подключение User Long Poll API

Данный модуль по умолчанию поддерживает **3** версию

### Инициализация

Модуль можно использовать с ключом доступа пользователя, полученным в Standalone-приложении через Implicit Flow(требуются права доступа: **messages**) или с ключом доступа сообщества(требуются права доступа: **messages**).

В начале необходимо инициализировать api:

```go
vk := api.NewVK("<TOKEN>")
```

А потом сам longpoll

```go
mode := longpoll.ReceiveAttachments + longpoll.ExtendedEvents
lp, err := longpoll.NewLongpoll(vk, mode)
// По умолчанию Wait = 25
// lp.Wait = 90 
// lp.Ts = 123
```

### HTTP client

В модуле реализована возможность изменять HTTP клиент - `lp.Client`

Пример прокси

```go
dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
httpTransport := &http.Transport{
	Dial:              dialer.Dial,
	// DisableKeepAlives: true,
}

client := &http.Client{
	Transport: httpTransport,
}

lp.Client = client
```

### Обработчик событий

Обработчики, которые возвращают полноценные структуры:
- Для [3 версии](https://github.com/SevereCloud/vksdk/tree/master/longpoll-user/v3)

Для каждого события существует отдельный обработчик, который передает функции `[]interface{}`.

Пример для `4` события

```go
lp.EventNew(4, func(event []interface{}) {
	...
})
```

Если вы хотите получать полный ответ от Long Poll(например для сохранения `ts` или специальной обработки `failed`), можно воспользоваться следующим обработчиком.

```go
lp.FullResponse(func(resp object.LongpollResponse) {
	...
})
```

Полный список событий и их структуру Вы найдёте [в документации](https://vk.com/dev/using_longpoll?f=3.%2B%D0%A1%D1%82%D1%80%D1%83%D0%BA%D1%82%D1%83%D1%80%D0%B0%2B%D1%81%D0%BE%D0%B1%D1%8B%D1%82%D0%B8%D0%B9)

### Запуск и остановка

```go
// Запуск
if err := lp.Run(); err != nil {
	log.Fatal(err)
}

// Безопасное завершение
// Ждет пока соединение закроется и события обработаются
lp.Shutdown()

// Закрыть соединение
// Требует lp.Client.Transport = &http.Transport{DisableKeepAlives: true}
lp.Client.CloseIdleConnections()
```

## Пример

```go
package main

import (
	"log"

	"github.com/SevereCloud/vksdk/api"

	longpoll "github.com/SevereCloud/vksdk/longpoll-user"
)

func main() {
	vk := api.NewVK("<TOKEN>")
	lp, err := longpoll.NewLongpoll(vk, 2)
	if err != nil {
		log.Fatal(err)
	}

	lp.EventNew(4, func(event []interface{}) error {
		log.Print(event[5].(string))

		return nil
	})

	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}

```
