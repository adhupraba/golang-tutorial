package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html>
	<body>
		<div>This</div>
		<h1>My First Heading</h1>
		<p>My first paragraph.</p>
		<p>HTML images are defined with the img tag:</p>
		<img src="xxx.jpg" width="104" height="142">
		<img src="xxx.jpg" width="104" height="142">
	</body>
</html>
`

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}

	words, pics := countWordsAndImages(doc)
	fmt.Printf("words = %d, images = %d\n", words, pics)
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int

	visit(doc, &words, &pics)

	return words, pics
}

// words *int, pics *int
// words, pics *int
func visit(n *html.Node, words, pics *int) {
	// if it's an element node, what tag does it have?

	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}
}
