package rff

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/viveksahu26/testing_in_go/practice/eeee"
)

func ReadFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	eeee.PrintError(err)

	var space []string

	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		fmt.Printf("\n %d: ", i)
		fmt.Print(scanner.Text())
		getEachLine := scanner.Text()
		space = append(space, getEachLine)
		i++
	}

	err = file.Close()
	eeee.PrintError(err)

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return space, err
}
