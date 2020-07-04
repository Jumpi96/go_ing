package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestLoadExistingFile(t *testing.T) {
	file := "ex2.html"
	_, err := loadHTML(file)
	if err != nil {
		t.Errorf("The file %s wasn't loaded but the file does exist.", file)
	}
}

func TestLoadNonExistingFile(t *testing.T) {
	file := "ex22.html"
	_, err := loadHTML(file)
	if err == nil {
		t.Errorf("The file %s was loaded but the file doesn't exist.", file)
	}
}

func TestSimpleLink(t *testing.T) {
	content := "<html><a href=\"/other-page\">A link to another page</a></html>"
	links := make([]link, 0)
	tokenizer := html.NewTokenizer(strings.NewReader(content))

	findLinks(&links, tokenizer, false, link{})

	expected := link{Href: "/other-page", Text: "A link to another page"}
	if links[0] != expected {
		t.Errorf("Found: %+v. Expected: {%+v", links[0], expected)
	}
}

func TestLinksWithNodesInside(t *testing.T) {
	content := `
		<a href="https://www.twitter.com/joncalhoun">
		Check me out on twitter
		<i class="fa fa-twitter" aria-hidden="true"></i>
		</a>
	`
	links := make([]link, 0)
	tokenizer := html.NewTokenizer(strings.NewReader(content))

	findLinks(&links, tokenizer, false, link{})

	expected := link{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"}
	if links[0] != expected {
		t.Errorf("Found: %+v. Expected: {%+v", links[0], expected)
	}
}

func TestLinksWithCommentsInside(t *testing.T) {
	content := "<a href=\"/dog-cat\">dog cat <!-- commented text SHOULD NOT be included! --></a>"
	links := make([]link, 0)
	tokenizer := html.NewTokenizer(strings.NewReader(content))

	findLinks(&links, tokenizer, false, link{})

	expected := link{Href: "/dog-cat", Text: "dog cat"}
	if links[0] != expected {
		t.Errorf("Found: %+v. Expected: {%+v", links[0], expected)
	}
}
