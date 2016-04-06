/* ROY CHARLES
** croy@student.42.fr
** France
** N-Puzzle project
*/

package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

func DrawPuzzle(){
	renderer.Clear()
	rect := sdl.Rect{0, 0, int32(winWidth), int32(winHeight)}
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.DrawRect(&rect)
	for i := 1; i <= PSurface; i++ {
		src = sdl.Rect{int32(Frame[i].Width), int32(Frame[i].Height), int32(imgWidth/Pwidth), int32(imgHeight/Pheight)}
		dst = sdl.Rect{int32(Win[i].Width), int32(Win[i].Height), int32(winWidth/Pwidth - 1), int32(winHeight/Pheight - 1)}
		renderer.Copy(texture, &src, &dst)
	}
	renderer.Present()
}
