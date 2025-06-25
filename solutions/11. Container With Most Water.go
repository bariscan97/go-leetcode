type Area struct {
	Area int
	L    int
	R    int
}

func maxArea(height []int) int {

	a := new(Area)
	a.R = len(height) - 1

	for a.L < a.R {
		if height[a.L] < height[a.R] {
			a.setArea(height[a.L])
			a.L++
		} else {
			a.setArea(height[a.R])
			a.R--
		}
	}

	return a.Area
}

func (a *Area) setArea(h int) {
	if a.Area < h*(a.R-a.L) {
		a.Area = h * (a.R - a.L)
	}
}