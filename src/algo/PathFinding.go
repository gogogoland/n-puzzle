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
	"container/list"
	"fmt"
)

//	TEST HEAP

// An Tabl is a min-heap of ints.

func (h PrioQueue) Len() int           { return len(h) }
func (h PrioQueue) Less(i, j int) bool { return h[i].rang < h[j].rang }
func (h PrioQueue) Swap(i, j int)      { h[i].rang, h[j].rang = h[j].rang, h[i].rang }

func (h *PrioQueue) Push(x interface{}) {
	*h = append(*h, Tabl{rang: x.(Tabl).rang, from: x.(Tabl).from})
}

//	Get last elem and delete it
func (h *PrioQueue) Pop() interface{} {
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
	h := &PrioQueue{Tabl{rang: 2, from: 2}}
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
func Pathfinding(board, objtf [][]int, long, large, algo int) *list.List {
	open := &PrioQueue{Tabl{
		rang:  0,
		from:  0,
		table: board,
		cur:   0,
		g:     0,
		h:     0,
		x:     0,
		y:     0}}
	(*open)[0].x, (*open)[0].y = MissPuzzle(board, long, large)
	var tmp [4]Tabl
	end := false
	id := 0

	//	Init Heap
	heap.Init(open)
	close := &PrioQueue{Tabl{
		rang:  0,
		from:  0,
		table: objtf,
		cur:   0,
		g:     0,
		h:     0,
		x:     0,
		y:     0}}
	(*close)[0].x, (*close)[0].y = MissPuzzle(objtf, long, large)
	heap.Init(close)
	end = false
	for len(*open) > 0 && !end {
		//	Get Highest priority of open
		cur := heap.Pop(open)
		//	Push current in close list (or Init close list with)
		heap.Push(close, cur)
		//	Find next path
		tmp, id = AlgoAStar(cur.(Tabl), long, large, id, algo)
		for i := 0; i < 4; i++ {
			if tmp[i].rang > -1 {
				//	Check if path exist already if so, check the fewest rang for the open list
				if !ComparePrioQueue(tmp[i], *close, long, large) && !ComparePrioQueue(tmp[i], *open, long, large) {
					heap.Push(open, tmp[i])
				}
				//	Check if it's final
				if CompareTable(tmp[i], (*close)[0], long, large) {
					end = true
				}
			}
		}
	}
	if end {
		path := list.New()
		i := 0
		cur := heap.Pop(close).(Tabl)
		path.PushFront(Path{x: cur.x, y: cur.y})
		from := cur.from
		for from != 0 {
			//	add to list final
			if (*close)[i].cur == from {
				path.PushFront(Path{x: (*close)[i].x, y: (*close)[i].y})
				i := 0
				from = (*close)[i].from
			}
		}
		return path
	}
	return nil
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
	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			if b1.table[x][y] != b2.table[x][y] {
				return false
			}
		}
	}
	return true
}

//	*	Find missing piece
func MissPuzzle(board [][]int, long, large int) (int, int) {
	x := 0
	y := 0

	for board[x][y] != long*large {
		y := 0
		for (board[x][y] != long*large) && (y < large) {
			y++
		}
		x++
	}
	return x, y
}

//	*	Algorithme
func AlgoAStar(cur Tabl, long, large, id, algo int) ([4]Tabl, int) {
	var path [4]Tabl
	mx := 0
	my := 0
	i := 0

	mx, my = MissPuzzle(cur.table, long, large)
	for x := -1; x < 2; x += 2 {
		for y := -1; y < 2; y += 2 {
			if (mx+x) < long && (mx+x) >= 0 && (my+y) >= 0 && (my+y) < large {
				cur.table[mx+x][my+y], cur.table[mx][my] = cur.table[mx][my], cur.table[mx+x][my+y]
				switch algo {
				case 0:
					path[i] = Marecages(cur, long, large)
					//case 1:
					//	path[i] = Euclidien(cur, long, large)
					//case 2:
					//	path[i] = Manahttan(cur, long, large)
				}
				id++
				path[i].cur, path[i].from = id, cur.cur
				path[i].g = cur.g + 1
				path[i].rang = path[i].g + path[i].h
				path[i].x, path[i].y = mx+x, my+y
				cur.table[mx+x][my+y], cur.table[mx][my] = cur.table[mx][my], cur.table[mx+x][my+y]
			}
			i++
		}
	}
	return path, id
}

//	*	*	Manahttan

//	*	*	Euclidien

//	*	*	Marecages
func Marecages(cur Tabl, long, large int) Tabl {
	res := Tabl{
		rang:  cur.rang,
		from:  0,
		table: cur.table,
		cur:   0,
		g:     0,
		h:     0,
		x:     0,
		y:     0}

	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			if cur.table[x][y] == 0 {
				tmpH := ((long / 2) + (long*large)/2) - ((x + 1) + (y * large))
			} else {
				tmpH := cur.table[x][y] - ((x + 1) + (y * large))
			}
			if tmpH < 0 {
				tmpH = -1 * tmpH
			}
			res.h += tmpH
		}
	}
	return res
}

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
