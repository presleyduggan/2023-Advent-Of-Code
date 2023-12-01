package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) []string {

	readFile, err := os.Open(fileName)
  
    if err != nil {
        fmt.Println(err)
		os.Exit(3) // failed
    }

	lines := []string{}

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
    }
  
    readFile.Close()

	return lines

}

func solvePartOne() {
	lines := readFile("input.txt")

	nums := []int{}

	for i := 0; i < len(lines); i++ {
		line_num := []int{}
		for j := 0; j < len(lines[i]); j++ {
			char := lines[i][j]
			val, err := strconv.Atoi(string(char))
			if err == nil {
				line_num = append(line_num, val)
			}
		}


		if len(line_num) > 1 {
			num, err := strconv.Atoi(fmt.Sprint(line_num[0]) + fmt.Sprint(line_num[len(line_num) - 1]))
			if err != nil{
				fmt.Println(err)
			}

			nums = append(nums, num)
		} else {
			num, err := strconv.Atoi(fmt.Sprint(line_num[0]) + fmt.Sprint(line_num[0]))
			if err != nil{
				fmt.Println(err)
			}

			nums = append(nums, num)
		}
		
	}

	solution := 0

	for i := 0; i < len(nums); i++{
		solution = solution + nums[i]
	}

	fmt.Println("Solution =")
	fmt.Println(solution)
}

// returns the int value of the word found
func checkForWord(str string) int{
	str_nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	for i := 0; i < len(str_nums); i++{
		if strings.Contains(str, str_nums[i]){
			return i + 1
		}
	}
	return -1
}

func solvePartTwo(){
	lines := readFile("input.txt")

	

	nums := []int{}

	for i := 0; i < len(lines); i++ {
		line_num := []int{}
		//for i, c := range lines[i]

		str := ""
		for j := 0; j < len(lines[i]); j++ {
			char := lines[i][j]

			str = str + string(char)

			// first check for word

			wordCheck := checkForWord(str)

			if wordCheck >= 0{
				// word found!
				line_num = append(line_num, wordCheck)
				str = string(char) // make sure to keep the last character in cases such as eighthree
			} else {
				// check for a digit
				val, err := strconv.Atoi(string(char))
				if err == nil {
					line_num = append(line_num, val)
					str = ""
				}
			}

		}

		// add number to the nums array

		if len(line_num) > 1 {
			num, err := strconv.Atoi(fmt.Sprint(line_num[0]) + fmt.Sprint(line_num[len(line_num) - 1]))
			if err != nil{
				fmt.Println(err)
			}

			nums = append(nums, num)
		} else {
			num, err := strconv.Atoi(fmt.Sprint(line_num[0]) + fmt.Sprint(line_num[0]))
			if err != nil{
				fmt.Println(err)
			}

			nums = append(nums, num)
		}
	}	

	solution := 0

	for i := 0; i < len(nums); i++{
		fmt.Println(nums[i], lines[i])
		solution = solution + nums[i] // lets see them side by side 
	}

	fmt.Println("Solution =")
	fmt.Println(solution)

}

func main() {
    //fmt.Println("Hello, World!")
	solvePartTwo()

	

}