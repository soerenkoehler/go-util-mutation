package util

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/bmatcuk/doublestar/v4"
)

func GlobCopy(srcDir, dstDir, pattern string) (err error) {
	files, err := doublestar.Glob(os.DirFS(srcDir), pattern)
	for _, file := range files {
		if err == nil {
			if isDir, copyErr := copyFile(srcDir, dstDir, file); isDir {
				err = GlobCopy(srcDir, dstDir, path.Join(file, "**"))
			} else {
				err = copyErr
			}
		}
	}
	return
}

func copyFile(srcDir, dstDir, file string) (isDir bool, err error) {
	Debug("src=%v dst=%v file=%v", srcDir, dstDir, file)

	src := path.Join(srcDir, file)

	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if sfi.IsDir() {
		return true, nil
	}
	if !sfi.Mode().IsRegular() {
		return false, fmt.Errorf(
			"CopyFile: non-regular source file %s (%q)",
			sfi.Name(),
			sfi.Mode().String())
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
			return false, fmt.Errorf(
				"CopyFile: non-regular destination file %s (%q)",
				dfi.Name(),
				dfi.Mode().String())
		}

		if os.SameFile(sfi, dfi) {
			return
		}
	}

	return false, copyFileContents(src, dst)
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
