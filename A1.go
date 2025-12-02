/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-02
 * @fileoverview Calculate average mark.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// declare variables
	var marksString string
	var marksNumber float64
	var total float64 = 0.0

	reader := bufio.NewReader(os.Stdin)

	// input for how many marks there are
	fmt.Print("How many marks will you enter today? \n")
	marksString, _ = reader.ReadString('\n')
	marksString = strings.TrimSpace(marksString)
	marksNumber, _ = strconv.ParseFloat(marksString, 64)

	// for statement
	for markAmount := 0; markAmount < int(marksNumber); markAmount++ {
		var gradeString string
		var gradeNumber float64
		fmt.Printf("Enter mark %d", markAmount + 1)
		gradeString, _ = reader.ReadString('\n')
		gradeString = strings.TrimSpace(gradeString)
		gradeNumber, _ = strconv.ParseFloat(gradeString, 64)
		total += gradeNumber
	}

	// average calculation
	var average float64 = total / marksNumber
	fmt.Printf("\nAverage mark: %.2f\n", average )

	// statements
	if average <= 49 {
		fmt.Println("The student is failing.")
	} else if average >= 50 && average <= 69 {
		fmt.Println("The student's performance is just below average.")
	} else if average >= 70 && average <= 79 {
		fmt.Println("The student's performance is average.")
	} else if average >= 80 {
		fmt.Println("The student is on the honour roll.")
	}

	fmt.Println("\nDone.")
}
