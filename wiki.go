package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	simple_page := &Page{Title: "simple_page", Body: []byte("This is a simple page")}
	simple_page.save()

	page, err := loadPage("simple_page")
  if err != nil {
    fmt.Println("Halt! Error!")
  }
	fmt.Println(string(page.Body))
}
