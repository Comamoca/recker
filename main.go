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
	realPaths := func() []string {
		var list []string

		for _, repoName := range ghqList() {
			list = append(list, repoName)
		}
		return list
	}

	repos := packing(realPaths())

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

	fullPath := filepath.Join(ghqRoot(), repos[idx].path)

	if filepath.Base(fullPath) == "sandbox" {
		list := packing(sandboxList(fullPath))

		if filepath.Base(fullPath) == "sandbox" {
			idx, err := fuzzyfinder.Find(
				list,
				func(i int) string {
					return list[i].name
				},
				fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
					if i == -1 {
						return ""
					}

					if i == 0 {
						return ""
					}

					return fmt.Sprintf("Name: %s\n\n%s",
						repos[i].name, repos[i].readme)
				}))
			if err != nil {
			}

			fmt.Println(list[idx].path)
		}
	} else {
		fmt.Println(fullPath)
	}
}

func packing(repoPaths []string) []Repo {
	var repos []Repo

	for _, path := range repoPaths {
		readme := READMEContent(path)
		repo := Repo{
			name:   filepath.Base(path),
			path:   path,
			readme: readme,
		}

		repos = append(repos, repo)
	}
	return repos
}
