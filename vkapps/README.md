# VK Mini Apps

[![PkgGoDev](https://pkg.go.dev/badge/github.com/SevereCloud/vksdk/v2/vkapps)](https://pkg.go.dev/github.com/SevereCloud/vksdk/v2/vkapps)
[![VK](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.com/dev/vk_apps_docs)

[VK Mini Apps](https://vk.com/dev/vk_apps_docs) — это платформа встраиваемых
кроссплатформенных приложений ВКонтакте. В данном модуле собраны инструменты для
Backend части приложения.

## Проверка параметров запуска

При запуске сервиса на указанный в управлении приложением URL передаются
дополнительные параметры, содержащие в себе данные о пользователе и об источнике
запуска.

Проверка подписи в ссылке:

```go
link := "https://example.com/?vk_user_id=494075&vk_app_id=6736218&vk_is_app_user=1&vk_are_notifications_enabled=1&vk_language=ru&vk_access_token_settings=&vk_platform=android&sign=htQFduJpLxz7ribXRZpDFUH-XEUhC9rBPTJkjUFEkRA"

v, _ := vkapps.ParamsVerify(link, "wvl68m4dR1UpLrVRli")

fmt.Println(v)

// Output:
// true
```

### VerifyMiddleware(next http.Handler) http.Handler

`VerifyMiddleware` это промежуточный http обработчик, который проверяет подпись.
Если подпись верна - вызывает следующий обработчик, иначе возвращает 403 ошибку.

```go
pv := vkapps.NewParamsVerification(clientSecret)

// Перед следующие обработчиками будет проверяться подпись
http.HandleFunc("/api/user/", pv.VerifyMiddleware(UserHandler))
http.HandleFunc("/api/user/details", pv.VerifyMiddleware(UserDetailsHandler))

```

Пример использования в [mux](https://github.com/gorilla/mux#middleware)

```go
pv := vkapps.NewParamsVerification(clientSecret)

r := mux.NewRouter()

// Перед следующие обработчиками будет проверяться подпись
s := r.PathPrefix("/api").Subrouter()
s.HandleFunc("/user/", UserHandler) // /api/user/
s.HandleFunc("/user/details", UserDetailsHandler) // /api/user/details
s.Use(pv.VerifyMiddleware)

// Проверки подписи не будет
r.HandleFunc("/", PublicHandler)
```
