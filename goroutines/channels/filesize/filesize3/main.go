package main

import (
	"book/goroutines/channels/filesize/du"
	"flag"
	"log"
	"sync"
	"time"
)

var progress = flag.Bool("v", false, "Show progress")

func main() {
	defer timeTrack(time.Now(), "CountingFilesWithWaitGroup")
	// determine the initial directories
	flag.Parse()
	roots := flag.Args()
	// if the len is 0 then we assume it's the current directory the program is being called from
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse the file tree
	fileSizes := make(chan int64)
	var nodesLeft sync.WaitGroup
	for _, root := range roots {
		nodesLeft.Add(1)
		go du.WorkDirWG(root, &nodesLeft, fileSizes)
	}

	go func() {
		nodesLeft.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *progress {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
	// loop label for breakage
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
	du.PrintDiskUsage(nfiles, nbytes) // final totals
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
