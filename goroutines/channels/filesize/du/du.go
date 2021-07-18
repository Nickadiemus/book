package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func WorkDir(dir string, fileSizes chan<- int64) {
	// check the entire directory
	for _, entry := range Dirents(dir) {
		if entry.IsDir() {
			// create the subdirectory
			subdir := filepath.Join(dir, entry.Name())
			WorkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func WorkDirWG(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	// check the entire directory
	for _, entry := range DirentsD(dir) {
		if entry.IsDir() {
			// increment the WaitGroup counter
			n.Add(1)
			// create the subdirectory
			subdir := filepath.Join(dir, entry.Name())
			go WorkDirWG(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// counting semaphore to help limit file opening
var sema = make(chan struct{}, 20)

// returns entries inside the directory
func Dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return nil
	}
	return entries
}

// returns entries inside the directory with
// sema
func DirentsD(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return nil
	}
	return entries
}

// Prints (in GB) number of files and total byte size
func PrintDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
