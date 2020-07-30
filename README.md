# Poster

A reverse http utility tool to easily send requests to other servers.

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

## PostForm

Send a post request from go, with a form content.

```go
package my_package

import "github.com/Alvarios/poster"

func main() {
    payload := map[string]interface{}{
        "message": "hello from go",
    }

    files := []poster.PostFile{
        poster.PostFile{
            Path: "local/path/to/my/file",
            Key: "filename",
        },
    }

    // payload is a map[string]interface{}
    res, err := poster.PostForm("https://www.domain.com/url", payload, files)
}
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/poster/blob/master/LICENSE)
