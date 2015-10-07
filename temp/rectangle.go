package main 

import (
	"fmt"
	"image/color"
)

type rectangle struct {
	x0, y0, x1, y1 	int
	fill			color.RGBA
}

func main() {
	rect := rectangle{4,8, 20,10, color.RGBA{0xFF, 0, 0, 0xFF}}
	fmt.Println(rect)
	resizeRect(&rect, 5, 5)
	fmt.Println(rect)

	y := 10
	upgradeY(&y)
	fmt.Println(y)
}

func resizeRect(rect *rectangle, dx, dy int){
	rect.x1 += dx
	rect.y1 += dy
}

func upgradeY(y *int){
	*y += 10
}