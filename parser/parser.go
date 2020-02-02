package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type link struct {
	Href string
	Text string
}

func main() {
	exampleFile := flag.String("example", "ex1.html", "a existent example file in examples/")
	flag.Parse()

	content, err := loadHTML(*exampleFile)

	if err != nil {
		panic(err)
	}

	links := make([]link, 0)
	tokenizer := html.NewTokenizer(content)

	findLinks(&links, tokenizer, false, link{})
	fmt.Println(links)

}

func findLinks(links *[]link, tokenizer *html.Tokenizer, opened bool, currentTag link) {
	token := tokenizer.Next()
	if token == html.ErrorToken {
		return
	} else if token == html.StartTagToken && !opened {
		tag := tokenizer.Token()
		if tag.Data == "a" {
			for _, a := range tag.Attr {
				if a.Key == "href" {
					currentTag.Text = ""
					currentTag.Href = a.Val
					break
				}
			}
			opened = true
		}
	} else if token == html.TextToken && opened {
		currentTag.Text += strings.Trim(tokenizer.Token().Data, " \n	")
	} else if token == html.EndTagToken && opened {
		if tokenizer.Token().Data == "a" {
			*links = append(*links, currentTag)
			opened = false
		}
	}
	findLinks(links, tokenizer, opened, currentTag)
}

func loadHTML(path string) (*os.File, error) {
	return os.Open("./examples/" + path)
}
