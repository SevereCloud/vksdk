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
https://skill-debugger.marusia.mail.ru/

Если тестируется навык, развернутый локально, отладчик скорее всего
столкнется с ограничениями, накладываемыми браузером на выполнение CORS.

В этом случае необходимо добавить в webhook поддержку предварительного
запроса (метод OPTIONS) и проброс CORS-заголовков:

	wh.EnableDebuging()

Пример

Пример скилла.


	type myPayload struct {
		Text string
		marusia.DefaultPayload
	}

	wh := marusia.NewWebhook()
	// wh.EnableDebuging()

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
			case "ссылка":
				resp.Text = marusia.CreateDeepLink(
					"e7a7d540-3928-4f11-87bf-a0de1244c096",
					map[string]string{"Text": "нагрузка из ссылки"},
				)
				resp.TTS = "Держи диплинк"
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
		case marusia.DeepLink:
			var p myPayload

			err := json.Unmarshal(r.Request.Payload, &p)
			if err != nil {
				resp.Text = "Что-то пошло не так"
				return
			}

			resp.Text = "Специальная ссылка. Полезная нагрузка: " + p.Text
			resp.TTS = "Вы перешли по ссылке"
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
	"strconv"
)

// RequestType тип ввода.
type RequestType string

// Возможные значения.
const (
	SimpleUtterance RequestType = "SimpleUtterance" // голосовой ввод
	ButtonPressed   RequestType = "ButtonPressed"   //  нажатие кнопки
	DeepLink        RequestType = "DeepLink"
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
	NLU *NLU `json:"nlu"`
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

// State данные состояния.
type State struct {
	// Хранение данных в сессии.
	Session json.RawMessage `json:"session"`

	// Персистентное хранение данных.
	User json.RawMessage `json:"user"`
}

// Request структура запроса.
type Request struct {
	// Информация об устройстве, с помощью которого пользователь общается с Марусей.
	Meta Meta `json:"meta"`

	// Данные, полученные от пользователя.
	Request RequestIn `json:"request"`

	// Данные о сессии.
	Session Session `json:"session"`

	// Хранение состояния.
	State State `json:"state"`

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

// Response данные для ответа пользователю.
type Response struct {
	// Текст, который следует показать и сказать пользователю. Максимум 1024
	// символа. Не должен быть пустым. В тексте ответа можно указать переводы
	// строк последовательностью «\n».
	//
	// TODO: поддержка массива строк.
	Text string `json:"text"`

	// Для того, чтобы передать SSML в ответе из внешнего скилла,
	// необходимо передать "ssml".
	TTSType string `json:"tts_type,omitempty"`

	// Speech Synthesis Markup Language представляет собой основанный
	// на XML язык разметки для приложений синтеза речи.
	SSML string `json:"ssml,omitempty"`

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

	// Для сохранения состояния внутри сессии.
	// При этом, если в очередном ответе не записать данные, даже если они не
	// изменились, то они затрутся, и в следующем запросе поле будет пустым.
	// Помимо этого, состояние потеряется если:
	//
	// Пользователь выходит из скилла;
	//
	// Скилл сам явно завершает работу, передав EndSession: true;
	//
	// Выход происходит по таймауту, когда пользователь не отвечает некоторое
	// время (1 минуту).
	//
	// Лимит размера json-объекта - 5 КБ.
	SessionState json.RawMessage `json:"session_state,omitempty"`

	// Для персистентного хранения данных о юзере.
	// Чтобы удалить конкретное поле из сохранённого json-объекта, нужно
	// положить null в это поле.
	// Лимит размера json-объекта - 5 КБ.
	UserStateUpdate json.RawMessage `json:"user_state_update,omitempty"`
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
