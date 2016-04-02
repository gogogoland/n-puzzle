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
	winWidth, winHeight int = 800, 600
	winTitle string = "N-Puzzle"
	imgName string = "../assets/cat.bmp"
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

	image, err = sdl.LoadBMP(imgName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s\n", err)
		return 3
	}
	defer image.Free()

	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return 4
	}
	defer texture.Destroy()

	src = sdl.Rect{0, 0, 1495, 1026}
	dst = sdl.Rect{0, 0, 800, 600}

	renderer.Clear()

	renderer.Copy(texture, &src, &dst)
	renderer.Present()

	sdl.Delay(4000)

	return 0
}

func Draw(){

}
