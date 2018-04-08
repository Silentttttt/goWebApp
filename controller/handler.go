package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//IndexHandler return the index.html to client
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route to /:", r.URL.Path)
	if strings.Compare(r.URL.Path, "/") != 0 {
		http.Error(w, "404 page not found", 404)
		return
	}

	w.Header().Set("Content-type", "text/html")
	indexPage, err := ioutil.ReadFile("bitlove/examples/parallax.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("index.html file error %v", err), 500)
	}
	fmt.Fprint(w, string(indexPage))
}

//Static maps static files
func Static(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route to static:", r.URL.Path)
	// Disable listing directories
	if strings.HasSuffix(r.URL.Path, "/") {
		http.Error(w, "404 page not found", 404)
		return
	}
	http.ServeFile(w, r, "bitlove" + r.URL.Path[0:])
}
