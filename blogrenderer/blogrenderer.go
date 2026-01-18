package blogrenderer

import "io"

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func Render(w io.Writer, p Post) error {
	return nil
}
