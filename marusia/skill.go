/*
Package marusia для создания скилла Маруси.

Для голосового помощника Маруси теперь можно создавать скиллы, которые
пополнят её базу навыков. Пользователям будет удобнее общаться с Марусей,
а разработчики и владельцы бизнеса смогут сделать голосовой интерфейс для
своих продуктов.

Документация: https://vk.com/dev/marusia_skill_docs

Регистрация приложения ВКонтакте

Скилл можно создать в разделе для разработчиков:
https://vk.com/editapp?act=create.
В день можно создать не более 3-х скиллов, и не более 10-и за 5 дней.

Для этого:

1. Выберите «Скилл Маруси» в типах приложения;

2. Добавьте название, которое будет совпадать с командой для активации скилла;

3. Введите в поле Webhook URL адрес сервера, по которому будет размещен навык,
например https://example.com/test-webhook;

4. Подтвердите действие.

Вы попадёте в интерфейс администрирования Вашего скилла.

Обратите внимание: имя является первой фразой-триггером для вызова скилла.

Фразы

Фразы для вызова скилла должны быть специфичны и уникальны, чтобы мы могли
использовать их для внешних скиллов. Например, фразу «Расскажи анекдот»
добавить не сможем, т.к. она уже используется во внутренних скиллах Маруси.
А вот фразу «Давай сделаем кодревью» — пока можно использовать для внешнего
скилла.

Длина фразы активации не может превышать 64 символа.

Фраза вызова скилла строится по следующей схеме:

	Слово «Маруся» + любая из дефолтных фраз вызова скилла + фраза активации.

Дефолтные фразы вызова скилла:

	хочу скилл/навык

	запусти скилл/навык

	включи скилл/навык

	открой скилл/навык

Примеры:
«Маруся, запусти навык шутка дня», «Маруся, включи скилл шутка дня».

Схема взаимодействия со скиллом

1. Пользователь произносит фразу для вызова скилла.

2. Получив и распознав выражение, сервер Маруси отправляет POST-запрос на
Webhook URL, который вы указали в настройках.

3. Обработчик скилла должен ответить на полученный от сервера Маруси запрос.
Таймаут ожидания ответа — 5 секунд, после чего сервер Маруси завершит сессию.

При создании скилла по умолчанию доступ к нему имеют только его
администраторы. Для того, чтобы проверить скилл, авторизуйтесь в приложении
Маруси для Android или iOS через VK Connect, используя номер, привязанный к
учётной записи ВКонтакте.

Запустить скилл можно как голосовой командой, так и при помощи чата в
приложении.

Также протестировать и отладить скилл можно в отладчике скиллов
https://skill-tester.marusia.mail.ru/

Пример

Пример скилла.


	type myPayload struct {
		Text string
		marusia.DefaultPayload
	}

	wh := marusia.NewWebhook()

	wh.OnEvent(func(r marusia.Request) (resp marusia.Response) {
		switch r.Request.Type {
		case marusia.SimpleUtterance:
			switch r.Request.Command {
			case marusia.OnStart:
				resp.Text = "Скилл запущен"
				resp.TTS = "Скилл запущен, жду команд"
			case "картинка":
				resp.Card = marusia.NewBigImage(
					"Заголовок",
					"Описание",
					457239017,
				)
			case "картинки":
				resp.Card = marusia.NewImageList(
					457239017,
					457239018,
				)
			case "кнопки":
				resp.Text = "Держи кнопки"
				resp.TTS = "Жми на кнопки"
				resp.AddURL("ссылка", "https://vk.com")
				resp.AddButton("подсказка без нагрузки", nil)
				resp.AddButton("подсказка с нагрузкой", myPayload{
					Text: "test",
				})
			case marusia.OnInterrupt:
				resp.Text = "Скилл закрыт"
				resp.TTS = "Пока"
				resp.EndSession = true
			default:
				resp.Text = "Неизвестная команда"
				resp.TTS = "Я вас не поняла"
			}
		case marusia.ButtonPressed:
			var p myPayload

			err := json.Unmarshal(r.Request.Payload, &p)
			if err != nil {
				resp.Text = "Что-то пошло не так"
				return
			}

			resp.Text = "Кнопка нажата. Полезная нагрузка: " + p.Text
			resp.TTS = "Вы нажали на кнопку"
		}

		return
	})

	http.HandleFunc("/", wh.HandleFunc)

	http.ListenAndServe(":8080", nil)
*/
package marusia // import "github.com/SevereCloud/vksdk/v2/marusia"

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"strconv"
)

// Version версия протокола.
const Version = "1.0"

// RequestType тип ввода.
type RequestType string

// Возможные значения.
const (
	SimpleUtterance RequestType = "SimpleUtterance" // голосовой ввод
	ButtonPressed   RequestType = "ButtonPressed"   //  нажатие кнопки
)

// NLU - Natural Language Understanding.
type NLU struct {
	Tokens   []string `json:"tokens"`
	Entities []string `json:"entities"`
}

// Типичные команды голосового ввода.
const (
	// OnStart команда запуска скилла. В скилл будет передана пустая строка
	// Command = "".
	OnStart = ""

	// OnInterrupt команда завершении скилла по команде "стоп", "выход" и т.д. в
	// скилл будет передано Command = "on_interrupt", чтобы у скилла была
	// возможность попрощаться с пользователем.
	OnInterrupt = "on_interrupt"
)

// RequestIn данные, полученные от пользователя.
type RequestIn struct {
	// Служебное поле: запрос пользователя, преобразованный для внутренней
	// обработки Марусей. В ходе преобразования текст, в частности, очищается
	// от знаков препинания, а числительные преобразуются в числа. При
	// завершении скилла по команде "стоп", "выход" и т.д. в скилл будет
	// передано "on_interrupt", чтобы у скилла была возможность попрощаться с
	// пользователем.
	Command string `json:"command"`

	// Полный текст пользовательского запроса, максимум 1024 символа.
	OriginalUtterance string `json:"original_utterance"`

	// Тип ввода.
	Type RequestType `json:"type"`

	// JSON, полученный с нажатой кнопкой от обработчика скилла (в ответе на
	// предыдущий запрос), максимум 4096 байт. Передаётся, только если была
	// нажата кнопка с payload.
	Payload json.RawMessage `json:"payload,omitempty"`

	// Объект, содержащий слова и именованные сущности, которые Маруся
	// извлекла из запроса пользователя.
	NLU NLU `json:"nlu"`
}

// Screen структура для Interfaces.
type Screen struct{}

// Interfaces интерфейсы, доступные на устройстве пользователя.
type Interfaces struct {
	// Пользователь может видеть ответ скилла на экране и открывать ссылки
	// в браузере.
	Screen *Screen `json:"screen,omitempty"`
}

// IsScreen пользователь может видеть ответ скилла на экране и открывать
// ссылки в браузере.
func (i *Interfaces) IsScreen() bool {
	return i.Screen != nil
}

// Meta информация об устройстве, с помощью которого пользователь общается
// с Марусей.
type Meta struct {
	// Идентификатор клиентского приложения
	ClientID string `json:"client_id"`

	// Язык в POSIX-формате, максимум 64 символа.
	Locale string `json:"locale"`

	// Название часового пояса, включая алиасы, максимум 64 символа
	Timezone string `json:"timezone"`

	// Интерфейсы, доступные на устройстве пользователя.
	Interfaces Interfaces `json:"interfaces"`

	// Город пользователя на русском языке.
	CityRu string `json:"_city_ru,omitempty"`
}

// Session данные о сессии.
type Session struct {
	// Уникальный идентификатор сессии, максимум 64 символа.
	SessionID string `json:"session_id"`

	// Идентификатор экземпляра приложения, в котором пользователь общается с
	// Марусей, максимум 64 символа.
	//
	// Deprecated: Важно! Это поле устарело, вместо него стоит использовать
	// Session.Application.ApplicationID.
	UserID string `json:"user_id"`

	// Идентификатор вызываемого скилла, присвоенный при создании.
	// Соответствует полю "Маруся ID" в настройках скилла.
	SkillID string `json:"skill_id"`

	// Признак новой сессии:
	//
	// true — пользователь начинает новый разговор с навыком,
	//
	// false — запрос отправлен в рамках уже начатого разговора.
	New bool `json:"new"`

	// Идентификатор сообщения в рамках сессии, максимум 8 символов.
	// Инкрементируется с каждым следующим запросом.
	MessageID int `json:"message_id"`

	// Данные о пользователе. Передаётся, только если пользователь
	// авторизован.
	User User `json:"user"`

	// Данные об экземляре приложения.
	Application Application `json:"application"`
}

// User данные о пользователе.
type User struct {
	// Идентификатор аккаунта пользователя (максимум 64 символа).
	// Уникален в разрезе: скилл + аккаунт.
	UserID string `json:"user_id"`
}

// Application данные об экземпляре приложения.
type Application struct {
	// Идентификатор экземпляра приложения, в котором пользователь общается
	// с Марусей (максимум 64 символа). Уникален в разрезе: скилл + приложение
	// (устройство).
	ApplicationID string `json:"application_id"`

	// Тип приложения (устройства). Возможные варинты: mobile, speaker, other.
	ApplicationType string `json:"application_type"`
}

// Request структура запроса.
type Request struct {
	// Информация об устройстве, с помощью которого пользователь общается с Марусей.
	Meta Meta `json:"meta"`

	// Данные, полученные от пользователя.
	Request RequestIn `json:"request"`

	// Данные о сессии.
	Session Session `json:"session"`

	// Версия протокола.
	Version string `json:"version"`
}

// BindingType тип для DefaultPayload.
type BindingType string

// Возможные значения.
const (
	BindingTypeSuggest BindingType = "suggest"
)

// DefaultPayload дефолтная нагрузка.
type DefaultPayload struct {
	BindingType    BindingType `json:"binding_type"`
	Index          int         `json:"index"`
	TargetPhraseID string      `json:"target_phrase_id"`
}

// Button кнопка.
type Button struct {
	// Текст кнопки, максимум 64 символа.
	Title string `json:"title"`

	// URL, который откроется при нажатии на кнопку, максимум 1024 байта.
	// Если свойство url не указано, по нажатию на кнопку навыку будет
	// отправлен текст кнопки. Пока кнопки с url поддерживаются только
	// на Android, на iOS появятся совсем скоро.
	URL string `json:"url,omitempty"`

	// Любой JSON, который нужно отправить скиллу, если данная кнопка будет
	// нажата, максимум 4096 байт.
	Payload interface{} `json:"payload,omitempty"`
}

// CardType тип карточки.
type CardType string

// Возможные значения.
const (
	// Одно изображение.
	BigImage CardType = "BigImage"

	// Набор изображений.
	ItemsList CardType = "ItemsList"

	// Карточка vk miniapp'а.
	MiniApp CardType = "MiniApp"
)

// CardItem элемент карточки.
type CardItem struct {
	// TODO: Заголовок изображения.
	// Title string `json:"title,omitempty"`

	// TODO: Описание изображения.
	// Description string `json:"description,omitempty"`

	// ID изображения из раздела "Медиа-файлы" в настройках скилла.
	ImageID int `json:"image_id"`
}

// Card описание карточки — сообщения с поддержкой изображений.
type Card struct {
	// Тип карточки.
	Type CardType `json:"type"`

	// Заголовок изображения.
	Title string `json:"title,omitempty"`

	// Описание изображения.
	Description string `json:"description,omitempty"`

	// ID изображения из раздела "Медиа-файлы" в настройках скилла
	// (игнорируется для типа ItemsList).
	ImageID int `json:"image_id,omitempty"`

	// Список изображений, каждый элемент является объектом формата BigImage.
	Items []CardItem `json:"items,omitempty"`

	// Ссылка для карточки типа Link. Для карточки типа MiniApp адрес мини-приложения.
	URL string `json:"url,omitempty"`
}

// NewBigImage возвращает карточку с картинкой.
func NewBigImage(title, description string, imageID int) *Card {
	return &Card{
		Type:        BigImage,
		Title:       title,
		Description: description,
		ImageID:     imageID,
	}
}

// NewItemsList возвращает карточку с набором картинок.
func NewItemsList(items ...CardItem) *Card {
	return &Card{
		Type:  ItemsList,
		Items: items,
	}
}

// NewImageList возвращает карточку с набором картинок.
func NewImageList(imageIDs ...int) *Card {
	items := make([]CardItem, len(imageIDs))

	for i := 0; i < len(imageIDs); i++ {
		items[i].ImageID = imageIDs[i]
	}

	return NewItemsList(items...)
}

// NewMiniApp возвращает карточку vk miniapp'а.
func NewMiniApp(url string) *Card {
	return &Card{
		Type: MiniApp,
		URL:  url,
	}
}

// Response данные для ответа пользователю.
type Response struct {
	// Текст, который следует показать и сказать пользователю. Максимум 1024
	// символа. Не должен быть пустым. В тексте ответа можно указать переводы
	// строк последовательностью «\n».
	//
	// TODO: поддержка массива строк.
	Text string `json:"text"`

	// Ответ в формате TTS (text-to-speech), максимум 1024 символа.
	// Поддерживается расстановка ударений с помощью '+'.
	TTS string `json:"tts,omitempty"`

	// Кнопки (suggest'ы), которые следует показать пользователю. Кнопки можно
	// использовать как релевантные ответу ссылки или подсказки для
	// продолжения разговора.
	Buttons []Button `json:"buttons,omitempty"`

	// Признак конца разговора:
	//
	// true — сессию следует завершить,
	//
	// false — сессию следует продолжить.
	EndSession bool `json:"end_session"`

	// Описание карточки — сообщения с поддержкой изображений.
	// Важно! Если указано данное поле, то поле text игнорируется.
	Card *Card `json:"card,omitempty"`
}

// AddURL добавляет к ответу кнопку с ссылкой.
func (r *Response) AddURL(title string, url string) {
	if r.Buttons == nil {
		r.Buttons = make([]Button, 0)
	}

	r.Buttons = append(r.Buttons, Button{
		Title: title,
		URL:   url,
	})
}

// AddButton добавляет к ответу кнопку с полезной нагрузкой.
//
// Если полезная нагрузка не нужна, можно передать nil.
func (r *Response) AddButton(title string, payload interface{}) {
	if r.Buttons == nil {
		r.Buttons = make([]Button, 0)
	}

	r.Buttons = append(r.Buttons, Button{
		Title:   title,
		Payload: payload,
	})
}

// responseSession данные о сессии.
type responseSession struct {
	SessionID   string      `json:"session_id"`
	MessageID   int         `json:"message_id"`
	UserID      string      `json:"user_id"`
	SkillID     string      `json:"skill_id"`
	New         bool        `json:"new"`
	User        User        `json:"user"`
	Application Application `json:"application"`
}

// response структура ответа серверу.
type response struct {
	Response Response        `json:"response"` // Данные для ответа.
	Session  responseSession `json:"session"`  // Данные о сессии.
	Version  string          `json:"version"`  // Версия протокола.
}

// Webhook структура.
type Webhook struct {
	event func(r Request) Response
}

// NewWebhook возвращает новый Webhook.
func NewWebhook() *Webhook {
	return &Webhook{}
}

// OnEvent обработчик скилла.
//
// Таймаут ожидания ответа — 5 секунд, после чего сервер Маруси завершит
// сессию.
func (wh *Webhook) OnEvent(f func(r Request) Response) {
	wh.event = f
}

// HandleFunc обработчик http запросов.
func (wh *Webhook) HandleFunc(w http.ResponseWriter, r *http.Request) {
	mediatype, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if mediatype != "application/json" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	resp := wh.event(req)

	fullResponse := response{
		Response: resp,
		Session: responseSession{
			SessionID:   req.Session.SessionID,
			MessageID:   req.Session.MessageID,
			UserID:      req.Session.UserID,
			SkillID:     req.Session.SkillID,
			New:         req.Session.New,
			User:        req.Session.User,
			Application: req.Session.Application,
		},
		Version: Version,
	}

	// Возвращаем данные
	w.Header().Add("Content-Type", "application/json; encoding=utf-8")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(fullResponse)
}

// SpeakerAudioVKID существует возможность вставлять в произносимую речь
// собственные звуки. Для этого необходимо на странице редактирования
// скилла воспользоваться формой загрузки медиафайлов. Загруженный аудиофайл
// будет доступен только для использования в вашем навыке.
func SpeakerAudioVKID(id string) string {
	return fmt.Sprintf(`<speaker audio_vk_id=%s>`, strconv.Quote(id))
}

// SpeakerAudio произносимый Марусей текст можно разнообразить звуковыми
// эффектами, которые входят в её библиотеку звуков.
//
// Список звуков можно найти на странице https://vk.com/dev/marusia_skill_docs4
func SpeakerAudio(name string) string {
	return fmt.Sprintf(`<speaker audio=%s>`, strconv.Quote(name))
}
