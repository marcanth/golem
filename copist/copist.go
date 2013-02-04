package copist

import (
	"../utils"
	"bufio"
	"io"
	"os"
	"path"
	"strings"
)

func Copy(src, dst string) {
	paths := utils.GetPaths(src)

	for _, path := range paths {
		if stat, err := os.Stat(path); stat.IsDir() {
			if err != nil {
				panic(err)
			}
			os.Mkdir(changePathName(src, dst, path), stat.Mode().Perm())
		} else {
			i, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer i.Close()
			r := bufio.NewReader(i)

			o, err := os.Create(changePathName(src, dst, path))
			if err != nil {
				panic(err)
			}
			defer o.Close()
			w := bufio.NewWriter(o)
			buf := make([]byte, 1024)
			for {
				n, err := r.Read(buf)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}
				w.Write(buf[:n])
			}
			w.Flush()
		}
	}
}

func changePathName(src, dst, file string) string {
	src = path.Clean(src)
	dst = path.Clean(dst)

	return strings.Replace(file, src, dst, -1)
}
