//+build windows

package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"syscall"
)

func (e *dirEntry) equals(other *dirEntry) bool {
	if e.path != other.path {
		fmt.Fprintf(os.Stderr, "%v: path does not match (%v != %v)\n", e.path, e.path, other.path)
		return false
	}

	if e.fi.Mode() != other.fi.Mode() {
		fmt.Fprintf(os.Stderr, "%v: mode does not match (%v != %v)\n", e.path, e.fi.Mode(), other.fi.Mode())
		return false
	}

	if !sameModTime(e.fi, other.fi) {
		fmt.Fprintf(os.Stderr, "%v: ModTime does not match (%v != %v)\n", e.path, e.fi.ModTime(), other.fi.ModTime())
		return false
	}

	return true
}

func nlink(info os.FileInfo) uint64 {
	return 1
}

func inode(info os.FileInfo) uint64 {
	return uint64(0)
}

func createFileSetPerHardlink(dir string) map[uint64][]string {
	linkTests := make(map[uint64][]string)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	for i uint64, f := range files {
		linkTests[i] = append(linkTests[i], f.Name())
		i++
	}
	return linkTests
}