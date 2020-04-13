package filereader

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(path string, output chan<- []string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()
	defer close(output)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		cols := strings.Split(row, ",")
		output <- cols
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
