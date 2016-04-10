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
	snail map[int]int
	right map[int]int
)

//	Functions
//	*	Mapping snails value Amd save missing value
//	*	*	TO TEST
func SaveSnail(board [][]int, long, large int) {
	x, y := 0, 0
	i, min, max := 1, 0, long*large
	snail := map[int]int{}
	right := map[int]int{}

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
		x++
	}
	if max%2 == 1 && long == large {
		snail[(x*large)+y+1] = i
		right[i] = (x * large) + y + 1
		i++
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
