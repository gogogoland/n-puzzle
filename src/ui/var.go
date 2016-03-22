/*
** Charles Roy
** croy@student.42.fr
*/

package ui

/*
** Stock section of img.
*/

type Sect struct {
	Xinit, Yinit int32
	Xsrc, Ysrc int32
}

/*
** Init Value size of puzzle, img, and window size
*/

var (
  winTitle string = "N-Puzzle"
  imageName string = "assets/cat.bmp"
  puzzX, puzzY int = 3, 3
  imgX, imgY, winWidth, winHeight int = 1500, 1026, 1850, 1026
	divX, divY int = imgX / puzzX, imgY / puzzY
  err error
)
