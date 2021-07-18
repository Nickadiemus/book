package main

import (
	"book/goroutines/channels/filesize/du"
	"flag"
	"log"
	"time"
)

var progress = flag.Bool("v", false, "Show progress")

func main() {
	defer timeTrack(time.Now(), "CountingFiles")
	// determine the initial directories
	flag.Parse()
	roots := flag.Args()
	// if the len is 0 then we assume it's the current directory the program is being called from
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse the file tree
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			du.WorkDir(root, fileSizes)
		}
		// make sure to close the channel
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *progress {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			du.PrintDiskUsage(nfiles, nbytes)
		}
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
