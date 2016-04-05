/* ROY CHARLES
** croy@student.42.fr
** France
** N-Puzzle project
*/

package ui

import (
	"fmt"
)

//Get Size of each frame of puzzle
func SeparateImgInMap(){
	Frame = make(map[int]Puzzle)
	Win = make(map[int]Window)
	x, y, i, min := 1, 1, 1, 0

	for i < PSurface{
		for x <= (Pwidth - min) && i <= PSurface {
			Frame[i] = Puzzle{
				(x * imgWidth / Pwidth),
				(y * imgHeight / Pheight),
			}
			Win[i] = Window{
				(x * winWidth / Pwidth),
				(y * winHeight / Pheight),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y, ", x, y)
			x++
			i++
		}
		x--
		for y + 1 <= (Pheight - min) && i <= PSurface {
			y++
			Frame[i] = Puzzle{
				(x * imgWidth / Pwidth),
				(y * imgHeight / Pheight),
			}
			Win[i] = Window{
				(x * winWidth / Pwidth),
				(y * winHeight / Pheight),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y, ", x, y)
			i++
		}
		for x - 1 > min && i <= PSurface {
			x--
			Frame[i] = Puzzle{
				(x * imgWidth / Pwidth),
				(y * imgHeight / Pheight),
			}
			Win[i] = Window{
				(x * winWidth / Pwidth),
				(y * winHeight / Pheight),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y ", x, y)
			i++
		}
		min++
		for y - 1 > min && i <= PSurface {
			y--
			Frame[i] = Puzzle{
				(x * imgWidth / Pwidth),
				(y * imgHeight / Pheight),
			}
			Win[i] = Window{
				(x * winWidth / Pwidth),
				(y * winHeight / Pheight),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y ", x, y)
			i++
		}
		x++
	}
}
