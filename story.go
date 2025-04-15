package main

import (
	"encoding/json"
	"io"
	"os"
)

func LoadJsonStory (file io.Reader) (Story, error){
	if file, err := os.Open("gopher.json"); err!=nil {
		return nil, err
	} else {
		decoder := json.NewDecoder(file)
		var story Story
		if err := decoder.Decode(&story); err != nil {
			return nil, err
		}
		return story, nil
	}
}
type Chapter struct {
	Title   string `json:"title"`
	Story []string `json:"story"`
	Options []Options `json:"options"`}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Story map[string]Chapter