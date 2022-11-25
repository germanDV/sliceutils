package main

import (
	"fmt"
	"strings"

	su "gitlab.com/germandv/sliceutils"
)

type File struct {
	Path string
	Name string
	Ext  string
}

func main() {
	paths := []string{"index.html", "scripts/main.js", "styles/main.css", "styles/responsive.css"}
	files := []*File{}

	// Convert to File structs.
	su.ForEach(paths, func(s string) {
		files = append(files, &File{Path: s})
	})

	// Parse file name.
	su.ForEach(files, parseName)

	// Parse file extension.
	su.ForEach(files, parseExt)

	// Print.
	su.ForEach(files, display)
}

func parseName(f *File) {
	parts := strings.Split(f.Path, "/")
	f.Name = parts[len(parts)-1]
}

func parseExt(f *File) {
	parts := strings.Split(f.Name, ".")
	f.Ext = parts[len(parts)-1]
}

func display(f *File) {
	fmt.Printf("%+v\n", f)
}
