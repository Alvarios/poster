package post

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPost(t *testing.T) {
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

	checkPostJSON(t, tsUrl("post_json"), map[string]interface{}{"message" : "value"}, "{\"status\":\"posted\"}")
}
