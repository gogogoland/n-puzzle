package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

/*
** Init Value size of puzzle, img, and window size
*/
var (
  winTitle string = "N-Puzzle"
  puzzX, puzzY int = 3, 3
  winWidth, winHeight int = 1496, 1026
  imageName string = "assets/cat.bmp"
	divX int = winWidth / puzzX
	divY int = winHeight / puzzY
  err error
)

/*
** Stock section of img.
*/

type Sect struct {
	Xinit, Yinit int32
	Xsrc, Ysrc int32
}

/*
** Main function
*/
func Window() int{
  var window    *sdl.Window
  var renderer  *sdl.Renderer

  window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
    winWidth, winHeight, sdl.WINDOW_SHOWN)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error : %s\n", err)
    return 0
  }
  defer window.Destroy()

  renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %s\n", err)
    return 0
  }
  defer renderer.Destroy()

  if (Load_image(renderer) != 1) {
    return 0
  }

  return 1
}
