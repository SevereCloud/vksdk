# Callback API

[![PkgGoDev](https://pkg.go.dev/badge/github.com/SevereCloud/vksdk/v2/callback)](https://pkg.go.dev/github.com/SevereCloud/vksdk/v2/callback)
[![VK](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://dev.vk.com/ru/api/callback/getting-started)

## Подключение Callback API

Для подключения Callback API нужно открыть раздел «Управление сообществом»
(«Управление страницей», если у Вас публичная страница), перейти во вкладку
«Работа с API».
Далее необходимо указать и подтвердить конечный адрес сервера, куда будут
направлены все запросы. Вы можете подключить до 10 серверов для Callback API,
задать каждому из них отдельный набор событий и версию API.

### Версия API

Данная библиотека поддерживает версию API **5.131**.

### Подтверждение сервера

Для подтверждения сервера задайте строку, которую должен вернуть сервер,
переменной `ConfirmationKeys[groupID]` для отдельной группы или `ConfirmationKey`

### Секретный ключ

Для проверки того, что запросы приходят от VK, укажите секретный ключ. Секретный
ключ задайте переменной `SecretKeys[groupID]` для отдельной группы или `SecretKey`

Если ключи не совпадают, Ваш сервер вернет `Bad Secret`

### Обработчик событий

Для каждого события существует отдельный обработчик, который передает функции
`ctx` и `object`.

Пример для события `message_new`

```go
cb.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
	...
})
```

Полный список событий Вы найдёте [в документации](https://dev.vk.com/ru/api/community-events/json-schema)

### Контекст

Поля `groupID` и `eventID` передаются в `ctx`. Чтобы получить их, можно
воспользоваться следующими функциями:

```go
groupID := events.GroupIDFromContext(ctx)
eventID := events.EventIDFromContext(ctx)
```

### Веб-сервер

Для модуля **net/http** воспользуйтесь функцией `HandleFunc`

```go
http.HandleFunc("/callback", cb.HandleFunc)
http.ListenAndServe(":8080", nil)
```

## Пример

```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/SevereCloud/vksdk/v2/callback"
	"github.com/SevereCloud/vksdk/v2/events"
)

func main() {
	cb := callback.NewCallback()

	cb.ConfirmationKey = "693d0ba9"
	// cb.ConfirmationKeys[170561776] = "693d0ba9"

	cb.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		log.Print(obj.Message.Text)
	})

	http.HandleFunc("/callback", cb.HandleFunc)

	http.ListenAndServe(":8080", nil)
}
```

## Автоматическая настройка

Для автоматической настройки callback сервера, существует метод `AutoSetting`.
Данный метод:

- проверяет существующие настройки callback
- удаляет сервер, если он сломан
- создает новый callback, если его нет
- генерирует секрет
- настраивает callback сервер с событиями, которые были прописаны в коде

`AutoSetting` требуется запускать вместе с веб-сервером.

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/callback"
	"github.com/SevereCloud/vksdk/v2/events"
)

func main() {
	groupToken := "<TOKEN>" // рекомендуется использовать os.Getenv("TOKEN")
	vk := api.NewVK(groupToken)

	cb := callback.NewCallback()
	cb.Title = "example-bot"

	cb.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		log.Print(obj.Message.Text)
	})

	http.HandleFunc("/callback", cb.HandleFunc)

	go func() {
		err := cb.AutoSetting(vk, "https://example.com/callback")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}()

	http.ListenAndServe(":8080", nil)
}
```
