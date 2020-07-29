# Poster

A reverse http utility tool to easily send requests to other servers.

## Post

Send a post request from go, in `application/json` format.

```go
// payload is a map[string]interface{}
res, err := httpUtils.Post(url, payload)
```

## PostForm

Send a post request from go, with a form content.

```go
files := []httpUtils.File{
    {
        Path: '/path/to/my/file',
        Key: 'file'
    }
}

// payload is a map[string]interface{}
res, err := httpUtils.PostForm(url, payload, files)
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/poster/blob/master/LICENSE)
