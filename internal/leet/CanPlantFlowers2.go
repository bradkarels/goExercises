package main

import "fmt"

func canPlaceFlowerz(flowerbed []int, n int) bool {
	newBed := make([]int, len(flowerbed))
	copy(newBed, flowerbed)
	newBed = append([]int{0}, newBed...)
	newBed = append(newBed, 0)
	nbr := n
	maxIdx := len(newBed) - 3
	idx := 0
	var c, next, last *int
	for {
		if idx > maxIdx || nbr <= 0 {
			break
		}
		c = &newBed[idx]
		next = &newBed[idx+1]
		last = &newBed[idx+2]
		fmt.Println(nbr, idx, *c, *next, *last)
		if *c+*next+*last == 0 {
			nbr--
			idx += 2
			continue
		}
		idx++
	}
	if nbr <= 0 {
		return true
	}
	return false
}
