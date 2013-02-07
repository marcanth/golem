package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	CONFIG_FILE = "config.json"
	SEPARATOR   = string(os.PathSeparator)
	DATE_FORMAT = "20060102150405"
)

type Config struct {
	Ignored List
	Destination string
}

func getConfig(s string) (c Config, err error) {
	_s := path.Clean(s)
	s = path.Join(_s, CONFIG_FILE)
	f, err := os.Open(s)
	if err == nil {
		d := json.NewDecoder(f)
		d.Decode(&c)
		c.Ignored = append(c.Ignored, CONFIG_FILE)
		for k, v := range c.Ignored {
			c.Ignored[k] = path.Join(_s, v)
		}
	}

	return
}

type List []string

func GetPaths(s string) (paths List) {
	var allPaths List

	s = path.Clean(s)
	filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
		allPaths = append(allPaths, path)
		return err
	})

	config, err := getConfig(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(config.Ignored)

	for k, v := range allPaths {
		fmt.Println(v)
		if config.Ignored.SearchBy(v) == false {
			paths = append(paths, allPaths[k])
		}
	}
	return
}

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

func ParseArgs(a []string) (args map[string]string) {
	args = make(map[string]string)
	switch len(a) {
	case 2:
		args["source"] = a[1]
		if c, err := getConfig(args["source"]); err == nil {
			args["dest"] = c.Destination
		} else {
			panic(err)
		}
	case 3:
		args["source"] = a[1]
		args["dest"] = a[2]
	}

	var zip string
	flag.StringVar(&zip, "o", time.Now().Format("DATE_FORMAT"), "Enter with the zip name here")
	args["zip"] = zip

	return
}
