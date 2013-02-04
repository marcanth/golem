package zipper

import (
	"../utils"
	"archive/zip"
	"io"
	"os"
	"path"
)

func Zip(src, dst string) {
	if _, err := os.Stat(src); err == nil {
		paths := utils.GetPaths(src)

		fout, err := os.Create(path.Clean(dst))
		if err != nil {
			panic(err)
		}
		defer fout.Close()

		w := zip.NewWriter(fout)
		defer w.Close()

		for _, file := range paths {
			fin, err := os.Open(file)
			if err != nil {
				panic(err)
			}
			defer fin.Close()

			if s, _ := fin.Stat(); s.IsDir() == false {
				f, err := w.Create(fin.Name())
				if err != nil && err != io.EOF {
					panic(err)
				}
				_buf := make([]byte, 1024)
				for {
					n, err := fin.Read(_buf)
					if err != nil && err != io.EOF {
						panic(err)
					}
					if n == 0 {
						break
					}
					f.Write(_buf[:n])
				}
			}
		}
	}
}
