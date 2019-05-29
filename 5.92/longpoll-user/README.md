# User Long Poll API

[Документация User Long Poll API](https://vk.com/dev/using_longpoll)

## Подключение User Long Poll API

Данная библиотека поддерживает версию **3**.

### Инициализация

Модуль можно использовать с ключом доступа пользователя, полученным в Standalone-приложении через Implicit Flow(требуются права доступа: **messages**) или с ключом доступа сообщества(требуются права доступа: **messages**).

В начале необходимо инициализировать api:

```go
vk := api.Init("<TOKEN>")
```

А потом сам longpoll

```go
lp, err := longpoll.Init(vk api.VK, groupID int)
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
httpTransport.Dial = dialer.Dial
lp.Client.Transport = httpTransport
```

### Обработчик событий

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

	"github.com/SevereCloud/vksdk/5.92/api"

	longpoll "github.com/SevereCloud/vksdk/5.92/longpoll-user"
)

func main() {
	vk := api.Init("<TOKEN>")
	lp, err := longpoll.Init(&vk, 2)
	if err != nil {
		log.Fatal(err)
	}

	lp.EventNew(4, func(event []interface{}) {
		log.Print(event[5].(string))
	})

	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}

```
