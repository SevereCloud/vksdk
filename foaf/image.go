package foaf // import "github.com/SevereCloud/vksdk/v2/foaf"

// Img - An image that can be used to represent some thing (ie. those
// depictions which are particularly representative of something, eg. one's photo on a homepage).
type Img struct {
	Image Image `xml:"Image"`
}

// Image digital images (such as JPEG, PNG, GIF bitmaps, SVG diagrams etc.).
type Image struct {
	Primary   string      `xml:"primary,attr"`
	Width     int         `xml:"width,attr"`
	Height    int         `xml:"height,attr"`
	About     string      `xml:"about,attr"`
	Thumbnail []Thumbnail `xml:"thumbnail"`
}

// Thumbnail a derived thumbnail image.
type Thumbnail struct {
	Width    int    `xml:"width,attr"`
	Height   int    `xml:"height,attr"`
	Resource string `xml:"resource,attr"`
}
