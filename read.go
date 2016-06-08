package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/btree"
)

const maxPics = 100

func read(root string) []pic {
	t := btree.New(2)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if isImage(filepath.Ext(path)) {
			t.ReplaceOrInsert(fts{path, info.ModTime()})
		}
		return nil
	})
	pics := []pic{}
	for _, filename := range flatten(t, maxPics) {
		f, err := os.Open(filename)
		if err != nil {
			log.Printf("%s: %v", filename, err)
			continue
		}
		config, _, err := image.DecodeConfig(f)
		if err != nil {
			log.Printf("%s: %v", filename, err)
			f.Close()
			continue
		}
		f.Close()
		pics = append(pics, pic{
			Source: picPathPrefix + strings.TrimPrefix(filename, root),
			Height: config.Height,
			Width:  config.Width,
		})
	}
	return pics
}

type pic struct {
	Source string
	Height int
	Width  int
}

func isImage(ext string) bool {
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	}
	return false
}

type fts struct {
	f  string
	ts time.Time
}

func (a fts) Less(b btree.Item) bool {
	return a.ts.After(b.(fts).ts) // newest first
}

func flatten(t *btree.BTree, max int) (res []string) {
	var count int
	t.Ascend(func(i btree.Item) bool {
		res = append(res, i.(fts).f)
		count++
		return count < max
	})
	return res
}
