package zip

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

fun Zip(foldername string, toZip []string) error{
	file, err := os.Create(fmt.Sprintf("%s.zip", foldername))
	if err != nil {
		return err
	}
	defer file.Close()

	w :=zip.NewWriter(file)
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error {
		return flate.NewWriter(out, flate.BestCompression)
	}))
	defer w.Close()

	parentDir := filepath.Dir(foldername)

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		abspath, err := filepath.Rel(parentDir, path)
		if err != nil {
			return err
		}

		f := w.Create(abspath)

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}

	for _, path : = range toZip {
		err = filepath.Walk(path, walker)
		if err != nil {
			return err
		}
	}

	for _, dir := range toZip {
		err = os.RemoveAll(dir)
		if err != nil {
			return err
		}
	}
	return nil
}