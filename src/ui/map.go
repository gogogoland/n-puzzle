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
func SeparateImgInMap(Pwidth, Pheight int){
	PSurface := Pheight * Pwidth
	Frame = make(map[int]Puzzle)
	Win = make(map[int]Window)
	x, y, i, min := 1, 1, 1, 0

	for i < PSurface{
		for x <= (Pwidth - min) && i <= PSurface {
			Frame[i] = Puzzle{
				((x - 1) * imgWidth / Pwidth),
				((y - 1) * imgHeight / Pheight),
			}
			Win[i] = Window{
				((x - 1) * winWidth / Pwidth),
				((y - 1) * winHeight / Pheight),
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
				((x - 1) * imgWidth / Pwidth),
				((y - 1) * imgHeight / Pheight),
			}
			Win[i] = Window{
				((x - 1) * winWidth / Pwidth),
				((y - 1) * winHeight / Pheight),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y, ", x, y)
			i++
		}
		for x - 1 > min && i <= PSurface {
			x--
			Frame[i] = Puzzle{
				((x - 1) * imgWidth / Pwidth),
				((y - 1) * imgHeight / Pheight),
			}
			Win[i] = Window{
				((x - 1) * winWidth / Pwidth),
				((y - 1) * winHeight / Pheight),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y ", x, y)
			i++
		}
		min++
		for y - 1 > min && i <= PSurface {
			y--
			Frame[i] = Puzzle{
				((x - 1) * imgWidth / Pwidth),
				((y - 1) * imgHeight / Pheight),
			}
			Win[i] = Window{
				((x - 1) * winWidth / Pwidth),
				((y - 1) * winHeight / Pheight),
			}
			//Debug
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y ", x, y)
			i++
		}
		x++
	}
	if PSurface % 2 == 1 && Pwidth == Pheight{
			Frame[i] = Puzzle{
				((x - 1) * imgWidth / Pwidth),
				((y - 1) * imgHeight / Pheight),
			}
			Win[i] = Window{
				((x - 1) * winWidth / Pwidth),
				((y - 1) * winHeight / Pheight),
			}
			fmt.Println(i, " => Window :", Win[i], "Frame :", Frame[i], "X,Y ", x, y)
		}
}
