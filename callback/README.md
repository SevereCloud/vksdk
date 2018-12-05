# Callback API

[Документация Callback API](https://vk.com/dev/callback_api)

## Подключение Callback API

Для подключения Callback API нужно открыть раздел «Управление сообществом» («Управление страницей», если у Вас публичная страница), перейти во вкладку «Работа с API». 
Далее необходимо указать и подтвердить конечный адрес сервера, куда будут направлены все запросы. Вы можете подключить до 10 серверов для Callback API, задать каждому из них отдельный набор событий и версию API.

### Версия API

Данная библиотека поддерживает версию API **5.92**.

### Подтверждение сервера

Для подтверждения сервера задайте строку, которую должен вернуть сервер, переменной `ConfirmationKeys[groupID]` для отдельной группы или `ConfirmationKey`

### Секретный ключ

Для проверки того, что запросы приходят от VK, укажите секретный ключ. Cекретный ключ задайте переменной `SecretKeys[groupID]` для отдельной группы или `SecretKey`

Если ключи не совпадают, Ваш сервер вернет `bad secret`

### Обработчик событий

Для каждого события существует отдельный обработчик, который передает функции `object` и `groupID`.

Пример для события `message_new`

```go
cb.MessageNew(func(object object.MessageNewObject, groupID int) {
	...
})
```

Полный список событий Вы найдёте [в документации](https://vk.com/dev/groups_events)

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
	"log"
	"net/http"

	"github.com/severecloud/vksdk/callback"
	"github.com/severecloud/vksdk/object"
)

func main() {
	var cb callback.Callback
	
	cb.ConfirmationKey = "693d0ba9"
	// cb.ConfirmationKeys[170561776] = "693d0ba9"

	cb.MessageNew(func(obj object.MessageNewObject, groupID int) {
		log.Print(obj.Text)
	})

	http.HandleFunc("/callback", cb.HandleFunc)

	http.ListenAndServe(":8080", nil)
}
```
