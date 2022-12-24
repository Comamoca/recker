package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// HACK: 取得したpathだと開けなかったのでハードコートした
// const ghqRoot = "/home/coma/ghq"

func ghqRoot() string {
	out, err := exec.Command("ghq", "root").Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Replace(string(out), "\n", "",  -1)
  
}

func ghqPath() string {
	out, err := exec.Command("ghq", "list").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func ghqList() []string {
	out, err := exec.Command("ghq", "list").Output()
	if err != nil {
		log.Fatal(err)
	}

	repoNames := strings.Split(string(out), "\n")

	return repoNames
}

func repoList(path string) []string {
	dirs, err := ioutil.ReadDir(ghqRoot())
	if err != nil {
		log.Fatal(err)
	}

	var results []string

	for _, dir := range dirs {
		if dir.IsDir() {
			results = append(results, dir.Name())
		}
	}

	return results
}

func READMEContent(name string) string {
  // TODO: READMEが含まれる文字列にしないとAsciiDocなどが表示されないので修正する
	READMEPath := filepath.Join(name, "README.md")

	f, err := os.Open(READMEPath)
	if err != nil {
		return "Not yet.."
	}
	defer f.Close()

	text, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	return string(text)
  }
	return string(text)
}

func cd(path string) {
	err := exec.Command("cd", path).Run()
	if err != nil {
		log.Fatal(err)
	}
}
