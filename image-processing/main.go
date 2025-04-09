package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/disintegration/imaging"
)

type result struct {
	image *image.NRGBA
	path  string
	err   error
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("please, provide input and output directories!")
	}

	start := time.Now()

	err := setPipeline(os.Args[1], os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done! Time taken: %s\n", time.Since(start))
}

func setPipeline(root string, dstPath string) error {
	done := make(chan struct{})

	defer close(done)

	paths, errors := processFiles(done, root)

	results := processImage(done, paths)

	for result := range results {
		if result.err != nil {
			return result.err
		}
		saveImage(result, dstPath)
	}

	if err := <-errors; err != nil {
		return err
	}

	return nil
}

func processFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errors := make(chan error, 1)

	go func() {
		defer close(paths)

		errors <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			contentType, _ := getFileContentType(path)

			if !(contentType == "image/jpeg" || contentType == "image/jpg") {
				return nil
			}

			select {
			case paths <- path:
			case <-done:
				return fmt.Errorf("traverse was cancelled")
			}
			return nil
		})
	}()

	return paths, errors
}

func processImage(done <-chan struct{}, paths <-chan string) <-chan *result {
	var wg sync.WaitGroup

	routinesNum := runtime.NumCPU() * 4
	results := make(chan *result)

	wg.Add(routinesNum)
	for range routinesNum {
		go func() {
			for path := range paths {
				src, err := imaging.Open(path)

				if err != nil {
					select {
					case results <- &result{nil, path, err}:
					case <-done:
						return
					}
				}

				select {
				case results <- &result{image: imaging.Thumbnail(src, 100, 100, imaging.Lanczos), path: path}:
				case <-done:
					return
				}
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}

func saveImage(res *result, dstPath string) {
	fileName := filepath.Base(res.path)
	dstFullPath := dstPath + "/" + fileName

	imaging.Save(res.image, dstFullPath)
	fmt.Printf("%s -> %s\n", res.path, dstFullPath)
}

func getFileContentType(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}
	defer file.Close()

	contentTypeBuffer := make([]byte, 512)

	_, err = file.Read(contentTypeBuffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(contentTypeBuffer)

	return contentType, nil
}
