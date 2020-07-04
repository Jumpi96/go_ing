package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	decoder "../../decoder"
	renderer "../../renderer"
)

func main() {
	arcName := ""

	if arcName == "" {
		arcName = "intro"
	}

	story, err := decoder.LoadStory("./gopher.json")
	if err != nil {
		panic(err)
	}

	h := newHandler(story)
	fmt.Println("Running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", h))
}

func newHandler(s decoder.Story) http.Handler {
	return handler{s}
}

type handler struct {
	s decoder.Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	arc, ok := h.s[path]
	if !ok {
		arc = h.s["intro"]
	}
	err := renderer.RenderArc(w, arc)
	if err != nil {
		panic(err)
	}
}
