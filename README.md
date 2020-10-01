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

    // payload has to be JSON friendly
    res, err := poster.Post("https://www.domain.com/url", payload)
}
```

## Get

Get request to a http server.

```go
package my_package

import "github.com/Alvarios/poster"

func main() {
    res, err := poster.Get("https://www.domain.com/url")
}
```

## Copyright
[License MIT](https://github.com/Alvarios/poster/blob/master/LICENSE), licensed by Kushuh with Alvarios.
