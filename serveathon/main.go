package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// rewrite
		if path.Ext(r.URL.Path) == "" && !strings.HasSuffix(r.URL.Path, "/") {
			target := r.URL.Path + "/"
			if r.URL.RawQuery != "" {
				target += "?" + r.URL.RawQuery
			}
			http.Redirect(w, r, target, http.StatusTemporaryRedirect)
			return
		}
		http.ServeFile(w, r, path.Join("./", r.URL.Path))
	})
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
