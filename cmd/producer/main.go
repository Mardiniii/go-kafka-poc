package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Mardiniii/go-kafka-poc/src/producer"
)

func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	fileName := os.Args[1]

	fileHandle, err := os.Open(fileName)
	failOnError(err, "Cannot open the file given")
	defer fileHandle.Close()

	fmt.Println("Reading content from:", fileName)
	fileScanner := bufio.NewScanner(fileHandle)

	// Publish a message for each line into the file
	for fileScanner.Scan() {
		producer.Start(fileScanner.Text())
	}

}
