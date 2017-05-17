package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

const usage = `
usage:
	concur <data-dir-path> <search-string>
`

func processFile(filePath string, ch chan int) {
	//TODO: open the file, scan each line,
	//do something with the word, and write
	//the results to the channel

	//open the file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// use scanner to read line by line
	scanner := bufio.NewScanner(f)
	n := 0
	// for each line
	for scanner.Scan() {
		// increase the number of words
		n++
		for i := 0; i < 100; i++ {
			//hash the word
			h := sha256.New()
			//write it to the buffer
			h.Write(scanner.Bytes())
			//you use _ because you dont care about the result of the hash write, just how long it took
			_ := h.Sum(nil)
		}
	}
	//close the scanner
	f.Close()
	// give the input to the channel
	ch <- n
}

func processFileWord(filePath string, q string, ch chan []string) {
	//TODO: open the file, scan each line,
	//do something with the word, and write
	//the results to the channel

	//open the file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// use scanner to read line by line
	scanner := bufio.NewScanner(f)
	// for each line
	for scanner.Scan() {
		// increase the number of words

		//read each word
		word := scanner.Text()
		// if q is in the word,
		if strings.Contains(word, q) {
			// append to the slice
			matches = append(matches, word)
		}

	}
	//close the scanner
	f.Close()
	// give the input to the channel
	ch <- matches
}

func processDir(dirPath string, q string) {
	//TODO: iterate over the files in the directory
	//and process each, first in a serial manner,
	//and then in a concurrent manner
	//get the file
	fileinfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	//make a channel
	ch1 := make(chan []string, len(fileinfos))
	ch := make(chan int)
	//process each file concurrently
	for _, fi := range fileinfos {
		go processFileWord(path.Join(dirPath, fi.Name()), q, ch1)
	}
	//to process number of words
	nWords := 0
	totalMatches := []string{}
	for i := 0; i < len(fileinfos); i++ {
		nWords += <-ch
	}
	// to process matching words
	for i := 0; i < len(fileinfos); i++ {
		matches := <-ch1
		totalMatches = append(totalMatches, matches...)
	}
	fmt.Printf("processed %d words\n", nWords)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	dir := os.Args[1]
	q := os.Args[2]

	fmt.Printf("processing directory %s...\n", dir)
	start := time.Now()
	processDir(dir, q)
	fmt.Printf("completed in %v\n", time.Since(start))
}
