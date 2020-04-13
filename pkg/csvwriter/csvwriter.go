package csvwriter

import (
	"encoding/csv"
	"log"
	"os"
)

func Write(payload [][]string, fname string) {
	file, err := os.Create(fname)
	checkError("Cannot create file", err)

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range payload {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
