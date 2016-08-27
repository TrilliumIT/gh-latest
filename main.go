package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Starting...")
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/repo/", handleRepo)
	http.ListenAndServe(":8080", nil)
}

func handleRepo(w http.ResponseWriter, r *http.Request) {
	loc := strings.Split(r.URL.Path, "/")
	resp, err := http.Head(fmt.Sprintf("https://github.com/%v/%v/releases/latest", loc[1], loc[2]))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>gh-latest</h1><p>A simple project to get specific files from the latest release of any github repo.</p><p>Use a url like /repo/$user/$project/$filename to download that file from a github release.</p>")
}
