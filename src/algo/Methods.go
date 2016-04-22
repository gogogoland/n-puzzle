/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   Methods.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbalea <tbalea@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/03/22 18:42:35 by tbalea            #+#    #+#             */
/*   Updated: 2016/03/22 18:42:35 by tbalea           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package algo

/*import (
"container/heap"
)*/

//	Methods
//	*	linked list for mouvement
type Path struct {
	X, Y  int
	Board [][]int
	Ret   []int
}

//	*	Puzzle composition and value of it
type Tabl struct {
	table [][]int
	rang  int
	g, h  int
	from  int
	cur   int
	x, y  int
}

//	*	Slice of Tabl
type PrioQueue []Tabl

//	*	Initialise Table
func InitTable(board [][]int, x, y int) Tabl {
	tmp := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		tmp[i] = make([]int, len(board[i]))
		copy(tmp[i], board[i])
	}

	table := Tabl{
		rang:  0,
		from:  0,
		table: tmp,
		cur:   0,
		g:     0,
		h:     0,
		x:     x,
		y:     y}
	return table
}

//	*	Initialise Heap List
func InitHeapList(board [][]int, long, large int) *PrioQueue {
	lx, ly := MissPuzzle(board, long, large)
	queue := &PrioQueue{InitTable(board, lx, ly)}
	return queue
}

//	*	ComparePrioQueue
func ComparePrioQueue(tbl Tabl, lst PrioQueue, long, large int) bool {
	max := len(lst)
	for i := 0; i < max; i++ {
		if tbl.h == lst[i].h {
			if CompareTable(tbl, lst[i], long, large) {
				//	if equals, get the fewest rang
				//	?	Maybe fix heap
				//	?	Only to open list
				if tbl.g < lst[i].g {
					lst[i].rang = tbl.rang
					lst[i].from = tbl.from
					lst[i].g = lst[i].g
				}
				return true
			}
		}
	}
	return false
}

//	*	Compare table
func CompareTable(b1, b2 Tabl, long, large int) bool {
	if &b1 == nil || &b2 == nil {
		return false
	}
	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			if b1.table[x][y] != b2.table[x][y] {
				return false
			}
		}
	}
	return true
}

//  *   Functions for method Path
//  *   *   Add Path
//func (cur *Path) AddPath(add *Path) {
//	cur.next = *add
//	add.prev = *cur
//}

//  *   *   Initialaze Path
//func (p Path) InitPath(x, y uint) {
//	p.next = nil
//	p.prev = nil
//	p.x = x
//	p.y = y
//}
