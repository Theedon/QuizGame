package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func readFile(filePath *string) [][]string {
	usageString := "input the file path of the csv file"
	filePtr := flag.String("file", *filePath, usageString)

	file, err := os.Open(*filePtr)

	if err != nil {
		log.Fatalf("Error while reading the %v file\n %v\n", filePath, err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records", err)
	}

	return records

}

func administerTest(ctx context.Context, records [][]string, score *int) {

OuterLoop:
	for _, record := range records {
		select {
		case <-ctx.Done():
			fmt.Println("Time elapsed...")
			break OuterLoop
		default:
			scoreQuestion(record, score)

		}

	}

	printScore(*score, len(records))

}

func scoreQuestion(record []string, score *int) {

	fmt.Printf("%v = ", record[0])
	scanner.Scan()
	text := scanner.Text()
	if text == record[1] {
		*score++
	}
}

func printScore(score int, len int) {

	perc := float64(score) / float64(len) * 100

	fmt.Printf("\nFinal Score = %v\nTest Percentage = %.2f\n", score, perc)
}

func main() {
	var score int = 0
	var testTime int
	var filePath string

	flag.IntVar(&testTime, "time", 30, "input the time")
	flag.StringVar(&filePath, "path", "problems.csv", "Enter the path to the csv file to be used for quiz")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(testTime)*time.Second)
	defer cancel()
	
	records := readFile(&filePath)
	go func() {
		administerTest(ctx, records, &score)
	}()



	<-ctx.Done()
	fmt.Println("\nTime elapsed...")
	printScore(score, len(records))
	os.Exit(0)
}
