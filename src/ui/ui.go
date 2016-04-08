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
	event			sdl.Event
	running 	bool
	src, dst 	sdl.Rect
	rect 			sdl.Rect
	err 		error
)

var (
	winWidth, winHeight int = 1024, 764
	winTitle string = "N-Puzzle"
	imgWidth, imgHeight int = 1500, 1000
	imgPuzzle string = "../assets/cat.bmp"
)

var (
	Frame map[int]Puzzle
	Win map[int]Window
	Pwidth, Pheight int = 5, 8
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

	//Init Img for Puzzle
	image, err = sdl.LoadBMP(imgPuzzle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR. Failed to load BMP: %s\n", err)
		return 3
	}
	//Destroy img for Puzzle
	defer image.Free()


	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR. Failed to create texture: %s\n", err)
		return 4
	}
	defer texture.Destroy()

	//algo to get size of each piece of the puzzle
	SeparateImgInMap()
	
	//Draw Puzzle and event
	DrawPuzzle()

	return 0
}
