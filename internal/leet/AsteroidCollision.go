package main

import "fmt"

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func asteroidCollision(asteroids []int) []int {
	var idx = 0
	for {
		collision := false
		if idx < len(asteroids)-1 {
			current := asteroids[idx]
			next := asteroids[idx+1]
			if current < 0 && next > 0 {
				idx++
				continue
			} else if current < 0 && next < 0 {
				idx++
				continue
			} else if current > 0 && next > 0 {
				idx++
				continue
			} else { // current > 0 && next < 0 - will collide
				absCurrent := abs(current)
				absNext := abs(next)
				if absCurrent == absNext {
					collision = true
					fmt.Printf("%d and %d will be destroyed...\n", current, next)
					asteroids = append(asteroids[:idx], asteroids[idx+2:]...)
					if idx > 0 {
						idx--
					}
				} else if absCurrent > absNext {
					collision = true
					fmt.Printf("%d will destroy %d...\n", current, next)
					asteroids = append(asteroids[:idx+1], asteroids[idx+2:]...)
				} else {
					collision = true
					fmt.Printf("%d will be destroyed by %d...\n", current, next)
					asteroids = append(asteroids[:idx], asteroids[idx+1:]...)
					if idx > 0 {
						idx--
					}
				}
			}
		} else {
			break
		}
		if collision {
			continue
		}
		break
	}
	return asteroids
}
