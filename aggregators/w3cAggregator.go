package aggregators

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/hyperupcall/knowledge/util"
)

// NOTE: DOES NOT INCLUDE SUPERSEDED STANDARDS TODO

func w3cAggregator() {
	name := "W3C"
	long := "World Wide Web Consortium"

	url := "https://www.w3.org/TR"
	var md strings.Builder

	md.WriteString("# W3C\n\n")

	resp, err := http.Get(url)
	util.Handle(err)

	body, err := io.ReadAll(resp.Body)
	util.Handle(err)

	r := regexp.MustCompile("<li data-title=(?s:.*?)class=\"profile\">(.*?)<(?s:.*?)<a href=\"(.*?)\"(?s:.*?)title=\"(?:.*?)\">(.*?)<(?s:.*?)class=deliverer>(.*?)<")
	matches := r.FindAllStringSubmatch(string(body), 200)
	for _, match := range matches {
		docType := match[1]
		link := match[2]
		title := match[3]

		fmt.Println(docType)
		fmt.Println(link)
		fmt.Println(title)
		fmt.Println()

		md.WriteString(fmt.Sprintf("- [%s](%s) (%s)\n", title, link, docType))
	}

	outFile, err := os.Create("specifications.md")
	util.Handle(err)
	defer outFile.Close()

	_, err = outFile.WriteString(md.String())
	util.Handle(err)
}
