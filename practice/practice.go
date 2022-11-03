// to main package
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/viveksahu26/testing_in_go/practice/eeee"
	"github.com/viveksahu26/testing_in_go/practice/rff"
	"github.com/viveksahu26/testing_in_go/practice/tifu"
)

func main() {
	fmt.Println("Hello World !!")
	fmt.Println("Hello Vivek !!")

	now := time.Now()
	year := now.Year()
	y, m, d := now.Date()

	fmt.Println("Year is: ", year)
	fmt.Println("Today's date is: ", d, m, y)

	random := "g#################"
	fmt.Println("random string: ", random)
	fmt.Println("Replace x#--> o")

	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(random)

	fmt.Println("fixed string: ", fixed)

	fmt.Println("--------------------------------------FAIL/PASS-----------------------------------")
	fmt.Print("Enter your grade in 10th board: ")

	marks, err := tifu.TakeInputFromUser()
	eeee.PrintError(err)

	// convert string into number
	marksInFloat, err := strconv.ParseFloat(marks, 64)
	eeee.PrintError(err)

	fmt.Println("The criteria for passing marks is more than 60%")

	if marksInFloat >= 95 {
		fmt.Println("PASS with Excellent marks: Topper")
	} else if marksInFloat > 60 {
		fmt.Println("You are PASS")
	} else {
		fmt.Println("You are FAIL")
	}
	fmt.Println("----------------------------- END -----------------------------")

	fmt.Println("Let's check whether file is present or not")
	fmt.Print("Enter the file name: ")

	file, err := tifu.TakeInputFromUser()
	eeee.PrintError(err)

	fileInfo, err := os.Stat(file)
	eeee.PrintError(err)

	fmt.Println("Size of file is: ", fileInfo.Size())
	fmt.Println("Is file Dir: ", fileInfo.IsDir())
	fmt.Println("File Name: ", fileInfo.Name())

	fmt.Println("***********************************READING************************************")

	fmt.Print("Enter fileName for Scanning: ")
	fileName, err := tifu.TakeInputFromUser()
	eeee.PrintError(err)

	// return an array
	myarray, err := rff.ReadFromFile(fileName)
	eeee.PrintError(err)
	fmt.Println("\n myarray: ", myarray)
}
