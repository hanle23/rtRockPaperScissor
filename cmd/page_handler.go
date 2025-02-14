package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

var validPath = func() *regexp.Regexp {
	f, err := getRootPath()
	if err != nil {
		log.Fatal("Unable to get templates entries from path:", err)
	}
	fp := f + "/views/templates"
	entries, err := getEntriesFromPath(fp)
	if err != nil {

		log.Fatal("Unable to get templates entries from path:", err)
	}
	var templateList []string
	for _, e := range entries {
		templateList = append(templateList, strings.Replace(e.Name(), ".html", "", -1))
	}
	var str strings.Builder
	for i := 0; i < len(templateList); i++ {
		str.WriteString(templateList[i])
		if i < len(templateList)-1 {
			str.WriteString("|")
		}
	}
	pathPattern := fmt.Sprintf("^/(%s)/[a-zA-Z0-9]+)$", str.String())
	return regexp.MustCompile(pathPattern)
}()

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil
}
