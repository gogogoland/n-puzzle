package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

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
	imgPuzzle string = "../assets/cat.bmp"
)


//Main Function

func Ui() int{

	//Init Window
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	//Destroy Window
	defer window.Destroy()


	//Init Render
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	//Destroy Render
	defer renderer.Destroy()

	image, err = sdl.LoadBMP(imgPuzzle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s\n", err)
		return 3
	}
	defer image.Free()

	//Init Img for Puzzle
	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return 4
	}
	//Destroy img for Puzzle
	defer texture.Destroy()

	//Draw Puzzle
	Draw()

	sdl.Delay(4000)

	return 0
}

func Draw(){
	src = sdl.Rect{0, 0, 1495, 1026}
	dst = sdl.Rect{0, 0, int32(winWidth), int32(winHeight)}

	renderer.Clear()

	renderer.Copy(texture, &src, &dst)
	renderer.Present()
}
