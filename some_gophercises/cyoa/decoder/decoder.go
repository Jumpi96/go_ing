package decoder

import (
	"encoding/json"
	"io/ioutil"
)

// Story is a struct containing everything to render a complete CYOA story.
type Story map[string]Arc

// Option to be chosen to get into the next arc.
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// Arc is a chapter into the story.
type Arc struct {
	Title   string   `json:"title"`
	Lines   []string `json:"story"`
	Options []Option `json:"options"`
}

// LoadStory gets the story from a file to a Story object.
func LoadStory(filename string) (story Story, err error) {
	base, _ := ioutil.ReadFile(filename)
	if err := json.Unmarshal(base, &story); err != nil {
		return nil, err
	}
	return story, nil
}
