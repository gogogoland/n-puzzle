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
func PrintAll(pr Tabl, long, large int) {
	fmt.Println("Move number", pr.g)
	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			fmt.Printf("%d ", pr.table[x][y])
		}
		fmt.Printf("\n")
	}
	fmt.Println("`~I~'")
}

/*
 * TEST
 */

//	Functions A*
//	*	Check if the board is solvable
func CheckSolvable(board [][]int, long, large, algo int) (bool, int) {
	var cross, i, j, max int
	var chain []int

	cross = 0
	max = long * large
	chain = BoardToString(board, long, large)
	for i = 0; i < max; i++ {
		for j = i + 1; j < max; j++ {
			if chain[j] != obv && chain[i] != obv && chain[j] < chain[i] {
				cross++
			}
		}
	}
	if cross != 0 && long%2 == cross%2 {
		fmt.Println("No solution")
		return false, 0
	}
	fmt.Println("Is soluble. .. I hope.")
	return true, cross
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

//	*	Implementation of A*
func Pathfinding(board [][]int, long, large, algo int) *list.List {
	if board == nil || long <= 0 || large <= 0 {
		return nil
	} else if algo < 0 || algo > 2 {
		fmt.Println("Algo value should be between 0 and 2 both inclued.")
		return nil
	}
	SaveSnail(board, long, large)
	ConvertToRight(board, long, large)
	solv, end := CheckSolvable(board, long, large, algo)
	if !solv {
		return nil
	}
	objtf := SetObjectifBoard(long, large)
	open := InitHeapList(board, long, large)
	close := InitHeapList(objtf, long, large)
	if end == 0 {
		return ListPath(long, large, (*close)[0], nil)
	}
	var tmp [4]Tabl
	id := 0
	//	Init Heap
	heap.Init(open)
	heap.Init(close)
	for len(*open) > 0 {
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
					return ListPath(long, large, tmp[i], *close)
				}
			}
		}
	}
	fmt.Println("No solution")
	return nil
}

func ListPath(long, large int, tmp Tabl, close []Tabl) *list.List {
	var i, from int

	//	if there a solution, set a list of path
	path := list.New()
	ConvertToSnail(tmp.table, long, large)
	path.PushFront(Path{Ret: BoardToString(tmp.table, long, large)})
	from = tmp.from
	//	TEST BEG
	defer fmt.Println("Number of case tested: ", len(close))
	defer PrintAll(tmp, long, large)
	//	TEST END
	i = 0
	for from != 0 {
		//	add to list final
		if close[i].cur == from {
			ConvertToSnail(close[i].table, long, large)
			path.PushFront(Path{Ret: BoardToString(close[i].table, long, large)})
			from = close[i].from
			//	TEST BEG
			defer PrintAll(close[i], long, large)
			//	TEST END
			i = 0
		}
		i++
	}
	return path
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

	mx, my = cur.x, cur.y
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
			path[i].rang = cur.h + path[i].h
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
	var tmpHx, tmpHy, x, y int

	for x = 0; x < long; x++ {
		for y = 0; y < large; y++ {
			tmpHx = ((res.table[x][y] - 1) / large) - x
			if tmpHx < 0 {
				tmpHx = -1 * tmpHx
			}
			tmpHy = ((res.table[x][y] - 1) % large) - y
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
	var tmpHx, tmpHy, tmpH, x, y int

	for x = 0; x < long; x++ {
		for y = 0; y < large; y++ {
			tmpHx = ((res.table[x][y] - 1) / large) - x
			if tmpHx < 0 {
				tmpHx = -1 * tmpHx
			}
			tmpHy = ((res.table[x][y] - 1) % large) - y
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
	var tmpH, tmpHx, tmpHy, tmpG, tmpGx, tmpGy, gx, gy int

	gx, gy = MissPuzzle(cur, long, large)
	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			tmpHx = ((res.table[x][y] - 1) / large) - x
			if tmpHx < 0 {
				tmpHx = -1 * tmpHx
			}
			tmpHy = ((res.table[x][y] - 1) % large) - y
			if tmpHy < 0 {
				tmpHy = -1 * tmpHy
			}
			tmpH = tmpHx + tmpHy
			tmpGx = (gx - x) * (gx - x)
			tmpGy = (gy - y) * (gy - y)
			tmpG = int(math.Sqrt(float64(tmpGx + tmpGy)))

			//tmpH = res.table[x][y] - ((x + 1) + (y * large))
			res.h += tmpG * tmpH
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
