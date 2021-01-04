package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	ch := make(chan map[string]int)

	result1 := make(map[string]int)
	result2 := make(map[string]int)

	files := readDir("./texts")

	// measure execution time of normal method
	start1 := time.Now()
	for _, fi := range files {
		res := CountWordsNormal(fi)
		for k, v := range res {
			result1[k] += v
		}
	}
	end1 := time.Since(start1)

	// measure execution time of concurrent method
	start2 := time.Now()
	for _, fi := range files {
		go CountWordsConcurrently(fi, ch)

		wordsCount := <-ch
		for k, v := range wordsCount {
			result2[k] += v
		}
	}
	end2 := time.Since(start2)

	switch end1 < end2 {
	case true:
		fmt.Println("normal method faster: ", end2-end1)
	default:
		fmt.Println("concurrent method faster: ", end1-end2)
	}

	fmt.Printf("%#v\n", result1)
	fmt.Println(strings.Repeat("<>", 40))
	fmt.Println(strings.Repeat("<>", 40))
	fmt.Println(strings.Repeat("<>", 40))
	fmt.Printf("%#v\n", result2)

}

// CountWordsNormal implements normal method to count word occurrences in a file
func CountWordsNormal(fi string) map[string]int {
	wordsCount := make(map[string]int)

	file, err := os.Open(fi)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		key := scanner.Text()

		if _, ok := wordsCount[key]; ok {
			wordsCount[key]++
		} else {
			wordsCount[key] = 1
		}
	}

	return wordsCount

}

// CountWordsConcurrently implements concurrent method to count word occurrences in a file
func CountWordsConcurrently(fi string, c chan map[string]int) {

	wordsCount := make(map[string]int)

	file, err := os.Open(fi)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		key := scanner.Text()

		if _, ok := wordsCount[key]; ok {
			wordsCount[key]++
		} else {
			wordsCount[key] = 1
		}

	}

	c <- wordsCount
}

func readDir(path string) (arrFiles []string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	files, err := f.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "text") {
			arrFiles = append(
				arrFiles,
				fmt.Sprintf("%s/%s", path, file.Name()),
			)
		}

	}

	return arrFiles
}
