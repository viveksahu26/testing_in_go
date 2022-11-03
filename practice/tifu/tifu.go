package tifu

import (
	"bufio"
	"os"
	"strings"

	"github.com/viveksahu26/testing_in_go/practice/eeee"
)

func TakeInputFromUser() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	marks, err := reader.ReadString('\n')
	eeee.PrintError(err)

	// removes extra speces such as new line, tab, etc
	marks = strings.TrimSpace(marks)

	// convert string into number
	return marks, err
}
