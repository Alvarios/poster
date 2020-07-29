package poster

import (
	"bytes"
	"fmt"
	fileUtils "github.com/Alvarios/kushuh-go-utils/file-utils"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
)

type File struct {
	Path string
	Key  string
}

func PostForm(u string, d map[string]string, f []File) (*http.Response, error) {
	b := bytes.NewBuffer(nil)
	w := multipart.NewWriter(b)

	for k, v := range d {
		fw, err := w.CreateFormField(k)
		if err != nil {
			return nil, err
		}

		if _, err = fw.Write([]byte(v)); err != nil {
			return nil, err
		}
	}

	for _, g := range f {
		if g.Path == "" {
			return nil, fmt.Errorf("missing Path in file struct")
		}

		if g.Key == "" {
			return nil, fmt.Errorf("missing Key in file struct")
		}

		fw, err := w.CreateFormFile(g.Key, g.Path)
		if err != nil {
			return nil, err
		}

		fd, err := fileUtils.Upsert(g.Path)
		if err != nil {
			return nil, err
		}

		// Write file field from file to upload
		_, err = io.Copy(fw, fd)
		if err != nil {
			return nil, err
		}

		_ = fd.Close()
	}

	_ = w.Close()

	req, err := http.NewRequest("POST", u, b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		rb, err := httputil.DumpRequest(req, true)
		if err != nil {
			return nil, err
		}

		err = fmt.Errorf("bad status : %s\n%s\n%s", res.Status, string(rb), b.String())
		return res, err
	}

	return res, nil
}
