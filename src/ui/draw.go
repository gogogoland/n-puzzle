/* ROY CHARLES
** croy@student.42.fr
** France
** N-Puzzle project
 */

package ui

import (
	"algo"
	"container/list"
	//"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

//Take an [][]path of int and draw while it's not the end.
func DrawPuzzle(lst *list.List, Pwidth, Pheight int) {
	running := true
	solve := 1
	PSurface := Pheight * Pwidth

	time.Sleep(500 * time.Millisecond)
	for running && lst != nil && lst.Front() != nil {
		renderer.Clear()
		for path := lst.Front(); path != nil && solve == 1; path = path.Next() {
			rect = sdl.Rect{0, 0, int32(winWidth), int32(winHeight)}
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.FillRect(&rect)
			for i := 1; i <= PSurface; i++ {
				switch {
				case (path.Value.(algo.Path).Ret[i-1] != PSurface):
					src = sdl.Rect{
						int32(Frame[path.Value.(algo.Path).Ret[i-1]].Width),
						int32(Frame[path.Value.(algo.Path).Ret[i-1]].Height),
						int32(imgWidth / Pwidth),
						int32(imgHeight / Pheight)}
					dst = sdl.Rect{
						int32(Win[i].Width + 1),
						int32(Win[i].Height + 1),
						int32(winWidth/Pwidth - 1),
						int32(winHeight/Pheight - 1)}

				case path.Next() == nil:
					src = sdl.Rect{
						int32(Frame[path.Value.(algo.Path).Ret[i-1]].Width),
						int32(Frame[path.Value.(algo.Path).Ret[i-1]].Height),
						int32(imgWidth / Pwidth),
						int32(imgHeight / Pheight)}
					dst = sdl.Rect{
						int32(Win[i].Width + 1),
						int32(Win[i].Height + 1),
						int32(winWidth/Pwidth - 1),
						int32(winHeight/Pheight - 1)}
				}
				renderer.Copy(texture, &src, &dst)
			}
			renderer.Present()
			time.Sleep(50 * time.Millisecond)
			renderer.Clear()
		}
		solve += 1
		sdl.Delay(1500)
		running = false
	}
}
