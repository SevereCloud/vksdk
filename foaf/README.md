# FOAF

[![PkgGoDev](https://pkg.go.dev/badge/github.com/Derad6709/vksdk/v2/foaf)](https://pkg.go.dev/github.com/Derad6709/vksdk/v2/foaf)

FOAF (акроним от Friend of a Friend - "друг друга") является машиночитаемым
языком для описания людей, групп и отношений между ними.

VK имеет модифицированный FOAF, который был сделан по примеру яндекса + добавили
свои теги(некоторых тегов нет в документации). К сожалению, документация яндекса
потерялась во времени ([archive](https://web.archive.org/web/20140909053226/http://api.yandex.ru/blogs/doc/indexation/concepts/what-is-foaf.xml)).

## Примеры

Получение пользователя

```go
ctx := context.Background()

person, err := foaf.GetPerson(ctx, 1)
if err != nil {
  log.Fatal(err)
}

log.Println(person)
```

```go
ctx := context.Background()

person, err := foaf.GetGroup(ctx, 1)
if err != nil {
  log.Fatal(err)
}

log.Println(person)
```

Получение пользователя, используя кастомный HTTP-клиент

```go
ctx := context.Background()

// Use the custom HTTP client
httpClient := &http.Client{Timeout: 2 * time.Second}
ctx = context.WithValue(ctx, foaf.HTTPClient, httpClient)

person, err := foaf.GetPerson(ctx, 1)
if err != nil {
  log.Fatal(err)
}

log.Println(person)
```
