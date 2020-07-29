package poster

import (
	"fmt"
	fileUtils "github.com/Alvarios/kushuh-go-utils/file-utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostForm(t *testing.T) {
	gin.SetMode(gin.ReleaseMode) // Avoid useless status messages in terminal (remove for advanced debug purposes)
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	tsUrl := func(p string) string {
		return fmt.Sprintf("%s/%s", ts.URL, p)
	}

	resp, err := http.Get(tsUrl("ping"))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
		return
	}

	if resp.StatusCode != 200 {
		t.Errorf("ping responded with wrong status %v", resp.StatusCode)
		return
	}

	root, err := fileUtils.FindProjectRoot("http-utils")
	if err != nil {
		t.Error(err.Error())
	}

	checkPost(
		t,
		tsUrl("post"),
		map[string]string{"message" : "hello world"},
		nil,
		"{\"message\":\"hello world\",\"status\":\"posted\"}",
	)

	checkPost(
		t,
		tsUrl("post_file"),
		nil,
		[]File{
			{
				Path: root + "/test.txt",
				Key:  "file",
			},
		},
		fmt.Sprintf("{\"message\":\"%s\",\"status\":\"posted\"}", root + "/test.txt"),
	)
}
