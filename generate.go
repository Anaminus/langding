package main

import (
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Entry struct {
	Name        string    `toml:"name"`
	URL         string    `toml:"url"`
	Date        time.Time `toml:"date"`
	Syntax      bool      `toml:"syntax"`
	NoScroll    bool      `toml:"noscroll"`
	Interactive bool      `toml:"interactive"`
}

var Languages = map[string]Entry{}

func main() {
	if len(os.Args) < 4 {
		fmt.Println(`Usage:

go run generate.go [languages.toml] [template.gohtml] [output.html]`)
		return
	}

	_, err := toml.DecodeFile(os.Args[1], &Languages)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.New("langding").ParseFiles("index.gohtml"))

	out, err := os.Create(os.Args[3])
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = tmpl.ExecuteTemplate(out, os.Args[2], Languages)
	if err != nil {
		panic(err)
	}
}
