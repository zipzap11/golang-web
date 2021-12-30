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

func SimpleHtmlEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.html"))
	t.ExecuteTemplate(writer, "index.html", []string{"Template Success, Congratulations you are now one of nakama !!! yeayy", "second message"})
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
	mux.HandleFunc("/", SimpleHtmlFile)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}