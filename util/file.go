package util

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
)

func ReadDir(path string) <-chan string {
	files := make(chan string)

	go func() {
		defer close(files)
		if err := fs.WalkDir(
			os.DirFS(path),
			".",
			func(path string, entry fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !entry.IsDir() {
					files <- path
				}
				return nil
			}); err != nil {
			panic(err)
		}
	}()

	return files
}

func CopyFile(srcDir, dstDir, file string) (err error) {
	src := path.Join(srcDir, file)

	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	dst := path.Join(dstDir, file)

	err = os.MkdirAll(path.Dir(dst), 0755)
	if err != nil {
		return
	}

	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}

		if os.SameFile(sfi, dfi) {
			return
		}
	}

	err = copyFileContents(src, dst)
	return
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()

	return
}
