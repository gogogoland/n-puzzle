/* ROY CHARLES
** croy@student.42.fr
** France
** N-Puzzle project
*/

package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

func DrawPuzzle(array []int){
	running = true

	for running {
		event = sdl.PollEvent()

		switch event.(type) {
		case *sdl.QuitEvent:
			running = false
		}

		renderer.Clear()

		for i := 1; i <= PSurface; i++ {
			src = sdl.Rect{int32(Frame[array[i - 1]].Width), int32(Frame[array[i - 1]].Height),
				int32(imgWidth/Pwidth), int32(imgHeight/Pheight)}
			dst = sdl.Rect{int32(Win[i].Width + 1), int32(Win[i].Height + 1),
				 int32(winWidth/Pwidth - 1), int32(winHeight/Pheight - 1)}
			renderer.Copy(texture, &src, &dst)
		}

		rect = sdl.Rect{0, 0, int32(winWidth), int32(winHeight)}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.DrawRect(&rect)

		renderer.Present()
	}
}
