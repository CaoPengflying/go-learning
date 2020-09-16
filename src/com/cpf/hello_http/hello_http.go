package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello world"))
	//})
	//http.HandleFunc("/time/", func(writer http.ResponseWriter, request *http.Request) {
	//	t := time.Now()
	//	timeStr := fmt.Sprintf("{\"time\": \"%s\"}",t)
	//	writer.Write([]byte(timeStr))
	//})
	//
	//http.ListenAndServe(":8888",nil)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8888", router))
}
