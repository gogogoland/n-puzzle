/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   Methods.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbalea <tbalea@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/03/22 18:42:35 by tbalea            #+#    #+#             */
/*   Updated: 2016/03/22 18:42:35 by tbalea           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package algo

/*import (
	"container/heap"
)*/

//	Methods
//	*	linked list for mouvement
type Path struct {
	x, y       int
	next, prev *Path
}

//	*	Puzzle composition and value of it
type Tabl struct {
	table [][]int
	rang  int
	from  int
}

//	*	Slice of Tabl
type List []Tabl

//  *   Functions for method Path
//  *   *   Add Path
//func (cur *Path) AddPath(add *Path) {
//	cur.next = *add
//	add.prev = *cur
//}

//  *   *   Initialaze Path
//func (p Path) InitPath(x, y uint) {
//	p.next = nil
//	p.prev = nil
//	p.x = x
//	p.y = y
//}
