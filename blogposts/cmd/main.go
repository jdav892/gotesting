package cmd

import (
	"log"
	"os"

	blogposts "github.com/quii/fstest-spike"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
