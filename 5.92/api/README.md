# API

[Документация VK API](https://vk.com/dev/first_guide)


### Версия API

Данная библиотека поддерживает версию API **5.92**.

### Инициализация

В начале необходимо инициализировать api с помощью [ключа доступа](https://vk.com/dev/access_token):

```go
vk := api.Init("<TOKEN>")
```

### Запросы к API

- `users.get` -> `vk.UsersGet(map[string]string{})`
- `groups.get` с extended=1 -> `vk.GroupsGetExtended(map[string]string{})`

Список всех методов можно найти на 
[данной странице](https://godoc.org/github.com/SevereCloud/vksdk/5.92/api#VK).

Пример запроса [`users.get`](https://vk.com/dev/users.get)

```go
users, vkErr := vk.UsersGet(map[string]string{
	"user_ids": "1"
})
if vkErr.Code != 0 {
	log.Fatal(vkErr.Message)
}
```

Описание ошибок:
- [документация VK](https://vk.com/dev/errors)
- [константы](https://godoc.org/github.com/SevereCloud/vksdk/5.92/object#pkg-constants)

Если `vkErr.Code` оказался отрицательным, значит ошибка на стороне клиента.

#### Запрос любого метода

Пример запроса [users.get](https://vk.com/dev/users.get)

```go
// Определяем структуру, которую вернет API
var response []object.UsersUser
var vkErr api.Error

params := map[string]string{
	"user_ids": "1"
}

// Делаем запрос
vk.RequestUnmarshal("users.get", params, &response, &vkErr)
if vkErr.Code != 0 {
	log.Fatal(vkErr.Message)
}
```

#### Execute

Универсальный метод, который позволяет запускать последовательность других 
методов, сохраняя и фильтруя промежуточные результаты.

[Документация VK](https://vk.com/dev/execute)

```go
var response struct {
	Text string `json:"text"`
}

vk.Execute(`return {text: "hello"};`, &response, &vkErr)
if vkErr.Code != 0 {
	log.Fatal(vkErr.Message)
}

log.Print(response.Text)
```

### Ограничитель запросов

К методам API ВКонтакте (за исключением методов из секций secure и ads) с 
ключом доступа пользователя или сервисным ключом доступа можно обращаться не 
чаще 3 раз в секунду. Для ключа доступа сообщества ограничение составляет 20 
запросов в секунду. Если логика Вашего приложения подразумевает вызов 
нескольких методов подряд, имеет смысл обратить внимание на метод execute. Он 
позволяет совершить до 25 обращений к разным методам в рамках одного запроса.

Для методов секции ads действуют собственные ограничения, ознакомиться с ними 
Вы можете на [этой странице](https://vk.com/dev/ads_limits).

Максимальное число обращений к методам секции secure зависит от числа 
пользователей, установивших приложение. Если приложение установило меньше 10 
000 человек, то можно совершать 5 запросов в секунду, до 100 000 — 8 запросов,
до 1 000 000 — 20 запросов, больше 1 млн. — 35 запросов в секунду.

Если Вы превысите частотное ограничение, сервер вернет ошибку с кодом 
**6: "Too many requests per second."**.

C помощью параметра `vk.Limit` можно установить ограничение на определенное 
количество запросов в секунду 

### Ошибка с Captcha

Если какое-либо действие (например, отправка сообщения) выполняется 
пользователем слишком часто, то запрос к API может возвращать ошибку 
"Captcha needed". При этом пользователю понадобится ввести код с изображения 
и отправить запрос повторно с передачей введенного кода Captcha в параметрах 
запроса.

**Код ошибки**: 14  
**Текст ошибки**: Captcha needed

Если возникает данная ошибка, то в сообщении об ошибке передаются также 
следующие параметры:
- `vkErr.CaptchaSID` - идентификатор captcha
- `vkErr.CaptchaImg` - ссылка на изображение, которое нужно показать 
  пользователю, чтобы он ввел текст с этого изображения.

В этом случае следует запросить пользователя ввести текст с изображения 
`vkErr.CaptchaImg` и повторить запрос, добавив в него параметры:
- `captcha_sid` - полученный идентификатор
- `captcha_key` - текст, который ввел пользователь

### HTTP client

В модуле реализована возможность изменять HTTP клиент с помощью параметра 
`vk.Client`

Пример прокси

```go
dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
httpTransport := &http.Transport{
	Dial:              dialer.Dial,
}
httpTransport.Dial = dialer.Dial
vk.Client.Transport = httpTransport
```

### Загрузка файлов

TODO: Загрузка файлов ещё не реализована
