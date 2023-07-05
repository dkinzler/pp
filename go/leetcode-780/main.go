package main

import "fmt"

func main() {
	sx, sy, tx, ty := 1, 1, 3, 5
	fmt.Println(reachingPoints(sx, sy, tx, ty))
	sx, sy, tx, ty = 1, 1, 2, 2
	fmt.Println(reachingPoints(sx, sy, tx, ty))
	sx, sy, tx, ty = 1, 1, 1, 1
	fmt.Println(reachingPoints(sx, sy, tx, ty))
	sx, sy, tx, ty = 4, 2, 2, 4
	fmt.Println(reachingPoints(sx, sy, tx, ty))
	sx, sy, tx, ty = 1, 16, 999999986, 16
	fmt.Println(reachingPoints(sx, sy, tx, ty))
	sx, sy, tx, ty = 3, 3, 12, 9
	fmt.Println(reachingPoints(sx, sy, tx, ty))
	sx, sy, tx, ty = 9, 10, 9, 19
	fmt.Println(reachingPoints(sx, sy, tx, ty))
}

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	// start with (sx, sy) want (tx, ty)
	// for (x, y) can either do (x, x+y) or (y, x+y)
	// sx, sy, tx, ty >= 1
	// lets look at it in reverse
	// we want to end with (tx, ty) assume tx < ty
	// what happens in the last move?
	// (x, y) -> (x, x+y) = (tx, ty) or (y, x+y) = (tx, ty)
	// i.e. one of the numbers is tx and the other is ty-tx
	// then we can do the same thing again?
	x, y := tx, ty
	for x >= sx && y >= sy {
		if x == sx && y == sy {
			return true
		}
		if x == 1 || y == 1 {
			return true
		}
		if x == y {
			return false
		}
		if x > y {
			if sx > y {
				v := (x - sx) / y
				if v == 0 {
					return false
				}
				x = x - v*y
			} else {
				if x%y == 0 {
					x = y
				} else {
					x = x % y
				}
			}
		} else {
			if sy > x {
				v := (y - sy) / x
				if v == 0 {
					return false
				}
				y = y - v*x
			} else {
				if y%x == 0 {
					y = x
				} else {
					y = y % x
				}
			}
		}
	}
	return false
}
