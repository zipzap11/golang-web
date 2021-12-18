package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	// 	fmt.Fprint(writer, "hello world")
	// }

	router := http.NewServeMux()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "root url")
	})
	router.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "login page")
	})
	router.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "register page")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
