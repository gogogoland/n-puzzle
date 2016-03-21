package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

func Load_image(renderer  *sdl.Renderer) int{
  var image     *sdl.Surface
  var texture   *sdl.Texture

  image, err = sdl.LoadBMP(imageName)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error : %s\n", err)
    return 0
  }
  defer image.Free()

  texture, err = renderer.CreateTextureFromSurface(image)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error : %s\n", err)
    return 0
  }
  defer texture.Destroy()

  TreatDataImageSrc(renderer, texture)

  return (1)
}

func TreatDataImageSrc(renderer *sdl.Renderer, texture *sdl.Texture) {
  var src, dst  sdl.Rect
	var Piece = make(map[int]Sect)

  for y, index := 1, 1; y <= puzzY; y++{
      for x := 1; x <= puzzX; x++{
        Piece[index] = Sect{ int32(winWidth * (x - 1) / puzzX),  int32(winHeight * (y - 1) / puzzY),
					int32(winWidth * x / puzzX), int32(winHeight * y / puzzY) }
        index++
      }
  }
	renderer.Clear()
	src = sdl.Rect{Piece[1].Xinit, Piece[1].Yinit, Piece[1].Xsrc, Piece[1].Ysrc}
	dst = sdl.Rect{int32(winWidth * 1 / puzzX), int32(winHeight *  1 / puzzY), Piece[1].Xsrc, Piece[1].Ysrc}
	renderer.Copy(texture, &src, &dst)
	renderer.Present()
	renderer.Clear()
	fmt.Println(Piece)
  sdl.Delay(4000)
}
