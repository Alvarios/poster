package poster

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"testing"
)

func ping (c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func postJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
	})
}

func post(c *gin.Context) {
	message := c.PostForm("message")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
	})
}

func postFile(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")

	if err != nil {
		c.String(500, err.Error())
	}

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": file.Filename,
	})
}

func setupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", ping)
	r.POST("/post", post)
	r.POST("/post_json", postJSON)
	r.POST("/post_file", postFile)

	return r
}

func checkPost(t *testing.T, url string, message map[string]string, files []File, expected string) {
	resp, err := PostForm(url, message, files)

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
