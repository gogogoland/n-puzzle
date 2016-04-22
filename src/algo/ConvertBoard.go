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

//	*	Return Value Final
func Return(board [][]int, long, large int) []int {
	var ret = make([]int, long*large)

	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			ret[(x*large)+y] = board[x][y]
		}
	}
	return ret
}
