package marusia // import "github.com/SevereCloud/vksdk/v2/marusia"

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

	// Стилизованная ссылка.
	Link CardType = "Link"
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

	// Заголовок изображения или ссылки.
	Title string `json:"title,omitempty"`

	// Описание ссылки.
	Text string `json:"text,omitempty"`

	// Описание изображения.
	//
	// Deprecated: исчезло из документации.
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

// NewLink возвращает карточку с стилизованной ссылкой.
func NewLink(url, title, text string, imageID int) *Card {
	return &Card{
		Type:    Link,
		URL:     url,
		Title:   title,
		Text:    text,
		ImageID: imageID,
	}
}
