package post

import (
	"fmt"
	"github.com/Alvarios/poster/_setup"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkPostJSON(t *testing.T, url string, c map[string]interface{}, expected string) {
	resp, err := Post(url, c)

	if err != nil {
		t.Errorf("error in post test : %s", err.Error())
		return
	}

	bd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("unable to read body of post : %s", err.Error())
		return
	}

	if resp.StatusCode != 200 {
		t.Errorf("post responded with wrong status %v - %s", resp.StatusCode, string(bd))
		return
	}

	if string(bd) != expected {
		t.Errorf("unexpected body of post\n\texpected %s\n\tgot %s", expected, string(bd))
		return
	}
}

func TestPost(t *testing.T) {
	gin.SetMode(gin.ReleaseMode) // Avoid useless status messages in terminal (remove for advanced debug purposes)
	ts := httptest.NewServer(_setup.SetupServer())
	defer ts.Close()

	tsUrl := func(p string) string {
		return fmt.Sprintf("%s/%s", ts.URL, p)
	}

	resp, err := http.Get(tsUrl("ping"))

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
		return
	}

	if resp.StatusCode != 200 {
		t.Errorf("ping responded with wrong status %v", resp.StatusCode)
		return
	}

	checkPostJSON(t, tsUrl("post_json"), map[string]interface{}{"message" : "value"}, "{\"status\":\"posted\"}")
}
