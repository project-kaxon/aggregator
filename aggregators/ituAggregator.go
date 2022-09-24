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

func ituAggregator() {
	name := "ITU"
	long := "International Telecommunication Union"

	var md strings.Builder
	md.WriteString("# ITU Standards\n\n")

	url := "https://www.itu.int/rec/T-REC-"
	letters := "ABCDEFGHIJKLMNOPQRSTUVXYZ"

	regex := "<strong>(.*?)</strong>(?s:(?:.*?))<p>(.*?)<"
	for _, letter := range letters {
		md.WriteString(fmt.Sprintf("## %c\n\n", letter))

		resp, err := http.Get(url + string(letter))
		util.Handle(err)
		body, err := io.ReadAll(resp.Body)
		util.Handle(err)

		r := regexp.MustCompile(regex)
		matches := r.FindAllStringSubmatch(string(body), 200)
		for _, match := range matches {
			// TODO: desc broken
			title := match[1]
			// desc := match[2]

			fmt.Println(match[1])
			// fmt.Println(match[2])
			fmt.Println()

			md.WriteString(fmt.Sprintf("- [%s]()\n", title))
		}
		md.WriteString("\n")
	}

	outFile, err := os.Create("specifications.md")
	util.Handle(err)
	defer outFile.Close()

	_, err = outFile.WriteString(md.String())
	util.Handle(err)
}
