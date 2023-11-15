package main

type Point struct{ x, y float64 }

// type
type Rectangle struct{ min, max Point }

// Overlap
func (r1 *Rectangle) Overlap(r2 *Rectangle) bool {
	return r1.min.x < r2.max.x && r1.max.x > r2.min.x &&
		r1.min.y < r2.max.y && r1.max.y > r2.min.y
}

// Inside
func (inner *Rectangle) Inside(outer *Rectangle) bool {
	return outer.min.x <= inner.min.x && outer.min.y <= inner.min.y &&
		outer.max.x >= inner.max.x && outer.max.y >= inner.max.y
}

func (r *Rectangle) Dx() float64 {
	return r.max.x - r.min.x
}

func (r *Rectangle) Dy() float64 {
	return r.max.y - r.min.y
}

// Type
type CollisionBody struct {
	outer      Rectangle   // Wrapper
	rectangles []Rectangle // BaseBlocks
}

// Center the body around an (x, y) point
func (b *CollisionBody) CenterAround(x, y float64) {
	x -= b.outer.Dx() / 2
	y -= b.outer.Dy() / 2

	for i := range b.rectangles {
		r := &b.rectangles[i]

		r.min.x += x
		r.max.x += x

		r.min.y += y
		r.max.y += y
	}
}

func (b1 *CollisionBody) Overlap(b2 *CollisionBody) bool {
	for i := range b1.rectangles {
		for j := range b2.rectangles {
			if b1.rectangles[i].Overlap(&b2.rectangles[j]) {
				return true
			}
		}
	}

	return false
}

func (inner *CollisionBody) Inside(outer *CollisionBody) bool {
	for i := range inner.rectangles {
		for j := range outer.rectangles {
			if !inner.rectangles[i].Inside(&outer.rectangles[j]) {
				return false
			}
		}
	}

	return true
}
