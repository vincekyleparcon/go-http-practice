package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {

	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1></div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}
