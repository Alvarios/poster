package poster

import "github.com/Alvarios/poster/post"

var Post = post.Post
var PostForm = post.PostForm
type PostFile = post.File

func main() {
	_, _ = Post, PostForm
}