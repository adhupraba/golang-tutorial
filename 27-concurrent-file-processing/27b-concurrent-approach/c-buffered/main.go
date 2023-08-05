package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type pair struct {
	hash, path string
}

type fileList []string

type results map[string]fileList

// compute hash
func hashFile(path string) pair {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return pair{fmt.Sprintf("%x", hash.Sum(nil)), path}
}

// collector
func collectHashes(pairs <-chan pair, result chan<- results) {
	hashes := make(results)

	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}

	result <- hashes
}

// call hashFiles to compute the hash of a file
func processFile(path string, pairs chan<- pair, wg *sync.WaitGroup, limits chan bool) {
	defer wg.Done()

	limits <- true

	defer func() {
		<-limits
	}()

	pairs <- hashFile(path)
}

// `<-chan` can read from channel
// `chan<-` can write into channel
// `limits chan bool` can read from as well as write into channel
func searchTree(dir string, pairs chan<- pair, wg *sync.WaitGroup, limits chan bool) error {
	defer wg.Done()

	visit := func(path string, d os.DirEntry, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err
		}

		info, _ := d.Info()

		if info.IsDir() && path != dir {
			wg.Add(1)
			go searchTree(path, pairs, wg, limits)

			return filepath.SkipDir
		}

		if info.Mode().IsRegular() && info.Size() > 0 {
			wg.Add(1)
			go processFile(path, pairs, wg, limits)
		}

		return nil
	}

	limits <- true

	defer func() {
		<-limits
	}()

	return filepath.WalkDir(dir, visit)
}

func run(dir string) results {
	workers := 2 * runtime.GOMAXPROCS(0)

	limits := make(chan bool, workers)
	pairs := make(chan pair)
	result := make(chan results)
	wg := new(sync.WaitGroup)

	// we need another goroutine so we don't block here
	go collectHashes(pairs, result)

	// multi threaded walk of the directory tree
	// we need a waitGroup because we don't know how many to wait for
	wg.Add(1)

	if err := searchTree(dir, pairs, wg, limits); err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	// by closing pairs we signal that all the hashes
	// have been collected; we have to do it here AFTER
	// all the workers are done
	close(pairs)

	return <-result
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing parameter, provide dir name!")
	}

	start := time.Now()

	if hashes := run(os.Args[1]); hashes != nil {
		for hash, files := range hashes {
			if len(files) > 1 {
				fmt.Println(hash[len(hash)-7:], len(files))

				for _, file := range files {
					fmt.Println("    ", file)
				}
			}
		}
	}

	t := time.Since(start).Round(time.Millisecond)

	fmt.Printf("total time taken is %s\n", t)
}
