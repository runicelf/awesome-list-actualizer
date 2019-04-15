package main

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	const awesomeListMarkdownRaw = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"

	req, err := http.NewRequest(http.MethodGet, awesomeListMarkdownRaw, nil)
	checkErr(err)

	client := http.Client{}
	res, err := client.Do(req)
	checkErr(err)

	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(res.Body)
	checkErr(err)

	body := buf.String()

	regExp, err := regexp.Compile(`(?s)(- \[Awesome Go].+?)##`)
	checkErr(err)

	matches := regExp.FindStringSubmatch(body)
	if len(matches) < 2 {
		panic("list missed")
	}

	cp := CategoryParser{}
	categories := cp.ParseCategoryList(matches[1])

	fmt.Println(categories)

	categories.Walk(func(c *Category) {
		fmt.Println(c.CategoryName)
	})

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
