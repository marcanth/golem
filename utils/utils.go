package utils

import (
	"os"
	"path/filepath"
)

func GetPaths(s string) []string {
	var paths []string = make([]string, 0)
	filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
		paths = append(paths, path)
		return err
	})
	return paths
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
