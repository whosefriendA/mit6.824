package main

import (
		"golang.org/x/tour/pic"
	)

	// func Pic(dx, dy int) [][]uint8 {
	// 	result := make([][]uint8, dy)
	// 	for i := range result {
	// 		result[i] = make([]uint8, dx)
	// 		for j := range result[i] {
	// 			result[i][j] = uint8(i * j)
	// 		}
	// 	}
	// 	return result
	// }
	func Pic(dy,dx int)[][]uint8{
		result:=make([][]uint8,dy)
		for i:=range result{
			row:=make([]uint8,dx)
			for j:=range row{
				row=append(row,uint8(i*j))
			}
			result=append(result,row)
		}
		return result
	}
	
func main() {
	pic.Show(Pic)
}