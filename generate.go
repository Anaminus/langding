package main

import (
	"fmt"
	"html/template"
	"os"
	"slices"
	"strings"
)

type Entry struct {
	Name        string
	URL         string
	Date        string
	Syntax      bool
	NoScroll    bool
	Interactive bool
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(`Usage:

go run *.go [template.gohtml] [output.html]`)
		return
	}

	tmpl := template.Must(template.New("langding").ParseFiles(os.Args[1]))

	out, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer out.Close()

	slices.SortStableFunc(Languages, func(a, b Entry) int {
		return strings.Compare(a.Name, b.Name)
	})

	err = tmpl.ExecuteTemplate(out, os.Args[1], Languages)
	if err != nil {
		panic(err)
	}
}
