package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
        //fmt.Println(fileScanner.Text())
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
		//for i, c := range lines[i]
		for j := 0; j < len(lines[i]); j++ {
			char := lines[i][j]
			//fmt.Printf("%c\n", char)
			val, err := strconv.Atoi(string(char))
			if err == nil {
				line_num = append(line_num, val)
			}
		}

		//fmt.Println(line_num)

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

func solvePartTwo(){
	lines := readFile("test-input2.txt")

	str_nums := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	nums := []int{}

	for i := 0; i < len(lines); i++ {
		line_num := []int{}
		//for i, c := range lines[i]
		for j := 0; j < len(lines[i]); j++ {
			char := lines[i][j]


			// TODO: do tomorrow :)
			// we need to keep checking until there is a num, or a word that is a number
			// then clear buffer and keep going until we find another word or num
			// repeat until line is done

		}
	}	
}

func main() {
    //fmt.Println("Hello, World!")
	solvePartOne()

	

}