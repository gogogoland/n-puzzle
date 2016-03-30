/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   PathFinding.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbalea <tbalea@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/03/17 19:24:03 by tbalea            #+#    #+#             */
/*   Updated: 2016/03/17 19:24:03 by tbalea           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package algo

import (
	"container/heap"
	"fmt"
)

//	TEST HEAP

// An Tabl is a min-heap of ints.

func (h List) Len() int           { return len(h) }
func (h List) Less(i, j int) bool { return h[i].rang < h[j].rang }
func (h List) Swap(i, j int)      { h[i].rang, h[j].rang = h[j].rang, h[i].rang }

func (h *List) Push(x interface{}) {
	*h = append(*h, Tabl{rang: x.(Tabl).rang, from: x.(Tabl).from})
}

func (h *List) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
 * TEST
 */
// This example inserts several ints into an Tabl, checks the minimum,
// and removes them in order of priority.
func TestHeap() {
	h := &List{Tabl{rang: 2, from: 2}}
	//tmp := Tabl{rang: []int{-1}}

	heap.Init(h)
	heap.Push(h, Tabl{rang: -1, from: 0})
	heap.Push(h, Tabl{rang: -7, from: 1000})
	heap.Push(h, Tabl{rang: 81, from: -10000})
	heap.Push(h, Tabl{rang: 0, from: 0})
	fmt.Printf("minimum: %d\n", (*h)[0].rang)
	//	TO TEST
	for h.Len() > 0 {
		fmt.Printf("rg:%d\n", heap.Pop(h).(Tabl).rang)
	}
}

/*
 * TEST
 */

//	Functions
//	*	Implementation of A*
func Pathfinding(board, objtf [][]int, long, large, algo int) *Path {
	open := &List{Tabl{rang: 0, from: 0, table: board}}
	var tmp [4]Tabl
	end := false

	//	Init Heap
	heap.Init(open)
	close := &List{Tabl{rang: 0, from: 0, table: objtf}}
	heap.Init(close)
	end = false
	for len(*open) > 0 && !end {
		//	Get Highest priority of open
		cur := heap.Pop(open)
		//	Push current in close list (or Init close list with)
		heap.Push(close, cur)
		//	Find next path
		//tmp := AlgoAStar(cur, long, large, algo)
		for i := 0; i < 4; i++ {
			if tmp[i].rang > -1 {
				//	Check if path exist already if so, check the fewest rang for the open list
				if !CompareList(tmp[i], *close, long, large) && !CompareList(tmp[i], *open, long, large) {
					heap.Push(open, tmp[i])
				}
				//	Check if it's final
				if CompareTable(tmp[i], (*close)[0], long, large) {
					end = true
				}
			}
		}
	}
	/*if end {
		for close {
			//	add to list final
		}
		return list
	}*/
	return nil
}

//	*	Compare List
func CompareList(tbl Tabl, lst List, long, large int) bool {
	max := len(lst)
	for i := 0; i < max; i++ {
		if CompareTable(tbl, lst[i], long, large) {
			//	if equals, get the fewest rang
			//	?	Maybe fix heap
			//	?	Only to open list
			if tbl.rang < lst[i].rang {
				tbl.rang = lst[i].rang
				tbl.from = lst[i].from
			}
			return true
		}
	}
	return false
}

//	*	Compare table
func CompareTable(b1, b2 Tabl, long, large int) bool {
	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			if b1.table[x][y] != b2.table[x][y] {
				return false
			}
		}
	}
	return true
}

//	*	Algorithme
func AlgoAStar(cur Tabl, long, large, algo int) [4]Tabl {
	var path [4]Tabl
	mx := 0
	my := 0
	i := 0

	for cur.table[mx][my] != long*large {
		for cur.table[mx][my] != long*large {
			my++
		}
		mx++
	}
	for x := -1; x < 2; x += 2 {
		for y := -1; y < 2; y += 2 {
			if (mx+x) < long && (mx+x) >= 0 && (my+y) >= 0 && (my+y) < large {
				cur.table[mx+x][my+y], cur.table[mx][my] = cur.table[mx][my], cur.table[mx+x][my+y]
				//switch algo {
				//case 0:
				//	path[i] = Manahttan(cur)
				//case 1:
				//	path[i] = ALGO2(cur)
				//case 2:
				//	path[i] = ALGO3(cur)
				cur.table[mx+x][my+y], cur.table[mx][my] = cur.table[mx][my], cur.table[mx+x][my+y]
			}
			i++
		}
	}
	return path
}

//	*	*	Manahttan
//func Manahttan ()

//	*	Set Table of Obstacle
//			? Only one time ?
/*func CalculObstacle(board [][]uint, long uint, larg uint) [][]uint {
	obstacle := make([][]uint, long, long)
	var tmp uint

	for x := 0; x < long; x++ {
		obstacle[x] = make([]uint, larg, larg)
		for y := 0; y < larg; y++ {
			tmp = board[x][y] - 1
			obstacle[x][y] = (tmp%larg + tmp/long) - (y + x*larg)
		}
	}
	return obstacle
}*/

//	*	Calcul Path
/*func	CalculPath(obstacle [][]uint, xBeg, yBeg, xObj, yObj) *Path {
	pathfind := make([][]int, long, long)
	var path

	for x := 0; x < long; x++ {
		obstacle[x] = make([]uint, larg, larg)
		for y := 0; y < larg; y++ {
			pathfind[x][y] = obstacle[x][y];
		}
	}
	pathfind[xBeg][yBeg] = 0;
	path := InitPath(xBeg, yBeg);
	for x, y := 0 path != nil {
		;
	}
	for path.x != xBeg && path.y != yBeg {
		;
	}
	return path
}*/
