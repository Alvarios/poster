package poster

import (
	"github.com/Alvarios/poster/get"
	"github.com/Alvarios/poster/post"
)

var Post = post.Post
var Get = get.Get

func main() {
	_ , _ = Post, Get
}