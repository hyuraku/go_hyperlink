package hyperlink

import "fmt"

func Hyperlink(url string, title string) map[string]string {
	html := fmt.Sprintf("html: <a href='%s'>%s</a>\n", url, title)
	markdown := fmt.Sprintf("markdown: [%s](%s)", title, url)
	results := map[string]string{
		"html": html,
		"markdown": markdown,
	}
	return results
}
