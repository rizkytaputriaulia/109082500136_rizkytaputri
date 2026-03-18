package main

import "fmt"

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	jarak1 := (x-cx1)*(x-cx1) + (y-cy1)*(y-cy1)
	jarak2 := (x-cx2)*(x-cx2) + (y-cy2)*(y-cy2)

	diDalam1 := jarak1 <= (r1 * r1)
	diDalam2 := jarak2 <= (r2 * r2)

	if diDalam1 && diDalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
