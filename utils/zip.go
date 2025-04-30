package utils

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func ZipFolder(source, target string, exclude []string) error {
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	w := zip.NewWriter(out)
	defer w.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, ex := range exclude {
			if info.IsDir() && info.Name() == ex {
				return filepath.SkipDir
			}
			if !info.IsDir() && filepath.Base(filepath.Dir(path)) == ex {
				return nil
			}
		}
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		f, err := w.Create(rel)
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		_, err = io.Copy(f, in)
		return err
	})
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			err = os.MkdirAll(fpath, 0755)
			if err != nil {
				return err
			}
			continue
		}
		err = os.MkdirAll(filepath.Dir(fpath), 0755)
		if err != nil {
			return err
		}
		outFile, err := os.Create(fpath)
		if err != nil {
			return err
		}
		inFile, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}
		_, err = io.Copy(outFile, inFile)
		outFile.Close()
		inFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func UnzipFromEmbed(zipData []byte, dest string) error {
	r, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return err
	}

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			err = os.MkdirAll(fpath, 0755)
			if err != nil {
				return err
			}
			continue
		}
		err = os.MkdirAll(filepath.Dir(fpath), 0755)
		if err != nil {
			return err
		}
		outFile, err := os.Create(fpath)
		if err != nil {
			return err
		}
		inFile, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}
		_, err = io.Copy(outFile, inFile)
		outFile.Close()
		inFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
