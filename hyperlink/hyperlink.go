package hyperlink

import "fmt"

type Format struct {
	Name string
	Url  string
}

func Hyperlink(url string, title string) []Format {
	html := fmt.Sprintf("<a href='%s'>%s</a>\n", url, title)
	markdown := fmt.Sprintf("[%s](%s)", title, url)
	results := []Format{
		{"html", html},
		{"markdown", markdown},
	}
	return results
}
