/* ROY CHARLES
** croy@student.42.fr
** France
** N-Puzzle project
 */

package main

import (
	"algo"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"ui"
)

func CheckArgs() (int, int) {
	//check if 2 args
	if len(os.Args) != 3 {
		fmt.Println("[Usage]:", os.Args[0], "<file of puzzle> -[your algo]")
		return 1, 0
	}

	//Open file and check if exist.
	file, err := os.Stat(os.Args[1])
	if err != nil || file.IsDir() {
		fmt.Println("[ERROR]: Can't open your File/Directory.")
		return 1, 0
	}

	input, err := strconv.Atoi(string(os.Args[2]))
	if err != nil {
		fmt.Println("[ERROR]: Atoi failed for somes obscurs reasons...")
		return 1, 0
	}
	return 0, input
}

//	Read file line by line
func ReadFile() ([][]int, int, int) {
	var long, large int

	long, large = 0, 0

	//	Open file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("[ERROR]: Can't open your file.")
		return nil, 0, 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//	Take size of puzzle
	for long == 0 && scanner.Scan() {
		array := strings.Split(scanner.Text(), " ")
		if array[0][0] != '#' {
			long, large = ReadFirst(strings.Split(scanner.Text(), " "))
		}
	}
	if long < 1 || large < 1 {
		return nil, -1, -1
	}

	//	Save each piece of puzzle and check it
	tabl := SaveData(long, large, scanner)
	if CheckData(tabl, long, large) == false {
		return nil, -1, -1
	}

	return tabl, long, large
}

//	*	Save Data in tabl [][]int
func SaveData(long, large int, scanner *bufio.Scanner) [][]int {
	var i, j, blank, ll int
	var tabl [][]int

	tabl = make([][]int, long)
	ll = long * large
	for scanner.Scan() && long > 0 && large > 0 {
		array := strings.Split(scanner.Text(), " ")
		blank = 0
		for blank < len(array) && len(array[blank]) < 1 {
			blank++
		}
		if len(array) == blank || array[blank][0] == '#' {
			continue
		}
		tabl[i] = make([]int, large)
		for j = blank; j < len(array); j++ {
			if len(array[j]) < 1 {
				blank++
				continue
			}
			if array[j][0] == '#' {
				break
			}
			num, err := strconv.Atoi(array[j])
			if err != nil {
				fmt.Println("Converting data failed at line", i, "element", array[j])
				return nil
			}
			if j-blank == large {
				j = 0
				break
			}
			if num == 0 {
				num = ll
			}
			tabl[i][j-blank] = num
		}
		if j-blank != large {
			fmt.Println("Wrong number of colomn.")
			return nil
		}
		i++
		long--
	}
	if long != 0 {
		fmt.Println("Wrong number of line.")
		return nil
	}
	return tabl
}

//	*	*	Read first line
func ReadFirst(array []string) (int, int) {
	var long, large int

	for i := 0; i < len(array); i++ {
		if array[i][0] == '#' {
			break
		}
		num, err := strconv.Atoi(array[i])
		if err != nil {
			fmt.Println("Converting data failed at first line.")
			return -1, -1
		} else if num > 0 && long == 0 {
			long = num
		} else if num > 0 && i < 2 {
			large = num
		} else {
			fmt.Println("Bad data in first line.")
			return -1, -1
		}
	}
	if large == 0 {
		large = long
	}
	return long, large
}

//	*	Check data
func CheckData(tabl [][]int, long, large int) bool {
	var max, i, j, n int
	var check []int

	if tabl == nil || long < 1 || large < 1 {
		return false
	}
	max = long * large
	check = make([]int, max+1)
	for n = 0; n < max; n++ {
		check[n] = 0
	}
	for i = 0; i < long; i++ {
		for j = 0; j < large; j++ {
			n = tabl[i][j]
			if n < 1 || n > max {
				fmt.Println(n, "is not beetween 0 and", max)
				return false
			} else if check[n-1] == 1 {
				fmt.Println(check)
				fmt.Println("Repetitive value of", n)
				return false
			}
			check[n-1] = 1
		}
	}
	return true
}

func main() {
	err, AlGore := CheckArgs()

	if err == 0 {
		board, long, large := ReadFile()
		if board == nil || long < 1 || large < 1 {
		}
		array := algo.Pathfinding(board, long, large, AlGore)
		ui.Ui(array, long, large)
		return
	}
	return
}
