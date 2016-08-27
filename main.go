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
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	loc := strings.Split(r.URL.Path, "/")
	url := fmt.Sprintf("https://github.com/%v/%v/releases/latest", loc[2], loc[3])
	fmt.Println(url)
	resp, err := client.Head(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>gh-latest</h1><p>A simple project to get specific files from the latest release of any github repo.</p><p>Use a url like /repo/$user/$project/$filename to download that file from a github release.</p>")
}
