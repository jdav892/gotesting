package blogrenderer_test

import (
	"bytes"
	"testing"
	"io"

	"github.com/jdav892/gotesting/blogrenderer"
	"github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		err := blogrenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post {
			Title: "hello world",
			Body: "This is a post",
			Description: "This is a description",
			Tags: []string{"go", "tdd"},
		}
	)

	for b.Loop(){
		blogrenderer.Render(io.Discard, aPost)
	}
}
