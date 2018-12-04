# Callback API

```go
package main

import (
    "net/http"
    "github.com/severecloud/vksdk/callback"
)

func main(){
    var cb callback.Callback
    cb.SecretKey := "abcdef"
    // cb.SecretKeys[<groupID>] = "abcdef"

    cb.MessageNew(func(obj MessageNewObject, groupID int) {
        //...
    })

    http.HandleFunc("/callback", cb.HandleFunc)

    http.ListenAndServe(":8080", nil)
}
```