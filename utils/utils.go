package utils

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	CONFIG_FILE = "config.json"
	SEPARATOR   = string(os.PathSeparator)
)

func GetPaths(s string) (paths List) {
	var ap, ip = getAllPaths(s), getIgnoredPaths(s)
	for k, v := range ap {
		v = strings.Join(strings.Split(v, SEPARATOR)[1:], SEPARATOR)
		if ip.SearchBy(v) == false {
			paths = append(paths, ap[k])
		}
	}
	return
}

func getAllPaths(s string) (paths List) {
	s = path.Clean(s)
	filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
		paths = append(paths, path)
		return err
	})
	return
}

func getIgnoredPaths(src string) List {
	src = path.Clean(src)
	src = path.Join(src, CONFIG_FILE)

	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(f)

	var file File
	d.Decode(&file)

	return file.Ignored
}

type File struct {
	Ignored []string
}

type List []string

func (l List) SearchBy(n interface{}) bool {
	switch t := n.(type) {
	default:
		return false
	case string:
		for _, v := range l {
			if v == n {
				return true
			}
		}
	case []string:
		for _, v := range t {
			return l.SearchBy(v)
		}
	}

	return false
}
