# API

[Документация VK API](https://vk.com/dev/first_guide)


### Версия API

Данная библиотека поддерживает версию API **5.92**.

### Инициализация

В начале необходимо инициализировать api:

```go
vk := api.Init("<TOKEN>")
```

### HTTP client

В модуле реализована возможность изменять HTTP клиент - `vk.Client`

Пример прокси

```go
dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
httpTransport := &http.Transport{
	Dial:              dialer.Dial,
}
httpTransport.Dial = dialer.Dial
vk.Client.Transport = httpTransport
```

### Тут должно быть продолжение

// TODO: дописать документацию