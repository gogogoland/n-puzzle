/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ConvertBoard.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbalea <tbalea@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/04/10 13:55:46 by tbalea            #+#    #+#             */
/*   Updated: 2016/04/10 13:55:46 by tbalea           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package algo

/*import "fmt"*/

var (
	obv   int
	snail []int
	right []int
)

//	Functions
//	*	Mapping snails value Amd save missing value
func SaveSnail(board [][]int, long, large int) {
	x, y := 0, 0
	i, min, max := 1, 0, long*large
	snail = make([]int, max+1)
	right = make([]int, max+1)

	snail[0] = 0
	right[0] = 0
	for i <= max {
		for i <= max && y < large-min {
			snail[(x*large)+y+1] = i
			right[i] = (x * large) + y + 1
			y++
			i++
		}
		y--
		for x+1 < long-min && i <= max {
			x++
			snail[(x*large)+y+1] = i
			right[i] = (x * large) + y + 1
			i++
		}
		for y-1 >= min && i <= max {
			y--
			snail[(x*large)+y+1] = i
			right[i] = (x * large) + y + 1
			i++
		}
		min++
		for x-1 >= min && i <= max {
			x--
			snail[(x*large)+y+1] = i
			right[i] = (x * large) + y + 1
			i++
		}
		y++
	}
	if i == max && max%2 == 1 && long == large {
		snail[(x*large)+y+1] = i
		right[i] = (x * large) + y + 1
	}
	i--
	obv = right[i]
}

//	*	ConvertBoard from snail to right
func ConvertToRight(board [][]int, long, large int) {
	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			board[x][y] = right[board[x][y]]
		}
	}
}

//	*	ConvertBoard from right to snail
func ConvertToSnail(board [][]int, long, large int) {
	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			board[x][y] = snail[board[x][y]]
		}
	}
}

//	*	Transform [][]int to []int
func BoardToString(board [][]int, long, large int) []int {
	var ret = make([]int, long*large)

	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			ret[(x*large)+y] = board[x][y]
		}
	}
	return ret
}

//	*	Set Waited board
func SetObjectifBoard(long, large int) [][]int {
	x, y, i := 0, 0, 1
	var objtf = make([][]int, long)
	for x < long {
		objtf[x] = make([]int, large)
		y = 0
		for y < large {
			objtf[x][y] = i
			i++
			y++
		}
		x++
	}
	return (objtf)
}

//	*	Check solvability by even/even inversion for initial state and final state
func CheckSolvability(board [][]int, long, large int) [][]int {
	var final [][]int
	var even int

	final = SetObjectifBoard(long, large)
	even = BoolToInt(large%2 == 0) * large * long
	ConvertToSnail(final, long, large)
	/*ConvertToSnail(board, long, large)
	fmt.Println("Size =", even, "\nfinal = {", BoardToString(final, long, large), "}\nboard = {", BoardToString(board, long, large), "}")
	fmt.Println("TODOCheck inversion 1 :", CheckInversion(BoardToString(board, long, large), even, large), "check inversion 2 :", CheckInversion(BoardToString(final, long, large), even, large))*/
	if CheckInversion(BoardToString(board, long, large), even, large) == CheckInversion(BoardToString(final, long, large), even, large) {
		ConvertToRight(final, long, large)
		return final
	}
	/* else {
		//TODO Check (Inversion final) != (Inversion first) even if it's solvable
		fmt.Println("TODOCheck inversion 1 :", CheckInversion(BoardToString(board, long, large), even, large), "check inversion 2 :", CheckInversion(BoardToString(final, long, large), even, large))
		ConvertToRight(board, long, large)
		return final
	}*/
	return nil
}

//	*	Check number of inversion
func CheckInversion(check []int, even int, large int) bool {
	var i, max, r, j, iseven int

	i, r, iseven = 0, 0, 0
	max = len(check)
	for i < max {
		if check[i] == even && iseven == 0 && even > 0 {
			iseven = i / large
			/*fmt.Println("iseven =", iseven)*/
		}
		if check[i] != large*large {
			for j = 1; j+i < max; j++ {
				if check[i] > check[i+j] {
					r++
				}
			}
		}
		i++
	}
	return (((r + iseven) % 2) == 0)
}

// Returns true if `start` belongs to the same permutation group as `goal`
/*template <uint size>
bool isSolvable(const Puzzle<size> & start, const Puzzle<size> & goal) {
	auto startInversions = inversions(start);
	auto goalInversions = inversions(goal);

	if (size % 2 == 0) { // In this case, the row of the '0' tile matters
		startInversions += start.indexOf(0) / size;
		goalInversions += goal.indexOf(0) / size;
	}

	return (startInversions % 2 == goalInversions % 2);
}*/
