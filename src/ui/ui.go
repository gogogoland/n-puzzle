/* ROY CHARLES
** croy@student.42.fr
** France
** N-Puzzle project
*/

package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)


type Puzzle struct {
	Width, Height int
}

type Window struct {
	Width, Height int
}

var (
	window 		*sdl.Window
	renderer 	*sdl.Renderer
	texture 	*sdl.Texture
	image 		*sdl.Surface
	src, dst 	sdl.Rect
	err 		error
)

var (
	winWidth, winHeight int = 1024, 764
	winTitle string = "N-Puzzle"
	imgWidth, imgHeight int = 1495, 1026
	imgPuzzle string = "../assets/cat.bmp"
)

var (
	Frame map[int]Puzzle
	Win map[int]Window
	Pwidth, Pheight int = 14, 16
	PSurface = Pheight * Pwidth
)

//Main Function

func Ui() int{

	//Init Window
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR. Failed to create window: %s\n", err)
		return 1
	}
	//Destroy Window
	defer window.Destroy()


	//Init Render
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR. Failed to create renderer: %s\n", err)
		return 2
	}
	//Destroy Render
	defer renderer.Destroy()

	image, err = sdl.LoadBMP(imgPuzzle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR. Failed to load BMP: %s\n", err)
		return 3
	}
	defer image.Free()

	//Init Img for Puzzle
	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR. Failed to create texture: %s\n", err)
		return 4
	}
	//Destroy img for Puzzle
	defer texture.Destroy()

	SeparateImgInMap()
	//Draw Puzzle
	DrawPuzzle()

	sdl.Delay(4000)

	return 0
}

func DrawPuzzle(){
	renderer.Clear()
	src = sdl.Rect{1495 - 498, 342 - 342, int32(1495), int32(342)}
	dst = sdl.Rect{341, 254, int32(341), int32(254)}
	renderer.Copy(texture, &src, &dst)
	src = sdl.Rect{996 - 498, 342 - 342, int32(996), int32(342)}
	dst = sdl.Rect{0, 0, int32(341), int32(254)}

	renderer.Copy(texture, &src, &dst)
	renderer.Present()
}
