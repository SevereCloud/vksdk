# Callback API

```go
package main

import (
	"log"
	"net/http"

	"github.com/severecloud/vksdk/callback"
	"github.com/severecloud/vksdk/object"
)

func main() {
	var cb callback.Callback
	cb.SecretKey = "693d0ba9"
	// cb.SecretKeys[<groupID>] = "abcdef"

	cb.MessageNew(func(obj object.MessageNewObject, groupID int) {
		log.Print(obj.Text)
	})

	http.HandleFunc("/callback", cb.HandleFunc)

	http.ListenAndServe(":8080", nil)
}
```