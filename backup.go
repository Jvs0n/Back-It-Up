package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func BackupFolder(src, dst string, verbose bool) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return copyFile(src, filepath.Join(dst, filepath.Base(src)), verbose)
	}

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, _ := filepath.Rel(src, path)
		targetPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			err := os.MkdirAll(targetPath, info.Mode())
			if err != nil {
				return err
			}
			return nil
		}

		return copyFile(path, targetPath, verbose)
	})
}

func copyFile(srcFile, dstFile string, verbose bool) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	err = os.MkdirAll(filepath.Dir(dstFile), 0755)
	if err != nil {
		return err
	}
	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	if verbose {
		fmt.Println("Copied:", srcFile, "->", dstFile)
	}
	return nil
}
