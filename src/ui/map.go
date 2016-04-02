package ui

import (
	"fmt"
)

func SeparateFrameInMap(){
	Frame = make(map[int]Puzzle)
	Win = make(map[int]Window)
	x, y, i, min := 1, 1, 1, 0

	for i < PSurface{
		for x < (Pwidth - min){
			Frame[i] = Puzzle{
				(x * imgWidth / 3),
				(y * imgHeight / 3),
			}
			Win[i] = Window{
				((x - 1) * winWidth / 3),
				((y - 1) * winHeight / 3),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y, Min : ", x, y, min)
			x++
			i++
		}
		for y < (Pheight - min){
			Frame[i] = Puzzle{
				(x * imgWidth / Pwidth),
				(y * imgHeight / Pheight),
			}
			Win[i] = Window{
				((x - 1) * winWidth / Pwidth),
				((y - 1) * winHeight / Pheight),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y, Min : ", x, y, min)
			y++
			i++
		}
		for x > min + 1{
			Frame[i] = Puzzle{
				(x * imgWidth / Pwidth),
				(y * imgHeight / Pheight),
			}
			Win[i] = Window{
				((x - 1) * winWidth / Pwidth),
				((y - 1) * winHeight / Pheight),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y, Min : ", x, y, min)
			x--
			i++
		}
		min++
		for y > min + 1{
			Frame[i] = Puzzle{
				(x * imgWidth / Pwidth),
				(y * imgHeight / Pheight),
			}
			Win[i] = Window{
				((x - 1) * winWidth / Pwidth),
				((y - 1) * winHeight / Pheight),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y, Min : ", x, y, min)
			y--
			i++
		}
	}
	//ADD the last Frame here.
	Frame[PSurface] = Puzzle{
		(x * imgWidth / Pwidth),
		(y * imgHeight / Pheight),
	}
	Win[PSurface] = Window{
		((x - 1) * winWidth / Pwidth),
		((y - 1) * winHeight / Pheight),
	}
	fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y, Min : ", x, y, min)
}
