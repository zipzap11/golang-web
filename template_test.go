package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	temp, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}

	temp.ExecuteTemplate(writer, "SIMPLE", "Hallo Dek!")
}

func SimpleHtmlFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.html"))
	t.ExecuteTemplate(writer, "index.html", []string{"Template Success, Congratulations you are now one of nakama !!! yeayy", "second message"})
}


//go:embed templates/*.html
var templates embed.FS

type Person struct {
	Name string
	Role string
}
func SimpleHtmlEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.html"))
	t.ExecuteTemplate(writer, "name.html", map[string]interface{}{
		"name": "Francisco",
		"role": "Software Engineeer",
	})
}

func SimpleHtmlData(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(writer, "name.html", Person{
		Name: "Ahmad Dhaniel",
		Role: "Singer",
	})
}

func SimpleHtmlAction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.html"))
	t.ExecuteTemplate(writer, "comparator.html", map[string]interface{}{
		"first": 9,
		"text": "haha",
	})
}

func SimpleHtmlIteration(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/iterator.html"))
	t.ExecuteTemplate(writer, "iterator.html", map[string]interface{}{
		"Users": []map[string]interface{}{
			{
				"Name": "Lorem",
				"Hobby": "Ping",
			},
			{
				"Name": "Ipsum",
				"Hobby": "Pong",
			},
		},
	})
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestServerTemplate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", SimpleHtmlIteration)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}