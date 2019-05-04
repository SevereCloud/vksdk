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
```

### Обработчик событий

Для каждого события существует отдельный обработчик, который передает функции `object` и `groupID`.

Пример для события `message_new`

```go
lp.MessageNew(func(object object.MessageNewObject, groupID int) {
	...
})
```

Полный список событий Вы найдёте [в документации](https://vk.com/dev/groups_events)

### Запуск и остановка

```go
// Запуск 
lp.Run()
// Завершение
lp.Shutdown()
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
	lp, err := longpoll.Init(vk, 12345678)
	if err != nil {
		panic(err)
	}

	lp.MessageNew(func(obj object.MessageNewObject, groupID int) {
		log.Print(obj.Text)
	})

	lp.Run()
}

```
