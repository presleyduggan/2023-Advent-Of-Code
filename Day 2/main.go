package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		lines = append(lines, fileScanner.Text())
    }
  
    readFile.Close()

	return lines

}

func solvePartOne() {
	lines := readFile("input.txt")
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	RED_CUBES := 12
	GREEN_CUBES := 13
	BLUE_CUBES := 14

	sum := 0

	// lets go through until we find a digit. if it is higher than the lowest (red) we will keep going
	// if it is lower we will check the next characters to see if it is green, red, or blue and compare

	valid_game := true

	for i := 0; i < len(lines); i++ {

		// chars 0,1,2,3 are GAME
		// char 4 is space
		// char 5 is the start of the game id
		// game id is always i + 1
		game_id := i + 1

		start := 6
		
		// j will start at : so we don't start at the game id
		for q := start; q < len(lines[i]); q++{
			if string(lines[i][q]) == ":"{
				start = q
				break
			}
		}

		// start loop at 6
		for j := start; j < len(lines[i]); j++{
			char := lines[i][j]
			val, err := strconv.Atoi(string(char))
				if err == nil {
					// lets make sure we get all of the number
					line := lines[i]
					find_num := true
					add_mod := 0 // add to the string index
					for find_num {
						for k := j + 1; k < len(lines[i]); k++ {
							// check if next character is " "
							if string(line[k]) == " " {
								// we are done
								find_num = false
								break
							} else {
								// add the value val
								add_mod = add_mod + 1
								next_val, err := strconv.Atoi(string(line[k]))
								if(err == nil) {
									strs, err := strconv.Atoi(fmt.Sprint(val) + fmt.Sprint(next_val))
									if(err == nil){
										val = strs
									}
								}
								
							}
						}
					}

					// we have a number.. lets check if its greater than the minimum
					if(val > RED_CUBES ){
						// lets check based on what comes next
						// there will be a space, then we'll compare the next 3 to RED BLU or GRE
						next_word := string(line[j+ 2 + add_mod]) + string(line[j+3+add_mod]) + string(line[j+4+add_mod])
						switch next_word {
						case "red":
							if val > RED_CUBES {
								valid_game = false
							}
						case "blu":
							if val > BLUE_CUBES {
								valid_game = false
							}
						case "gre":
							if val > GREEN_CUBES {
								valid_game = false
							}
						default:
							// how did we get here?
							fmt.Printf("default?")
						}
					
					}
				}
		}


		// if valid game, add the id
		if valid_game {
			fmt.Println("valid game: ", game_id)
			fmt.Println("sum - ", sum)
			sum = sum + game_id
		}

		valid_game = true

	}

	fmt.Println("Part 1 Solution =")
	fmt.Println(sum)
}

func solvePartTwo(){
	lines := readFile("input.txt")
	sum := 0

	// lets go through until we find a digit. if it is higher than the lowest (red) we will keep going
	// if it is lower we will check the next characters to see if it is green, red, or blue and compare


	for i := 0; i < len(lines); i++ {

		// chars 0,1,2,3 are GAME
		// char 4 is space
		// char 5 is the start of the game id
		// game id is always i + 1
		game_id := i + 1

		start := 6

		red_nums := []int{}
		green_nums := []int{}
		blue_nums := []int{}
		
		// j will start at : so we don't start at the game id
		for q := start; q < len(lines[i]); q++{
			if string(lines[i][q]) == ":"{
				start = q
				break
			}
		}

		// start loop at 6
		for j := start; j < len(lines[i]); j++{
			char := lines[i][j]
			val, err := strconv.Atoi(string(char))
				if err == nil {
					// lets make sure we get all of the number
					line := lines[i]
					find_num := true
					add_mod := 0 // add to the string index
					for find_num {
						for k := j + 1; k < len(lines[i]); k++ {
							// check if next character is " "
							if string(line[k]) == " " {
								// we are done
								find_num = false
								break
							} else {
								// add the value val
								add_mod = add_mod + 1
								next_val, err := strconv.Atoi(string(line[k]))
								if(err == nil) {
									strs, err := strconv.Atoi(fmt.Sprint(val) + fmt.Sprint(next_val))
									if(err == nil){
										val = strs
									}
								}
								
							}
						}
					}

					// we have a number.. lets check if its greater than the minimum
						// lets check based on what comes next
						next_word := string(line[j+ 2 + add_mod]) + string(line[j+3+add_mod]) + string(line[j+4+add_mod])
						//fmt.Println("word = ", next_word)
						switch next_word {
						case "red":
							red_nums = append(red_nums, val)
						case "blu":
							blue_nums = append(blue_nums, val)
						case "gre":
							green_nums = append(green_nums, val)
						default:
							// how did we get here?
							fmt.Printf("default?")
						}
					
						j = j + add_mod
					
				}
		}


		// find the max of each set, multiply and add to sum
		
		fmt.Println("game: ", game_id, red_nums, green_nums, blue_nums)
		r_min := slices.Max(red_nums)
		g_min := slices.Max(green_nums)
		b_min := slices.Max(blue_nums)

		power := r_min * g_min * b_min


		//fmt.Println("sum - ", sum)
		sum = sum + power
		


	}

	fmt.Println("Part 2 Solution =")
	fmt.Println(sum)

}

func main() {
	solvePartOne()
	solvePartTwo()

	

}