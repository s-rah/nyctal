package utils

import "image"

func IntersectsRect(mx int, my int, rect image.Rectangle) (intersects bool, ix int, iy int) {
	if mx >= rect.Min.X && my >= rect.Min.Y && mx < rect.Max.X && my < rect.Max.Y {
		return true, mx - rect.Min.X, my - rect.Min.Y
	}
	return false, -1, -1
}
