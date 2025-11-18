// Author: Natnael Debesay
// Version: 1.0.0
// Date: 2025-11-18
// Fileoverview: This program will ask the user for their name, a aourse
// 		they are currently taking in school, and the last three test marks
//		for that course. The output of the program will be the student's nemt, 
//		their course, the three marks, and the average of the three test marks.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variables
	var userName string
	var courseName string
	var mark1AsString string
	var mark2AsString string
	var mark3AsString string
	var mark1AsNumber int
	var mark2AsNumber int
	var mark3AsNumber int
	var averageMark float64

	reader := bufio.NewReader(os.Stdin)

	// input
	fmt.Print("What is your name? ")
	userName, _ = reader. ReadString('\n')
	userName = strings.TrimSpace(userName)

	fmt.Print("What course are you currently taking? ")
	courseName, _ = reader.ReadString('\n')
	courseName = strings.TrimSpace(courseName)

	fmt.Print("Enter the  first test mark: ")
	mark1AsString, _ = reader.ReadString('\n')
	mark1AsString = strings.TrimSpace(mark1AsString)

	fmt.Print("Enter the  second test mark: ")
	mark2AsString, _ = reader.ReadString('\n')
	mark2AsString = strings.TrimSpace(mark2AsString)

	fmt.Print("Enter the third test mark: ")
	mark3AsString, _ = reader.ReadString('\n')
	mark3AsString = strings.TrimSpace(mark3AsString)

	// process
	mark1AsNumber, _ = strconv.Atoi(mark1AsString)
	mark2AsNumber, _ = strconv.Atoi(mark2AsString)
	mark3AsNumber, _ = strconv.Atoi(mark3AsString)
	averageMark = float64(mark1AsNumber+mark2AsNumber+mark3AsNumber) / 3.0
	// output
	fmt.Println()
	fmt.Printf("%s", "Hello, " + userName + "!\n")
	fmt.Printf("%s", "You are taking: " + courseName + ".\n")
	fmt.Printf("%s", "Your marks are: " + strconv.Itoa(mark1AsNumber) + ", " +
			strconv.Itoa(mark2AsNumber) + ", and " +
			strconv.Itoa(mark3AsNumber) + ".\n")
	fmt.Printf("%s", "Your average mark is: " +
			fmt.Sprintf("%.2f", averageMark) + ".\n")

	fmt.Println("\nDone.")
}
