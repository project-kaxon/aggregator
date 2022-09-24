package aggregators

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/hyperupcall/knowledge/util"
)

func main() {
	name := "IETF"
	long := "Internet Engineering Task Force"

	var md strings.Builder
	md.WriteString("# IETF Standards\n\n")

	url := "https://www.rfc-editor.org/rfc-index.html"

	resp, err := http.Get(url)
	util.Handle(err)
	body, err := io.ReadAll(resp.Body)
	util.Handle(err)

	r := regexp.MustCompile("<noscript>([[:digit:]]+)</noscript>(?s:.*?)<b>(.*?)</b>(?s:(.*?))<")
	matches := r.FindAllStringSubmatch(string(body), 9500)
	for _, match := range matches {
		num := match[1]
		bold := match[2]
		// desc := match[3]

		fmt.Println(num)
		fmt.Println(bold)
		fmt.Println()

		n, err := strconv.Atoi(num)
		util.Handle(err)
		md.WriteString(fmt.Sprintf("- [http://www.rfc-editor.org/info/rfc%d](%s)\n", n, bold))
	}

	outFile, err := os.Create("specifications.md")
	util.Handle(err)
	defer outFile.Close()

	_, err = outFile.WriteString(md.String())
	util.Handle(err)
}
