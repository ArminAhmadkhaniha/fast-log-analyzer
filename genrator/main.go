package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

func main(){
	file, err := os.Create("dummy_logs.txt")
	if err != nil{
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	var logLevel string

	for i := 0; i < 1000; i++{
		number := rand.Intn(100)
		if number < 90 {
			logLevel = "INFO"
		} else {
			logLevel = "ERROR"
		
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(file, "%s %s This is a log message number %d\n", now, logLevel, i)

		
		
	}
	fmt.Println("dummy_logs created")
}

