package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {

	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	// 	fmt.Println("Hello")
	// 	http.HandleFunc("/", getRoot)
	// 	http.HandleFunc("/hello", getHello)

	// 	err := http.ListenAndServe(":3333", nil)

	// 	if errors.Is(err, http.ErrServerClosed) {
	// 		fmt.Printf("Server Closed\n")
	// 	} else if err != nil {
	// 		fmt.Printf("error stating server: %s\n", err)
	// 		os.Exit(1)
	// 	}
	// }

	// func getRoot(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Printf("got / request\n")
	// 	io.WriteString(w, "This is my website\n")
	// }

	// func getHello(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Printf("got /hello request\n")
	// 	io.WriteString(w, "Hello, HTTP!\n")
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
