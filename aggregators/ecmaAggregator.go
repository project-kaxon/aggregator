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

func ecmaAggregator() {
	var md strings.Builder
	md.WriteString("# ECMA Standards\n\n")

	url := "https://www.ecma-international.org/publications-and-standards/standards"

	resp, err := http.Get(url + "/?order=number")
	util.Handle(err)
	body, err := io.ReadAll(resp.Body)
	util.Handle(err)

	r := regexp.MustCompile(fmt.Sprintf("<a href=\"%s/(ecma-[[:digit:]]+)/\">(.*?)</a>(?:.*?)<a (?:.*?)>(.*?)<", url))
	matches := r.FindAllStringSubmatch(string(body), 150)
	for _, match := range matches {

		url := fmt.Sprintf("%s/%s", url, match[1])
		name := match[2]
		desc := match[3]

		fmt.Println(url)
		fmt.Println(name)
		fmt.Println(desc)
		fmt.Println()
		md.WriteString(fmt.Sprintf("- [%s](%s)\n  - %s\n", name, url, desc))
	}

	outFile, err := os.Create("specifications.md")
	util.Handle(err)
	defer outFile.Close()

	_, err = outFile.WriteString(md.String())
	util.Handle(err)
}
