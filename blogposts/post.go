package blogposts

import (
	"io"
)

func newPost(postFile io.Reader) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData)[7:]}
	return post, nil
}
