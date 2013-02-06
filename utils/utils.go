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
	var ignoredPaths, allPaths List

	s = path.Clean(s)
	filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
		allPaths = append(allPaths, path)
		return err
	})

	s = path.Join(s, CONFIG_FILE)
	f, err := os.Open(s)
	if err == nil {
		d := json.NewDecoder(f)
		var file File
		d.Decode(&file)
		ignoredPaths = file.Ignored
		ignoredPaths = append(ignoredPaths, CONFIG_FILE)
	}

	for k, v := range allPaths {
		v = strings.Join(strings.Split(v, SEPARATOR)[1:], SEPARATOR)
		if ignoredPaths.SearchBy(v) == false {
			paths = append(paths, allPaths[k])
		}
	}
	return
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
