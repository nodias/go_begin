package main

import (
	"log"
	"path/filepath"
	"sync"
)

func main(){
	defer recoverPanic()

	var urls = []string{
		"http://127.0.0.1:7000/static/img/cat01.jpg",
		"http://127.0.0.1:7000/static/img/cat02.jpg",
		"http://127.0.0.1:7000/static/img/cat03.jpg",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls{
		go func(url string){
			defer wg.Done()
			Must(download(url))
		}(url)
	}
	wg.Wait()

	filenames, err := filepath.Glob("*.jpg")
	if err != nil {
		panic(err)
	}

	err = writeZip("images.zip", filenames)
	if err != nil {
		log.Fatal(err)
	}
}

