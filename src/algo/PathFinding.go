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
	"math"
)

//	Functions for heap list
func (h PrioQueue) Len() int {
	return len(h)
}

func (h PrioQueue) Less(i, j int) bool {
	return h[i].rang < h[j].rang
}

func (h PrioQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

//	*	Add Tabl structure in heap list
func (h *PrioQueue) Push(x interface{}) {
	*h = append(*h, Tabl{
		rang:  x.(Tabl).rang,
		from:  x.(Tabl).from,
		table: x.(Tabl).table,
		cur:   x.(Tabl).cur,
		g:     x.(Tabl).g,
		h:     x.(Tabl).h,
		x:     x.(Tabl).x,
		y:     x.(Tabl).y,
	})
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

func PrintAll(board [][]int, long, large int) {
	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			fmt.Printf("%d ", board[x][y])
		}
		fmt.Printf("\n")
	}
	fmt.Println("`~I~'")
}

/*
 * TEST
 */

//	Functions A*
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

//	*	Implementation of A*
func Pathfinding(board [][]int, long, large, algo int) *list.List {
	if board == nil || long <= 0 || large <= 0 || algo < 0 || algo > 2 {
		return nil
	}
	SaveSnail(board, long, large)
	ConvertToRight(board, long, large)
	objtf := SetObjectifBoard(long, large)
	open := InitHeapList(board, long, large)
	close := InitHeapList(objtf, long, large)
	var tmp [4]Tabl
	var from int
	end := false
	id := 0
	//	Init Heap
	heap.Init(open)
	heap.Init(close)
	for len(*open) > 0 && !end {
		//	Get Highest priority of open
		cur := heap.Pop(open)
		//	Push current in close list (or Init close list with)
		heap.Push(close, cur)
		heap.Fix(close, len(*close)-1)
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
					from = tmp[i].from
					id = i
					end = true
				}
			}
		}
	}
	//	if there a solution, set a list of path
	if end {
		path := list.New()
		ConvertToSnail(tmp[id].table, long, large)
		path.PushFront(Path{Ret: Return(tmp[id].table, long, large)})
		//	TEST BEG
		PrintAll(tmp[id].table, long, large)
		//	TEST END
		i := 0
		for from != 0 {
			//	add to list final
			if (*close)[i].cur == from {
				ConvertToSnail((*close)[i].table, long, large)
				path.PushFront(Path{Ret: Return((*close)[i].table, long, large)})
				from = (*close)[i].from
				//	TEST BEG
				PrintAll((*close)[i].table, long, large)
				//	TEST END
				i = 0
			}
			i++
		}
		//	else return null
		return path
	}
	fmt.Println("No solution")
	//	else return null
	return nil
}

//	*	Find missing piece
func MissPuzzle(board [][]int, long, large int) (int, int) {
	var x int
	var y int

	x = 0
	y = 0
	for board[x][y] != obv {
		for y < large && (board[x][y] != obv) {
			y++
		}
		if y == large {
			y = 0
			x++
		}
	}
	return x, y
}

//	*	Algorithme
func AlgoAStar(cur Tabl, long, large, id, algo int) ([4]Tabl, int) {
	var path [4]Tabl
	mx := 0
	my := 0

	mx, my = MissPuzzle(cur.table, long, large)
	for i := 0; i < 4; i++ {
		x := (i - 2) % 2
		y := (i - 1) % 2
		if (mx+x) < long && (mx+x) >= 0 && (my+y) >= 0 && (my+y) < large {
			//	Switch place
			cur.table[mx+x][my+y], cur.table[mx][my] = cur.table[mx][my], cur.table[mx+x][my+y]
			//	Execute algo
			switch algo {
			case 0:
				path[i] = Marecages(cur.table, long, large, mx+x, my+y)
			case 1:
				path[i] = Euclidien(cur.table, long, large, mx+x, my+y)
			case 2:
				path[i] = Manahttan(cur.table, long, large, mx+x, my+y)
			}
			id++
			//	Save for list path
			path[i].cur, path[i].from = id, cur.cur
			//	Calcul cost of path
			path[i].g = cur.g + 1
			path[i].rang = path[i].g + path[i].h
			//	Save position of obv
			path[i].x, path[i].y = mx+x, my+y
			//	Reswitch place
			cur.table[mx+x][my+y], cur.table[mx][my] = cur.table[mx][my], cur.table[mx+x][my+y]
		} else {
			path[i].rang = -1
		}
	}
	return path, id
}

//	*	*	Manahttan
func Manahttan(cur [][]int, long, large, mx, my int) Tabl {
	res := InitTable(cur, mx, my)
	tmpHx := 0
	tmpHy := 0

	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			tmpHx = (res.table[x][y] % long) - (x + 1)
			if tmpHx < 0 {
				tmpHx = -1 * tmpHx
			}
			tmpHy = (res.table[x][y] / long) - y
			if tmpHy < 0 {
				tmpHy = -1 * tmpHy
			}
			res.h += tmpHy + tmpHx
		}
	}
	return res
}

//	*	*	Euclidien
func Euclidien(cur [][]int, long, large, mx, my int) Tabl {
	res := InitTable(cur, mx, my)
	tmpHx := 0
	tmpHy := 0
	tmpH := 0

	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			tmpHx = (res.table[x][y] % long) - (x + 1)
			if tmpHx < 0 {
				tmpHx = -1 * tmpHx
			}
			tmpHy = (res.table[x][y] / long) - y
			if tmpHy < 0 {
				tmpHy = -1 * tmpHy
			}
			tmpH = (tmpHx + tmpHy) * (tmpHy + tmpHx)
			res.h += int(math.Sqrt(float64(tmpH)))
		}
	}
	return res
}

//	*	*	Marecages
func Marecages(cur [][]int, long, large, mx, my int) Tabl {
	res := InitTable(cur, mx, my)
	tmpH := 0

	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			tmpH = res.table[x][y] - ((x + 1) + (y * large))
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
