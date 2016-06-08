package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		title = flag.String("title", "Pica", "title of website")
		root  = flag.String("root", mustGetwd(), "path containing images")
		addr  = flag.String("addr", ":6174", "HTTP listen address")
	)
	flag.Parse()

	log.Printf("%q serving from %s", *title, *root)
	log.Printf("listening on %s", *addr)
	http.HandleFunc("/", handleIndex(*title, *root))
	http.HandleFunc(picPathPrefix, handlePic(*root))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func mustGetwd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
