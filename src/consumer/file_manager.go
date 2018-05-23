package consumer

import (
	"fmt"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func createFile() (file *os.File) {
	fileName := "programming_languages.txt"
	// Search for file
	_, err := os.Stat(fileName)

	// Delete file if already exists
	if err == nil {
		deleteFile(fileName)
	}

	// Create file
	file, err = os.Create(fileName)
	failOnError(err, "File cannot be created")
	fmt.Println("File", fileName, "was created successfull")

	return file
}

func deleteFile(fileName string) {
	var err = os.Remove(fileName)
	failOnError(err, "File cannot be deleted")
}
