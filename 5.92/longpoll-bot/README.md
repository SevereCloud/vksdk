# Bots Long Poll API

[Документация Bots Long Poll API](https://vk.com/dev/bots_longpoll)

## Подключение Bots Long Poll API

Чтобы использовать Bots Long Poll API, откройте раздел «Управление сообществом», на вкладке «Работа с API»→«Long Poll API» выберите «Включён».

### Версия API

Данная библиотека поддерживает версию API **5.92**.

### Инициализация

Модуль можно использовать с ключом доступа пользователя, полученным в Standalone-приложении через Implicit Flow(требуются права доступа: **groups**) или с ключом доступа сообщества(требуются права доступа: **manage**).

В начале необходимо инициализировать api:

```go
vk := api.Init("<TOKEN>")
```

А потом сам longpoll

```go
lp, err := longpoll.Init(vk api.VK, groupID int)
// По умолчанию Wait = 25
// lp.Wait = 90 
// lp.Ts = "123"
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

Для каждого события существует отдельный обработчик, который передает функции `object` и `groupID`.

Пример для события `message_new`

```go
lp.MessageNew(func(object object.MessageNewObject, groupID int) {
	...
})
```

Если вы хотите получать полный ответ от Long Poll(например для сохранения `ts` или специальной обработки `failed`), можно воспользоваться следующим обработчиком.

```go
lp.FullResponse(func(resp object.LongpollBotResponse) {
	...
})
```

Полный список событий Вы найдёте [в документации](https://vk.com/dev/groups_events)

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

	longpoll "github.com/SevereCloud/vksdk/5.92/longpoll-bot"
	"github.com/SevereCloud/vksdk/5.92/object"
)

func main() {
	vk := api.Init("<TOKEN>")
	lp, err := longpoll.Init(&vk, 12345678)
	if err != nil {
		panic(err)
	}

	lp.MessageNew(func(obj object.MessageNewObject, groupID int) {
		log.Print(obj.Text)
	})

	lp.Run()
}

```
