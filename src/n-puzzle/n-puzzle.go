/* ROY CHARLES
** croy@student.42.fr
** France
** N-Puzzle project
*/

package main

import (
  "ui"
  "os"
  "fmt"
  "strconv"
)

//Debug
var (
  array = [][]int{{1, 4, 5, 6, 7, 8, 9, 3, 2, 10, 11, 12}, {11, 12, 1, 7, 5, 6, 4, 8, 9, 2, 3, 10}, {11, 12,10, 7, 1, 8, 6, 4, 5, 9, 2, 3}}
)

func check_args() (int, int){
  //check if 2 args
  if len(os.Args) != 3{
    fmt.Println("[Usage]:", os.Args[0], "<file of puzzle> -[your algo]")
    return 1, 0
  }

  //Open file and check if exist.
  file, err := os.Stat(os.Args[1])
  if err != nil || file.IsDir(){
    fmt.Println("[ERROR]: Can't open your File/Directory.")
    return 1, 0
  }

  input, err := strconv.Atoi(string(os.Args[2]))
  if (err != nil) {
    fmt.Println("[ERROR]: Atoi failed for somes obscurs reasons...")
    return 1, 0
  }
  return 0, input
}

func main() {
  err, _ := check_args()

  if err == 0 {
    ui.Ui(array, 2, 5)
  }
}
