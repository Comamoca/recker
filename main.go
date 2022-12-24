package main

import (
	"fmt"
	"path/filepath"

	"github.com/ktr0731/go-fuzzyfinder"
)

type Repo struct {
	name   string
	path   string
	readme string
}

func tp() {
	fmt.Println(len(ghqList()))

	for _, name := range ghqList() {
		fmt.Println(name)
	}
}

func main() {
	root := ghqRoot()
	repoPaths := ghqList()

	var repos = []Repo{}

	for _, path := range repoPaths {
		realPath := filepath.Join(root, path)
		readme := READMEContent(realPath)
		repo := Repo{
			name:   path,
			path:   path,
			readme: readme,
		}

		repos = append(repos, repo)
	}

	idx, err := fuzzyfinder.Find(
		repos,
		func(i int) string {
			return repos[i].name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Name: %s\n\n%s",
				repos[i].name, repos[i].readme)
		}))

	if err != nil {
	}
	fmt.Println(filepath.Join(ghqRoot(), repos[idx].path))
}

func launchFuzzy() {
}
