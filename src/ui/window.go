/*
** Charles Roy
** croy@student.42.fr
*/

package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

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
