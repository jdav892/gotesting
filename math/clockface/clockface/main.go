package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jdav892/gotesting/math/clockface"
)

func main() {
	t := time.Now()
	sh := clockface.SecondHand(t)
	io.WriteString(os.Stdout, svgStart+"\n")
	io.WriteString(os.Stdout, bezel+"\n")
	io.WriteString(os.Stdout, secondHandTag(sh)+"\n")
	io.WriteString(os.Stdout, svgEnd+"\n")
}

func secondHandTag(p clockface.Point) string {
	// format to 2 decimal places and make the tag self-closing
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.2f" y2="%.2f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
