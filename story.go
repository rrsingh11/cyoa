package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)
type Story map[string]Chapter
type Chapter struct {
	Title   string   `json:"title"`
	Paragraphs   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:text`
	Arc string `json:"arc`
}

//  Create a html template
var htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>

    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}

    <ul>
        {{range .Options}}
        <li><a href=/{{.Arc}}>{{.Text}}</a></li>
		{{end}}
    </ul>

</body>
</html>`



func Exit(err string) {
	fmt.Println("Error: ", err)
	os.Exit(1)
}

func ParseJSON(f io.Reader) Story {
	var story Story
	dc := json.NewDecoder(f)
	if err := dc.Decode(&story); err != nil {
		Exit("Problem in Decoding")
	}
	return story
}	

func OpenFile(f string) *os.File {
	file, err := os.Open(f)
	if(err != nil) {
		Exit("Cannot open the file")
	}
	return file
}

func NewHandler(s Story) handler {
	return handler{s: s}
}

type handler struct {
	s Story
}

// We are handling routes here so not using mux
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New(" ").Parse(htmlTemplate))
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		if err := tpl.Execute(w, chapter); err != nil {
			Exit("Unable to Execute")
		}
	}
}
