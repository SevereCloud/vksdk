# API

[![документация](https://godoc.org/github.com/SevereCloud/vksdk/errors?status.svg)](http://godoc.org/github.com/SevereCloud/vksdk/api)
[![VK](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.com/dev/first_guide)


### Версия API

Данная библиотека поддерживает версию API **5.101**.

### Инициализация

В начале необходимо инициализировать api с помощью [ключа доступа](https://vk.com/dev/access_token):

```go
vk := api.Init("<TOKEN>")
```

### Запросы к API

- `users.get` -> `vk.UsersGet(map[string]string{})`
- `groups.get` с extended=1 -> `vk.GroupsGetExtended(map[string]string{})`

Список всех методов можно найти на 
[данной странице](https://godoc.org/github.com/SevereCloud/vksdk/api#VK).

Пример запроса [`users.get`](https://vk.com/dev/users.get)

```go
users, err := vk.UsersGet(map[string]string{
	"user_ids": "1"
})
if err != nil {
	log.Fatal(err)
}
```

#### Обработка ошибок

[![документация](https://godoc.org/github.com/SevereCloud/vksdk/errors?status.svg)](http://godoc.org/github.com/SevereCloud/vksdk/errors)
[![VK](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.com/dev/errors)

Пример обработки ошибки

```go
// import "github.com/SevereCloud/vksdk/errors"

switch errors.GetType(err) {
case errors.NoType:
	log.Print("Ошибка не связанная с работай самого апи")
case Captcha:
	log.Print("Требуется ввод кода с картинки (Captcha)")
case 1:
	log.Print("Код ошибки 1")
default:
	log.Print("Другая ошибка")
}
```

Получение ошибки, [отправленной ВК](http://godoc.org/github.com/SevereCloud/vksdk/object#Error)

```go
// import "github.com/SevereCloud/vksdk/errors"

vkErr := errors.GetErrorContext(err)
```

#### Запрос любого метода

Пример запроса [users.get](https://vk.com/dev/users.get)

```go
// Определяем структуру, которую вернет API
var response []object.UsersUser
var err api.Error

params := map[string]string{
	"user_ids": "1"
}

// Делаем запрос
err = vk.RequestUnmarshal("users.get", params, &response)
if err != nil {
	log.Fatal(err)
}
```

#### Execute

[![документация](https://godoc.org/github.com/SevereCloud/vksdk/errors?status.svg)](http://godoc.org/github.com/SevereCloud/vksdk/api#VK.Execute)
[![VK](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.com/dev/execute)

Универсальный метод, который позволяет запускать последовательность других 
методов, сохраняя и фильтруя промежуточные результаты.

```go
var response struct {
	Text string `json:"text"`
}

err = vk.Execute(`return {text: "hello"};`, &response)
if err != nil {
	log.Fatal(err)
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

[![VK](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.com/dev/captcha_error)

Если какое-либо действие (например, отправка сообщения) выполняется 
пользователем слишком часто, то запрос к API может возвращать ошибку 
"Captcha needed". При этом пользователю понадобится ввести код с изображения 
и отправить запрос повторно с передачей введенного кода Captcha в параметрах 
запроса.

**Код ошибки**: 14  
**Текст ошибки**: Captcha needed

Если возникает данная ошибка, то в сообщении об ошибке передаются также 
следующие параметры:
- `err.CaptchaSID` - идентификатор captcha
- `err.CaptchaImg` - ссылка на изображение, которое нужно показать 
  пользователю, чтобы он ввел текст с этого изображения.

В этом случае следует запросить пользователя ввести текст с изображения 
`err.CaptchaImg` и повторить запрос, добавив в него параметры:
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

[![VK](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.com/dev/upload_files)

#### 1. Загрузка фотографий в альбом

Допустимые форматы: JPG, PNG, GIF. 
Файл объемом не более 50 МБ, соотношение сторон не менее 1:20

Загрузка фотографий в альбом для текущего пользователя:

```go
photosPhoto, err = vk.UploadPhoto(albumID, response.Body)
```

Загрузка фотографий в альбом для группы:

```go
photosPhoto, err = vk.UploadPhotoGroup(groupID, albumID, response.Body)
```

#### 2. Загрузка фотографий на стену

Допустимые форматы: JPG, PNG, GIF. 
Файл объемом не более 50 МБ, соотношение сторон не менее 1:20

```go
photosPhoto, err = vk.UploadWallPhoto(response.Body)
```

Загрузка фотографий в альбом для группы:

```go
photosPhoto, err = vk.UploadWallPhotoGroup(groupID, response.Body)
```

#### 3. Загрузка главной фотографии пользователя или сообщества

Допустимые форматы: JPG, PNG, GIF.
Ограничения: размер не менее 200x200px, соотношение сторон от 0.25 до 3, сумма высоты и ширины не более 14000px, файл объемом не более 50 МБ, соотношение сторон не менее 1:20.

Загрузка главной фотографии пользователя

```go
photosPhoto, err = vk.UploadUserPhoto(file)
```

Загрузка фотографии пользователя или сообщества с миниатюрой

```go
photosPhoto, err = vk.UploadOwnerPhoto(ownerID, squareСrop,file)
```

Для загрузки главной фотографии сообщества необходимо передать его идентификатор со знаком «минус» в параметре `ownerID`.

Дополнительно Вы можете передать параметр `squareСrop` в формате "x,y,w" (без кавычек), где x и y — координаты верхнего правого угла миниатюры, а w — сторона квадрата. Тогда для фотографии также будет подготовлена квадратная миниатюра.

Загрузка фотографии пользователя или сообщества без миниатюры:

```go
photosPhoto, err = vk.UploadOwnerPhoto(ownerID, "", file)
```

#### 4. Загрузка фотографии в личное сообщение

Допустимые форматы: JPG, PNG, GIF.
Ограничения: сумма высоты и ширины не более 14000px, файл объемом не более 50 МБ, соотношение сторон не менее 1:20.

```go
photosPhoto, err = vk.UploadMessagesPhoto(peerID, file)
```

#### 5. Загрузка главной фотографии для чата

Допустимые форматы: JPG, PNG, GIF.
Ограничения: размер не менее 200x200px, соотношение сторон от 0.25 до 3, сумма высоты и ширины не более 14000px, файл объемом не более 50 МБ, соотношение сторон не менее 1:20.

Без обрезки:

```go
messageInfo, err = vk.UploadChatPhoto(peerID, file)
```

С обрезкой:

```go
messageInfo, err = vk.UploadChatPhotoCrop(peerID, cropX, cropY, cropWidth, file)
```

#### 6. Загрузка фотографии для товара

Допустимые форматы: JPG, PNG, GIF.
Ограничения: минимальный размер фото — 400x400px, сумма высоты и ширины не более 14000px, файл объемом не более 50 МБ, соотношение сторон не менее 1:20.

Если Вы хотите загрузить основную фотографию товара, необходимо передать параметр `mainPhoto = true`.  Если фотография не основная, она не будет обрезаться.

Без обрезки:

```go
photosPhoto, err = vk.UploadMarketPhoto(groupID, mainPhoto, file)
```

Основную фотографию c обрезкой:

```go
photosPhoto, err = vk.UploadMarketPhotoCrop(groupID, cropX, cropY, cropWidth, file)
```

#### 7. Загрузка фотографии для подборки товаров

Допустимые форматы: JPG, PNG, GIF.
Ограничения: минимальный размер фото — 1280x720px, сумма высоты и ширины не более 14000px, файл объемом не более 50 МБ, соотношение сторон не менее 1:20.

```go
photosPhoto, err = vk.UploadMarketAlbumPhoto(groupID, file)
```

#### 9. Загрузка видеозаписей

Допустимые форматы: AVI, MP4, 3GP, MPEG, MOV, MP3, FLV, WMV.

[Параметры](https://vk.com/dev/video.save)

```go
videoUploadResponse, err = vk.UploadVideo(params, file)
```

После загрузки видеозапись проходит обработку и в списке видеозаписей может появиться спустя некоторое время.

#### 10. Загрузка документов

Допустимые форматы: любые форматы за исключением mp3 и исполняемых файлов.
Ограничения: файл объемом не более 200 МБ.

`title` - название файла с расширением

`tags` - метки для поиска

`typeDoc` - тип документа.

- doc - обычный документ;
- audio_message - голосовое сообщение
- graffiti - граффити


Загрузить документ:

```go
docsDoc, err = vk.UploadDoc(title, tags, file)
```

Загрузить документ в группу:

```go
docsDoc, err = vk.UploadGroupDoc(groupID, title, tags, file)
```

Загрузить документ, для последующей отправки документа на стену:

```go
docsDoc, err = vk.UploadWallDoc(title, tags, file)
```

Загрузить документ в группу, для последующей отправки документа на стену:

```go
docsDoc, err = vk.UploadGroupWallDoc(groupID, title, tags, file)
```

Загрузить документ в личное сообщение:

```go
docsDoc, err = vk.UploadMessagesDoc(peerID, typeDoc, title, tags, file)
```

#### 11. Загрузка обложки сообщества

Допустимые форматы: JPG, PNG, GIF.
Ограничения: минимальный размер фото — 795x200px, сумма высоты и ширины не более 14000px, файл объемом не более 50 МБ. Рекомендуемый размер: 1590x400px. В сутки можно загрузить не более 1500 обложек.

Необходимо указать координаты обрезки фотографии в параметрах `cropX`, `cropY`, `cropX2`, `cropY2`.

```go
photosPhoto, err = vk.UploadOwnerCoverPhoto(groupID, cropX, cropY, cropX2, cropY2, file)
```

#### 12. Загрузка аудиосообщения

Допустимые форматы: Ogg Opus.
Ограничения: sample rate 16kHz, variable bitrate 16 kbit/s, длительность не более 5 минут.

```go
docsDoc, err = vk.UploadMessagesDoc(peerID, "audio_message", title, tags, file)
```

#### 13. Загрузка истории

Допустимые форматы:​ JPG, PNG, GIF.
Ограничения:​ сумма высоты и ширины не более 14000px, файл объемом не более 10МБ. Формат видео: h264 video, aac audio, максимальное разрешение 720х1280, 30fps.

Загрузить истроию с фотографией. [Параметры](https://vk.com/dev/stories.getPhotoUploadServer)

```go
uploadInfo, err = vk.UploadStoriesPhoto(params, file)
```

Загрузить истроию с видео. [Параметры](https://vk.com/dev/stories.getVideoUploadServer)

```go
uploadInfo, err = vk.UploadStoriesVideo(params, file)
```

#### Загрузка фоновой фотографии в опрос

Допустимые форматы: JPG, PNG, GIF. 
Ограничения: сумма высоты и ширины не более 14000px, файл объемом не более 50 МБ, соотношение сторон не менее 1:20.

```go
photosPhoto, err = vk.UploadPollsPhoto(file)
```

```go
photosPhoto, err = vk.UploadOwnerPollsPhoto(ownerID, file)
```

Для загрузки фотографии сообщества необходимо передать его идентификатор со знаком «минус» в параметре `ownerID`.

#### Загрузка фотографии для карточки

Для карточек используются квадратные изображения минимальным размером 400х400. В случае загрузки неквадратного изображения, оно будет обрезано до квадратного.
Допустимые форматы: JPG, PNG, BMP, TIFF или GIF. Ограничения: файл объемом не более 5 МБ.

```go
photo, err = vk.UploadPrettyCardsPhoto(file)
```

#### Загрузка обложки для формы

Для форм сбора заявок используются прямоугольные изображения размером 1200х300. В случае загрузки изображения другого размера, оно будет автоматически обрезано до требуемого. Допустимые форматы: JPG, PNG, BMP, TIFF или GIF. Ограничения: файл объемом не более 5 МБ.

```go
photo, err = vk.UploadLeadFormsPhoto(file)
```

Полученные данные можно использовать в методах [leadForms.create](https://vk.com/dev/leadForms.create) и [leadForms.edit](https://vk.com/dev/leadForms.edit).

Полученные данные можно использовать в методах [prettyCards.create](https://vk.com/dev/prettyCards.create) и [prettyCards.edit](https://vk.com/dev/prettyCards.edit).

#### Загрузки фотографии в коллекцию приложения для виджетов приложений сообществ

`imageType` (string) - тип изображения. 

Возможные значения:
- 24x24
- 50x50
- 160x160
- 160x240
- 510x128

```go
photosPhoto, vkErr = vk.UploadAppImage(imageType, file)
```

#### Примеры

Загрузка фотографии в альбом:

```go
response, err := os.Open("photo.jpeg")
if err != nil {
	log.Fatal(err)
}
defer response.Body.Close()

photo, err = vk.UploadPhoto(albumID, response.Body)
if err != nil {
	log.Fatal(err)
}
```

Загрузка фотографии в альбом из интернета:

```go
response, err := http.Get("https://sun9-45.userapi.com/c638629/v638629852/2afba/o-dvykjSIB4.jpg")
if err != nil {
	log.Fatal(err)
}
defer response.Body.Close()

photo, err = vk.UploadPhoto(albumID, response.Body)
if err != nil {
	log.Fatal(err)
}
```
