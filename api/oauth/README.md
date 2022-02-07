# Получение ключа доступа

[![PkgGoDev][doc-badge]][doc]
[![VK][dev-badge]](https://vk.com/dev/access_token)

Для работы со всеми методами API Вам необходимо передавать в запросе
**access_token** — специальный ключ доступа. Он представляет собой строку из
латинских букв и цифр и может соответствовать отдельному пользователю,
сообществу или самому Вашему приложению.

ВКонтакте поддерживает несколько способов получения ключа доступа по OAuth 2.0:

1. **Implicit flow** — требует встраивания браузера в ваше приложение.
Ключ возвращается на устройство пользователя, где был открыт диалог авторизации
(в виде дополнительного параметра URL). Такой ключ может быть использован
только для запросов непосредственно с устройства пользователя.

2. **Authorization code flow** — двухэтапный вариант с дополнительной
аутентификацией Вашего сервера. Ключ доступа возвращается непосредственно на
сервер и может быть использован, например, для автоматизированных запросов.

3. **Direct Authorization** — прямая авторизация используя логин и пароль.

## Права доступа приложения

[![VK][dev-badge]](https://vk.com/dev/permissions)

Права доступа определяют возможность использования токена для работы с тем или
иным разделом данных. Так, например, для отправки личного сообщения от имени
пользователя токен должен быть получен с правами `oauth.ScopeUserMessage`.

Обратите внимание, что некоторые права доступа пользователя из списка
(например, messages) могут быть запрошены только Standalone-приложением — что
означает необходимость использования Implicit Flow для запроса таких прав.

Если Вы хотите получить права на **Доступ к друзьям** и
**Доступ к статусам пользователя**, то Ваша битовая маска будет равна:

```go
scope := oauth.ScopeUserFriends + oauth.ScopeUserStatus // 1026
```

С помощью метода
[account.getAppPermissions](https://vk.com/dev/account.getAppPermissions),
можно получить битовую маску настроек текущего пользователя в данном приложении.

Если, имея битовую маску 1026, Вы хотите проверить, имеет ли она доступ к
друзьям — Вы можете сделать 1026 & 2.

```go
scope, err := vk.AccountGetAppPermissions(nil)
if err != nil {
    log.Fatal(err)
}

if oauth.CheckScope(scope, oauth.ScopeUserFriends, oauth.ScopeUserStatus) {
    log.Println("Имеется доступ к друзьям и статусу")
}
```

## Ключ доступа пользователя

### Authorization code flow

[![VK][dev-badge]](https://vk.com/dev/authcode_flow_user)

Используйте Authorization Code Flow для вызова методов API ВКонтакте с
серверной части Вашего приложения. Ключ доступа, полученный таким способом, не
привязан к IP-адресу, но набор прав, которые может получить приложение,
ограничен из соображений безопасности.

#### Сгенерируйте адрес

```go
acf := NewAuthCodeFlowUser(oauth.UserParams{
    ClientID:    123456,
    RedirectURI: "https://example.com/callback",
    Scope:       oauth.ScopeUserPhotos + oauth.ScopeUserDocs,
}, clientSecret)
u := acf.URL().String()
```

Необходимо перенаправить браузер пользователя по сгенерированному адресу.

Если пользователь не вошел на сайт, то в диалоговом окне ему будет предложено
ввести свой логин и пароль.

#### Разрешение прав доступа

После успешного входа на сайт пользователю будет предложено авторизовать
приложение, разрешив доступ к необходимым настройкам, запрошенным при помощи
параметра scope.

#### Получение access_token

После успешной авторизации приложения браузер пользователя будет перенаправлен
по адресу redirect_uri, указанному при открытии диалога авторизации. При этом
код для получения ключа доступа `code` будет передан как GET-параметр.

```go
func callback(w http.ResponseWriter, req *http.Request) {
    t, err := acf.Token(req.URL)
    if err != nil {
        fmt.Printf("%#v\n", err)
    }

    fmt.Printf(
        "Токен %s для id%d действует %d секунд",
        t.AccessToken,
        t.UserID,
        t.ExpiresIn,
    )
}
```

`acf.Token(req.URL)` выполнит запрос с вашего сервера, чтобы получить ключ
доступа.

### Implicit flow

[![VK][dev-badge]](https://vk.com/dev/implicit_flow_user)

Используйте Implicit Flow для вызова методов API ВКонтакте непосредственно с
устройства пользователя.

#### Сгенерируйте адрес

```go
u := oauth.ImplicitFlowUser(oauth.UserParams{
    ClientID: 123456,
    Scope:    oauth.ScopeUserPhotos + oauth.ScopeUserDocs,
})
```

RedirectURI по умолчанию `https://oauth.vk.com/blank.html`

#### Разрешение прав доступа

Необходимо перенаправить **встроенный браузер** по адресу сгенерированному на
предыдущем шаге.

Рекомендуется проверять, что страница вернула код 200. Если код отличается, а
`content-type` содержит JSON, необходимо распарсить этот JSON в ошибку:

```go
var errOAuth oauth.Error
err = json.Unmarshal(data, &errOAuth)
```

Если пользователь не авторизован ВКонтакте в используемом браузере, то в
диалоговом окне ему будет предложено ввести свой логин и пароль.

После успешного входа на сайт пользователю будет предложено авторизовать
приложение, разрешив доступ к необходимым настройкам, запрошенным при помощи
параметра scope.

#### Получение access_token

После успешной авторизации приложения браузер пользователя будет перенаправлен
по адресу redirect_uri, указанному при открытии диалога авторизации. При этом
ключ доступа к API access_token и другие параметры будут переданы в
URL-фрагменте ссылки.

```go
t, err := oauth.NewUserTokenFromURL(u)
if err != nil {
    fmt.Printf("%#v\n", err)
}

fmt.Printf(
    "Токен %s для id%d действует %d секунд",
    t.AccessToken,
    t.UserID,
    t.ExpiresIn,
)
```

`ExpiresIn` содержит 0, если токен бессрочный
(при использовании scope=offline).

### Direct Authorization

[![VK][dev-badge]](https://vk.com/dev/auth_direct)

> **Внимание! Доступ к этому типу авторизации может быть получен только после**
> **предварительного согласования с администрацией ВКонтакте.**
>
> Для подачи заявки на получение доступа Вам необходимо обратиться в службу
> поддержки, указав ID Вашего приложения.
>
> В настоящий момент эта возможность предоставляется только для платформ, не
> поддерживающих стандартную авторизацию. В заявке необходимо кратко описать
> функционал приложения.

Доверенные приложения могут получить неограниченный по времени access_token для
доступа к API, передав логин и пароль пользователя.

Обратите внимание, что приложение не должно хранить пароль пользователя.
Выдаваемый access_token не привязан к IP-адресу пользователя, поэтому его
достаточно для последующей работы с API без повторения процедуры авторизации.

```go
func example(params oauth.DirectAuthParams) (*oauth.UserToken, error) {
	t, err := oauth.DirectAuth(params)
	if err == nil {
		return t, nil
	}

	var e *oauth.Error
	if !errors.As(err, &e) {
		return nil, err
	}

	switch e.Type {
	case oauth.ErrNeedCaptcha:
		message := "Требуется ввести код с картинки:\n"
		message += e.CaptchaImg

		_, _ = fmt.Println(message)

		// ...

		params.CaptchaKey = ""
		params.CaptchaSID = e.CaptchaSID

		return example(params)
	case oauth.ErrNeedValidation:
		message := "Введите код"

		switch e.ValidationType {
		case oauth.ValidationSMS:
			message += ", который пришел в сообщении на номер " + e.PhoneMask
		case oauth.ValidationApp:
			message += " из приложения для генерации кодов"
		}

		fmt.Println(message)

		// ...

		params.Code = ""

		return example(params)
	}

	return nil, err
}


params := DirectAuthParams{
    ClientSecret: "secret",
    ClientID:     123456,
    Scope:        oauth.ScopeUserPhotos + oauth.ScopeUserDocs,
}

params.Username = "username"
params.Password = "password"

t, err := example(params)
if err != nil {
    fmt.Printf("%#v\n", err)
}


fmt.Printf(
    "Токен %s для id%d действует %d секунд",
    t.AccessToken,
    t.UserID,
    t.ExpiresIn,
)
```

## Ключ доступа сообщества

### Получение списка администрируемых сообществ

Получить ключ доступа сообщества через OAuth может только его администратор.
Чтобы получить ключи доступа сразу для всех или нескольких сообществ
пользователя, мы рекомендуем добавить этот дополнительный шаг в процесс
авторизации.

Получите ключ доступа пользователя с правами `scope=groups` и сделайте запрос к
методу [groups.get](https://vk.com/dev/groups.get) с параметром `filter=admin`,
чтобы получить список идентификаторов администрируемых сообществ.

Затем используйте все полученные значения или их часть в качестве параметра
`GroupIDs`.

### Authorization code flow

[![VK][dev-badge]](https://vk.com/dev/authcode_flow_group)

Используйте Authorization Code Flow для вызова методов API ВКонтакте с серверной
части Вашего приложения. Ключ доступа, полученный таким способом, не привязан к
IP-адресу.

#### Сгенерируйте адрес

```go
acf := NewAuthCodeFlowGroup(oauth.GroupParams{
    ClientID:    123456,
    GroupIDs:    []int{1234},
    RedirectURI: "https://example.com/callback",
    Scope:       oauth.ScopeGroupPhotos + oauth.ScopeGroupDocs,
}, clientSecret)
u := acf.URL().String()
```

Необходимо перенаправить браузер пользователя по сгенерированному адресу.

Если пользователь не вошел на сайт, то в диалоговом окне ему будет предложено
ввести свой логин и пароль.

#### Получение access_token

После успешной авторизации приложения браузер пользователя будет перенаправлен
по адресу redirect_uri, указанному при открытии диалога авторизации. При этом
код для получения ключа доступа `code` будет передан как GET-параметр.

```go
func callback(w http.ResponseWriter, req *http.Request) {
    t, err := acf.Token(req.URL)
    if err != nil {
        fmt.Printf("%#v\n", err)
    }

    for _, groupToken := range t.Groups {
        fmt.Printf(
            "Токен %s для club%d действует %d секунд",
            groupToken.AccessToken,
            groupToken.GroupID,
            t.ExpiresIn,
        )
    }
}
```

`acf.Token(req.URL)` выполнит запрос с вашего сервера, чтобы получить ключ
доступа.

### Implicit flow

[![VK][dev-badge]](https://vk.com/dev/implicit_flow_group)

Используйте Implicit Flow для вызова методов API ВКонтакте непосредственно с
устройства пользователя.

#### Сгенерируйте адрес

Перед открытием диалога авторизации рекомендуется убедиться, что пользователь
является администратором сообщества, для которого необходимо получить ключ
доступа

```go
u := oauth.ImplicitFlowGroup(oauth.GroupParams{
    ClientID: 123456,
    GroupIDs: []int{1234},
    Scope:    oauth.ScopeGroupPhotos + oauth.ScopeGroupDocs,
})
```

RedirectURI по умолчанию `https://oauth.vk.com/blank.html`

#### Разрешение прав доступа

Необходимо перенаправить **встроенный браузер** по адресу сгенерированному на
предыдущем шаге.

Рекомендуется проверять, что страница вернула код 200. Если код отличается, а
`content-type` содержит JSON, необходимо распарсить этот JSON в ошибку:

```go
var errOAuth oauth.Error
err = json.Unmarshal(data, &errOAuth)
```

После успешного входа на сайт пользователю будет предложено авторизовать
приложение, разрешив доступ к необходимым настройкам, запрошенным при помощи
параметра scope.

#### Получение access_token

После успешной авторизации приложения браузер пользователя будет перенаправлен
по адресу redirect_uri, указанному при открытии диалога авторизации. При этом
ключ доступа к API access_token и другие параметры будут переданы в
URL-фрагменте ссылки.

```go
t, err := oauth.NewGroupTokensFromURL(u)
if err != nil {
    fmt.Printf("%#v\n", err)
}

for _, groupToken := range t.Groups {
    fmt.Printf(
        "Токен %s для club%d действует %d секунд",
        groupToken.AccessToken,
        groupToken.GroupID,
        t.ExpiresIn,
    )
}
```

## Сервисный ключ доступа

Сервисный ключ нужен для запросов, которые не требуют авторизации пользователя
или сообщества. Это такие методы, как
[secure.sendNotification][secure.sendNotification] для отправки уведомлений от
приложения, или [secure.addAppEvent][secure.addAppEvent] для добавления
информации о достижениях, а также, начиная с апреля 2017 года, открытые методы,
например, [users.get][users.get].

Получить сервисный ключ доступа можно в [настройках][appsmanage] Вашего
приложения. Ключ не привязан к IP-адресу при использовании с открытыми
методами, срок его действия не ограничен. Если ключ был скомпрометирован, Вы
можете сгенерировать новый ключ, при этом старый будет аннулирован.

Сервисный ключ доступа идентифицирует Ваше приложение. Все запросы к API,
совершённые с использованием Вашего ключа доступа, будут считаться совершёнными
от имени Вашего приложения. Сервисный ключ доступа можно использовать только
для запросов с серверной стороны приложения, его нельзя передавать и хранить
на клиенте.

[doc-badge]: https://pkg.go.dev/badge/github.com/SevereCloud/vksdk/v2/api/oauth
[doc]: https://pkg.go.dev/github.com/SevereCloud/vksdk/v2/api/oauth
[dev-badge]: https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white
[users.get]: https://vk.com/dev/users.get
[secure.addAppEvent]: https://vk.com/dev/secure.addAppEvent
[secure.sendNotification]: https://vk.com/dev/secure.sendNotification
[appsmanage]: https://vk.com/apps?act=manage
