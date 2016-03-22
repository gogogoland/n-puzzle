/*
** Charles Roy
** croy@student.42.fr
*/

package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"time"
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
					int32(divX), int32(divY) }
        index++
      }
  }
	renderer.Clear()
	fmt.Println(Piece)

//Test for random number
	list := []int{9,7,5,3,4,2,8,1,6}
//	list := []int{1,2,3,4,5,6,7,8,9}

	for y, k := 1, 0; k < len(list); y++{
		for x := 1; x <= puzzX; x++{
			fmt.Println(k)
			fmt.Println(Piece[list[k]])

			src = sdl.Rect{Piece[list[k]].Xinit, Piece[list[k]].Yinit, Piece[list[k]].Xsrc, Piece[list[k]].Ysrc}
			dst = sdl.Rect{int32(winWidth * (x - 1)/ puzzX), int32(winHeight *  (y - 1) / puzzY),
				int32(divX), int32(divY)}

			time.Sleep(400 * time.Millisecond)
			fmt.Println(dst)
			renderer.Copy(texture, &dst, &src)
			k++
		}
	}
	renderer.Present()
	renderer.Clear()
  sdl.Delay(4000)
}
