package main

import (
	"log"
	"net/http"
	//TODO: a name to replace foreverWall

	"github.com/Silentttttt/goWebApp/controller"
)

// var (
// 	workDir = "/users/ys/go/src/github.com/silent/foreverWall/bitlove/asserts"
// )

func main() {
	//	TODO: add a route package
	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/dist/", controller.Static)
	http.HandleFunc("/examples/", controller.Static)
	http.HandleFunc("/vendors/", controller.Static)
	http.HandleFunc("/jquery.fullPage.css", controller.Static)
	http.HandleFunc("/jquery.fullPage.js", controller.Static)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server Exit!")
}
