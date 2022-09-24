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

func WhatwgAggregator() {
	name := "WHATWG"
	long := "Web Hypertext Application Technology Working Group"

	var md strings.Builder
	md.WriteString("# WhatWG\n\n")

	url := "https://spec.whatwg.org"

	resp, err := http.Get(url)
	util.Handle(err)
	body, err := io.ReadAll(resp.Body)
	util.Handle(err)

	r := regexp.MustCompile("<dt><a href=\"(.*?)\">(.*?)<(?s:.*?)<p>(.*?)</p>")
	matches := r.FindAllStringSubmatch(string(body), 200)
	for _, match := range matches {
		url := strings.TrimSuffix(match[1], "/")
		name := match[2]
		desc := match[3]

		fmt.Println(name)
		fmt.Println(url)
		fmt.Println(desc)
		fmt.Println()

		md.WriteString(fmt.Sprintf("- [%s](%s)\n", name, url))
		md.WriteString(fmt.Sprintf("  - %s\n", desc))
	}

	outFile, err := os.Create("specifications.md")
	util.Handle(err)
	defer outFile.Close()

	_, err = outFile.WriteString(md.String())
	util.Handle(err)
}
