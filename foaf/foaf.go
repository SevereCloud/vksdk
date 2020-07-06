/*
Package foaf wrapper for FOAF.

Specification https://web.archive.org/web/20140909053226/http://api.yandex.ru/blogs/doc/indexation/concepts/what-is-foaf.xml
*/
package foaf // import "github.com/SevereCloud/vksdk/foaf"

import (
	"context"
	"encoding/xml"
	"net/http"

	"github.com/SevereCloud/vksdk/internal"
)

// BaseURL url foaf.
const BaseURL = "https://vk.com/foaf.php"

// Context keys to use with https://golang.org/pkg/context
// WithValue function to associate.
const (
	HTTPClient = internal.HTTPClientKey
	UserAgent  = internal.UserAgentKey
)

// rdf is a standard model for data interchange on the Web.
//
// See https://www.w3.org/RDF/
type rdf struct {
	XMLName xml.Name `xml:"RDF"`
	Lang    string   `xml:"lang,attr"`
	Rdf     string   `xml:"rdf,attr"`
	Rdfs    string   `xml:"rdfs,attr"`
	Foaf    string   `xml:"foaf,attr"`
	Ya      string   `xml:"ya,attr"`
	Img     string   `xml:"img,attr"`
	Dc      string   `xml:"dc,attr"`
	Person  Person   `xml:"Person"`
	Group   Group    `xml:"Group"`
}

// URI struct.
type URI struct {
	Primary  string `xml:"primary,attr"`
	Resource string `xml:"resource,attr"`
}

// Weblog struct.
type Weblog struct {
	Title    string `xml:"title,attr"`
	Resource string `xml:"resource,attr"`
}

// Date may be used to express temporal information at any level of granularity.
//
// Use time.Parse(time.RFC3339, v.Date).
type Date struct {
	Date string `xml:"date,attr"`
}

// getFoaf return RDF.
//
// BUG: VK return invalid XML char (example &#12;).
func getFoaf(ctx context.Context, req *http.Request) (r rdf, err error) {
	resp, err := internal.DoRequest(ctx, req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)

	// decoder.Strict set to false, the parser allows input containing
	// common mistakes:
	//	* If an element is missing an end tag, the parser invents
	//	  end tags as necessary to keep the return values from Token
	//	  properly balanced.
	//	* In attribute values and character data, unknown or malformed
	//	  character entities (sequences beginning with &) are left alone.
	decoder.Strict = false
	decoder.CharsetReader = internal.CharsetReader // For support WINDOWS-1251
	err = decoder.Decode(&r)

	return
}
