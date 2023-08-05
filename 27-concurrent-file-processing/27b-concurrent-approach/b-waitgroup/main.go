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
func processFiles(paths <-chan string, pairs chan<- pair, done chan<- bool) {
	for path := range paths {
		pairs <- hashFile(path)
	}

	done <- true
}

func searchTree(dir string, paths chan<- string, wg *sync.WaitGroup) error {
	defer wg.Done()

	visit := func(path string, d os.DirEntry, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err
		}

		info, _ := d.Info()

		if info.IsDir() && path != dir {
			wg.Add(1)
			go searchTree(path, paths, wg)

			return filepath.SkipDir
		}

		if info.Mode().IsRegular() && info.Size() > 0 {
			paths <- path
		}

		return nil
	}

	return filepath.WalkDir(dir, visit)
}

func run(dir string) results {
	workers := 2 * runtime.GOMAXPROCS(0)

	paths := make(chan string)
	pairs := make(chan pair)
	done := make(chan bool)
	result := make(chan results)
	wg := new(sync.WaitGroup)

	for i := 0; i < workers; i++ {
		go processFiles(paths, pairs, done)
	}

	// we need another goroutine so we don't block here
	go collectHashes(pairs, result)

	// multi threaded walk of the directory tree
	// we need a waitGroup because we don't know how many to wait for
	wg.Add(1)

	if err := searchTree(dir, paths, wg); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	// we must close the paths channel so the workers stop
	close(paths)

	for i := 0; i < workers; i++ {
		<-done
	}

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
