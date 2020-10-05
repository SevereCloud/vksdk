# Пакер

Пакер батчит все выполняемые VKSDK запросы в execute запросы.

## Пример

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/api/packer"
)

func main() {
	token := os.Getenv("TOKEN")
	vk := api.NewVK(token)
	packer.Default(vk, packer.Debug())

	resp, err := vk.UtilsResolveScreenName(
		params.NewUtilsResolveScreenNameBuilder().
			ScreenName("durov").Params,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("durov id:", resp.ObjectID)
}
```

## Параметры

Параметры передаются в виде аргументов в методы `packer.Default()` и `packer.New()`

- `packer.Debug()` включает вывод дебаг инфы
 (без этой опции пакер будет использовать токены применяющиеся в запросах)
- `packer.MaxPackedRequests(num)` устанавливает максимальное кол-во запросов в пачке
- `packer.Rules(mode, methods...)` устанавливает правила фильтрации методов\

Пример:

```go
// батчить только messages.send и messages.edit
packer.Default(vk, packer.Rules(packer.Allow, "messages.send", "messages.edit"))

// батчить все методы кроме groups.getMembers и board.getTopics
packer.Default(vk, packer.Rules(packer.Ignore, "groups.getMembers", "board.getTopics"))

// P.S. метод execute всегда выполняется отдельно
```
