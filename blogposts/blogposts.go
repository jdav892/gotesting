package blogposts

import (
	"testing/fstest"
)

type Post struct{}

func NewPostsFromFs(fileSystem fstest.MapFS) []Post {
	return []Post{{}, {}}
}
