package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Result struct{
	Level string
}

func worker(id int, jobs <-chan string, results chan<-Result, wg *sync.WaitGroup){
	defer wg.Done()
	for line := range jobs {
		if strings.Contains(line, "ERROR"){
			results <-Result{Level: "ERROR"}
		} else {
			results <-Result{Level: "INFO"}
		}

	}

}

func main(){
	startTime := time.Now()
	file, err := os.Open("dummy_logs.txt")
	if err != nil{
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	jobs := make(chan string, 100)
	results := make(chan Result, 100)
	var wg sync.WaitGroup

	numWorkers := 5
	for i:=0; i<numWorkers; i++{
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	
	}
	go func(){
		wg.Wait()
		close(results)

	}()
	go func(){
		scanner := bufio.NewScanner(file)
		for scanner.Scan(){
			jobs <- scanner.Text()
		}
		close(jobs)
	}()

	errorCount := 0
	totalCount := 0
	for res := range results{
		totalCount++
		if res.Level == "ERROR"{
			errorCount++
		}
	}

	fmt.Printf("Processing took %s\n", time.Since(startTime))
	fmt.Printf("Total lines processed: %d\n", totalCount)	
	fmt.Printf("Error count: %d\n", errorCount)
	
}
