package main

func maxArea(height []int) int {
	l := len(height)
	maxVol := 0
	mvp := &maxVol
	for i := l - 1; i > -1; i-- {
		heightAtI := height[i]
		if heightAtI == 0 {
			continue
		}
		for idx := 0; idx < i; idx++ {
			heightAtIdx := height[idx]
			if heightAtIdx == 0 {
				continue
			}
			width := i - idx
			if heightAtI > heightAtIdx {
				volume := heightAtIdx * width
				if volume > *mvp {
					*mvp = volume
				}
				continue
			}
			volume := heightAtI * width
			if volume > *mvp {
				*mvp = volume
			}
		}
	}
	return maxVol
}

func maxArea2(height []int) int {
	m := 0
	l := len(height)
	var h, w []int
	for ir := l - 1; ir > -1; ir-- {
		hir := height[ir]
		if hir == 0 {
			continue
		}
		for il := 0; il < ir; il++ {
			hil := height[il]
			if hil == 0 {
				continue
			}

			w = append(w, ir-il)
			if hir > hil {
				h = append(h, hil)
			} else {
				h = append(h, hir)
			}
		}
	}
	for i, h := range h {
		v := h * w[i]
		if v > m {
			m = v
		}
	}
	return m
}
