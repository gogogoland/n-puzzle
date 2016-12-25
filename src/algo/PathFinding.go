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
	"time"
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
	if board == nil || long <= 0 || large <= 0 {
		return nil
	} else if algo < 0 || algo > 7 {
		fmt.Println("Algo value should be between 0 and 7 both inclued.")
		return nil
	}
	SaveSnail(board, long, large)
	ConvertToRight(board, long, large)
	objtf := SetObjectifBoard(long, large)
	open := InitHeapList(board, long, large)
	close := InitHeapList(objtf, long, large)
	var tmp [4]Tabl
	id := 0
	//	Init Heap
	heap.Init(open)
	heap.Init(close)
	start := time.Now()
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
					elapsed := time.Since(start)
					defer fmt.Println("Time required: ", elapsed)
					return ListPath(long, large, tmp[i], *close)
				}
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Binomial took ", elapsed)
	fmt.Println("No solution founded")
	return nil
}

//	Fill list with found path
func ListPath(long, large int, tmp Tabl, close []Tabl) *list.List {
	var i, from int

	defer fmt.Println("Number of case tested: ", len(close))
	defer PrintAll(tmp, long, large)

	path := list.New()
	ConvertToSnail(tmp.table, long, large)
	path.PushFront(Path{Ret: BoardToString(tmp.table, long, large)})
	from = tmp.from
	i = 0
	for from != 0 {
		i++
		if close[i].cur == from {
			ConvertToSnail(close[i].table, long, large)
			path.PushFront(Path{Ret: BoardToString(close[i].table, long, large)})
			from = close[i].from
			i = 0

			defer PrintAll(close[i], long, large)
		}
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
				path[i] = Manahttan(cur.table, long, large, mx+x, my+y)
			case 1:
				path[i] = Euclidien(cur.table, long, large, mx+x, my+y)
			case 2:
				path[i] = Chebyshev(cur.table, long, large, mx+x, my+y)
			case 3:
				path[i] = Marecages(cur.table, long, large, mx+x, my+y)
			case 4:
				path[i] = FeuFollet(cur.table, long, large, mx+x, my+y, cur.objx, cur.objy)
			case 5:
				path[i] = Gollum(cur.table, long, large, mx+x, my+y, cur.objx, cur.objy)
			case 6:
				path[i] = IsWrong(cur.table, long, large, mx+x, my+y)
			case 7:
				path[i] = IsRight(cur.table, long, large, mx+x, my+y)
			}

			//	Save data
			id++
			path[i].cur, path[i].from = id, cur.cur
			path[i].g = cur.g + 1
			path[i].rang = cur.h + path[i].h
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
			//tmpH = (tmpHx + tmpHy) * (tmpHy + tmpHx)
			tmpH = (tmpHx * tmpHx) + (tmpHy + tmpHy)
			res.h += int(math.Sqrt(float64(tmpH)))
		}
	}
	return res
}

//	*	*	Marecages
func Marecages(cur [][]int, long, large, mx, my int) Tabl {
	res := InitTable(cur, mx, my)
	var tmpH, tmpHx, tmpHy, tmpG, tmpGx, tmpGy, gx, gy int

	tmpGx, tmpGy = 0, 0
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
			if tmpH > 0 {
				tmpGx = gx - x
				if tmpGx < 0 {
					tmpGx *= -1
				}
				tmpGy = gy - y
				if tmpGy < 0 {
					tmpGy *= -1
				}
				tmpG = tmpGx + tmpGy
			}
			res.h += tmpH + tmpG
		}
	}
	return res
}

//	*	*	Swamp Wisp
func FeuFollet(cur [][]int, long, large, mx, my, ox, oy int) Tabl {
	res := InitTable(cur, mx, my)
	var tmpH, tmpHx, tmpHy, tmpG, tmpGx, tmpGy, gx, gy int

	tmpG, tmpGx, tmpGy = 0, 0, 0
	gx, gy = MissPuzzle(cur, long, large)
	if gx == ox && gy == oy {
		ox, oy = -1, -1
	}
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
			if tmpH > tmpG && (oy != -1 || ox != -1) {
				tmpG = tmpH
				ox, oy = x, y
				tmpGx = gx - x
				if tmpGx < 0 {
					tmpGx *= -1
				}
				tmpGy = gy - y
				if tmpGy < 0 {
					tmpGy *= -1
				}
				/*if tmpG > tmpGx+tmpGy {
					tmpG = tmpGx + tmpGy
					ox, oy = x, y
				}*/
			}
			res.h += tmpH
		}
	}
	res.objx, res.objy = ox, oy
	tmpG = tmpGx + tmpGy
	res.h += tmpG
	return res
}

//	*	*	Smeagol
func Gollum(cur [][]int, long, large, mx, my, ox, oy int) Tabl {
	res := InitTable(cur, mx, my)
	var tmpH, tmpHx, tmpHy, tmpG, tmpM, gx, gy int

	tmpG, tmpM = 0, 0
	gx, gy = MissPuzzle(cur, long, large)
	if gx == ox && gy == oy {
		ox, oy = -1, -1
	}
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
			if tmpH > 0 && (oy != -1 || ox != -1) {
				if tmpH > tmpM {
					tmpM = tmpH
				}
				if tmpH < tmpG {
					tmpG = tmpH
					ox, oy = x, y
				}
			}
			res.h += tmpH
		}
	}
	res.objx, res.objy = ox, oy
	res.h += tmpM
	return res
}

//	*	*	Chebyshev
func Chebyshev(cur [][]int, long, large, mx, my int) Tabl {
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
			if tmpHx > tmpHy {
				res.h += tmpHx
			} else {
				res.h += tmpHy
			}
		}
	}
	return res
}

//	*	*	Number of incorrect case
func IsWrong(cur [][]int, long, large, mx, my int) Tabl {
	res := InitTable(cur, mx, my)

	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			if res.table[x][y] != cur[x][y] {
				res.h++
			}
		}
	}
	return res
}

//	*	*	Number of correct case
func IsRight(cur [][]int, long, large, mx, my int) Tabl {
	res := InitTable(cur, mx, my)

	for x := 0; x < long; x++ {
		for y := 0; y < large; y++ {
			if res.table[x][y] == cur[x][y] {
				res.h--
			}
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

//	*	Check if the board is solvable
/*func CheckSolvable(board [][]int, long, large, algo int) (bool, int) {
	var cross, i, j, max, d_pos, d_crs, d_lrg int
	var chain []int

	d_pos, cross = MissPuzzle(board, long, large)
	cross = 0
	max = long * large
	chain = BoardToString(board, long, large)
	for i = 0; i < max; i++ {
		if chain[i] == obv {
			d_pos = d_pos - (i / large)
		}
		for j = i + 1; j < max; j++ {
			if chain[j] < chain[i] {
				cross++
			}
		}
	}
	if cross == 0 {
		fmt.Println("Already finished. Come on, Test me !")
		return true, 0
	}
	d_lrg = large % 2
	d_crs = cross % 2
	d_pos = d_pos % 2 * d_pos % 2
	fmt.Println("d_l=", d_lrg, "d_c=", d_crs, "d_p=", d_pos, "cross=", cross)
	if (d_lrg == 1 && d_crs == 0) || (d_lrg == 0 && d_pos == d_crs) {
		fmt.Println("No solution")
		return false, cross
	}
	fmt.Println("Is soluble. .. I hope.")
	return true, cross
}*/
