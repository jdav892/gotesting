package main

import (
	"os"
	"time"

	"github.com/jdav892/gotesting/math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
