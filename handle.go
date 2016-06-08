package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

var (
	indexTemplate = template.Must(template.New("index").Parse(tmplstr))
	picPathPrefix = "/pic/"
)

func handleIndex(title, root string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Be aggressive.
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		begin := time.Now()
		pics := read(root)
		err := indexTemplate.Execute(w, map[string]interface{}{
			"title": title,
			"pics":  pics,
		})

		log.Printf(
			"%s: %s %s: %d pic(s) in %s (%v)",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			len(pics),
			time.Since(begin),
			err,
		)
	}
}

func handlePic(root string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Don't permit non-images to be served
		if !isImage(filepath.Ext(r.URL.Path)) {
			http.NotFound(w, r)
			return
		}

		// Convert /pic/foo/bar.jpg to /path/to/root/foo/bar.jpg
		rel := strings.TrimPrefix(r.URL.Path, picPathPrefix)
		abs := filepath.Join(root, rel)
		http.ServeFile(w, r, abs)
	}
}
