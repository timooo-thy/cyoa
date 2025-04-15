package main

import (
	"net/http"
	"os"
	"strings"
)

func main() {

	mux := http.NewServeMux()
	file, err := os.Open("gopher.json")
	if err != nil {	
		panic(err)
	}	
	story, err := LoadJsonStory(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {	
		w.Write([]byte("Welcome to the Gopher Story! Head to /story/ to start your adventure.\n"))
	})

	mux.HandleFunc("/story/", func(w http.ResponseWriter, r *http.Request) {
		storyHandler(w, r, story)
	})

	http.ListenAndServe(":8080", mux)
}

func storyHandler(w http.ResponseWriter, r *http.Request, story Story) {
	// Load the JSON file
	arc := r.URL.Path[len("/story/"):]
	if arc == "" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Start with chapter intro at /story/intro.\n"))
		return
	}
	html, err := showChapter(arc, story)
	if err != nil {
		http.Error(w, "Error loading story", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func showChapter(chap string, story Story) (string, error) {
	chapter, ok := story[chap]
	if !ok {
		return "", nil
	}

	var sb strings.Builder

	sb.WriteString("<html><body>")
	sb.WriteString("<h1>" + chapter.Title + "</h1>")

	for _, line := range chapter.Story {
		sb.WriteString("<p>" + line + "</p>")
	}

	if len(chapter.Options) == 0 {
		sb.WriteString("<p><b>You reached the end of the story. Thanks for playing.</b></p>")
	} else {
		sb.WriteString("<h3>Options:</h3><ul>")
		for _, opt := range chapter.Options {
			sb.WriteString("<li><a href=\"/story/" + opt.Arc + "\">" + opt.Text + "</a></li>")
		}
		sb.WriteString("</ul>")
	}

	sb.WriteString("</body></html>")

	return sb.String(), nil
}
