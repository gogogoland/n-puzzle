/*
** Charles Roy
** croy@student.42.fr
*/

package ui

/*
** Init Value size of puzzle, img, and window size
*/

var (
  winTitle string = "N-Puzzle"
  puzzX, puzzY int = 3, 3
  winWidth, winHeight int = 1500, 1026
  imageName string = "assets/cat.bmp"
	divX, divY int = winWidth / puzzX, winHeight / puzzY
  err error
)

/*
** Stock section of img.
*/

type Sect struct {
	Xinit, Yinit int32
	Xsrc, Ysrc int32
}
