// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"os"
	"path/filepath"
)

// FileSystem struct
type FileSystem struct {
}

// NewFileSystem creates a new instance
func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

// FileExists reports whether the named file exists
func (fs *FileSystem) FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}

	return false
}

// DirExists reports whether the dir exists
func (fs *FileSystem) DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}

	return false
}

// EnsureDir ensures that directory exists
func (fs *FileSystem) EnsureDir(dirName string, mode int) error {
	err := os.MkdirAll(dirName, os.FileMode(mode))

	if err == nil || os.IsExist(err) {
		return nil
	}

	return err
}

// DeleteDir deletes a dir
func (fs *FileSystem) DeleteDir(dir string) error {
	err := os.RemoveAll(dir)

	if err != nil {
		return err
	}

	return nil
}

// StoreFile stores a file content
func (fs *FileSystem) StoreFile(path, content string) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, 0775)

	if err != nil {
		return err
	}

	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)

	return err
}
