package main

import (
	"regexp"
	"strings"
)

const indentSpaceAmount = 4

type CategoryParser struct {
	raw          string
	currPosition int
	buffer       string
}

func (cp *CategoryParser) Parse(raw string) *CategoryParser {
	cp.raw = raw
	return cp
}

func (cp *CategoryParser) Next() bool {
	if cp.currPosition >= len(cp.raw) {
		return false
	}

	var sb strings.Builder

	for ; cp.currPosition < len(cp.raw); cp.currPosition++ {
		currSymbol := string(cp.raw[cp.currPosition])
		if currSymbol != "\n" {
			sb.WriteString(currSymbol)
		} else {
			cp.currPosition++
			break
		}
	}

	trimmedLine := strings.Trim(sb.String(), " \r\n")
	if len(trimmedLine) == 0 {
		return cp.Next()
	}

	cp.buffer = sb.String()

	return true
}

func (cp *CategoryParser) Scan() string {
	return cp.buffer
}

func (cp *CategoryParser) ParseCategoryList(raw string) CategoryList {
	cp.Parse(raw)

	var (
		categories    CategoryList
		lastIndent    int
		categoryStack CategoryStack
	)

	for cp.Next() {
		regExp, err := regexp.Compile(`(\s*?)- \[(.+)]`)
		checkErr(err)

		matches := regExp.FindStringSubmatch(cp.buffer)
		if len(matches) < 3 {
			panic("list missed")
		}

		indent := len(matches[1])
		categoryName := matches[2]

		category := Category{CategoryName: categoryName}

		if indent == 0 {
			categories.Add(&category)
			categoryStack.Clear()
			categoryStack.Push(&category)
			lastIndent = indent
			continue
		}

		nesting := indent - lastIndent
		if nesting > 0 {
			categoryStack.Tail().AddSubCategory(&category)
			categoryStack.Push(&category)
		} else if nesting < 0 {
			categoryStack.Pop()
			for ; nesting != 0; nesting += indentSpaceAmount {
				categoryStack.Pop()
			}
			categoryStack.Tail().AddSubCategory(&category)
			categoryStack.Push(&category)
		} else {
			categoryStack.Pop()
			categoryStack.Tail().AddSubCategory(&category)
			categoryStack.Push(&category)
		}

		lastIndent = indent
	}

	return categories
}
