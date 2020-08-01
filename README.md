# Poster

A reverse http utility tool to easily send requests to other servers.

```cgo
go get github.com/Alvarios/poster
```

## Post

Send a post request from go, in `application/json` format.

```go
package my_package

import "github.com/Alvarios/poster"

func main() {
    payload := map[string]interface{}{
        "message": "hello from go",
    }

    // payload is a map[string]interface{}
    res, err := poster.Post("https://www.domain.com/url", payload)
}
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/poster/blob/master/LICENSE)
