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
	reqLocationSlice := strings.Split(r.URL.Path, "/")
	user := reqLocationSlice[2]
	project := reqLocationSlice[3]
	file := reqLocationSlice[4]

	tag, err := getTag(user, project)
	if err != nil {
		http.Error(w, err.Error(), 502)
	}

	fileUrl := ""
	if file == "archive.tar.gz" || file == "archive.zip" {
		parts := strings.Split(file, ".")
		atype := "tar.gz"
		if len(parts) == 2 {
			atype = "zip"
		}
		fileUrl = fmt.Sprintf("https://github.com/%v/%v/archive/%v.%v", user, project, tag, atype)
	}

	if fileUrl == "" {
		fileUrl = fmt.Sprintf("https://github.com/%v/%v/releases/download/%v/%v", user, project, tag, file)
	}

	http.Redirect(w, r, fileUrl, 302)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>gh-latest</h1><p>A simple project to get specific files from the latest release of any github repo.</p><p>Use a url like /repo/$user/$project/$filename to download that file from a github release.</p><p>Use 'archive.tar.gz' or 'archive.zip' as a filename to get the whole archive of the latest release.")
}

func getTag(user, project string) (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	url := fmt.Sprintf("https://github.com/%v/%v/releases/latest", user, project)
	fmt.Println(url)
	resp, err := client.Head(url)
	defer resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("Error retrieving latest location from github.com")
	}

	if resp.StatusCode != 302 {
		return "", fmt.Errorf("Unexpected status code from github.com")
	}

	redirLocation := resp.Header.Get(http.CanonicalHeaderKey("Location"))
	redirLocationSlice := strings.Split(redirLocation, "/")

	return redirLocationSlice[len(redirLocationSlice)-1], nil
}
