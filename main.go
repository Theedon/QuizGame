package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func readFile(filePath *string, score *int) {
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

	for _, record := range records {
		scoreQuestion(record, score)
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
	var filePath string = "problems.csv"
	var score int = 0

	readFile(&filePath, &score)

}
