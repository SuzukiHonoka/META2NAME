package main

import (
	"fmt"
	"github.com/dhowden/tag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	scanALL = true
)

var music = []string{"mp3", "flac", "wav"}

func main() {
	cwd, _ := os.Getwd()
	audios, _ := ioutil.ReadDir(cwd)
	for _, f := range audios {
		name := f.Name()
		if !scanALL {
			if isvalid("music", name) {
				continue
			}
		}
		data, _ := os.Open(name)
		m, err1 := tag.ReadFrom(data)
		if err1 != nil {
			fmt.Println(name, "Not a audio file.")
			continue
		}
		abss := filepath.Join(cwd, name)
		dname := ""
		if !scanALL {
			dname = filepath.Join(cwd, m.Artist()+" - "+m.Title()+filepath.Ext(name))
		} else {
			dname = filepath.Join(cwd, m.Artist()+" - "+m.Title()+"."+strings.ToLower(string(m.FileType())))
		}
		err2 := os.Rename(abss, dname)
		if err2 == nil {
			fmt.Println(name, "renamed to", dname)
		}
		_ = data.Close()
	}
}

func isvalid(ext string, name string) bool {
	fext := filepath.Ext(name)
	if len(fext) > 1 {
		lext := strings.ToLower(fext[1:])
		if ext == "music" {
			for _, f := range music {
				if lext != f {
					continue
				}
				return true
			}
		}
	}
	return false
}
