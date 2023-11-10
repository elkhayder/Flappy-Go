package main

type Point struct{ x, y float64 }
type CollisionBody struct{ min, max Point }

// Overlap
func (b1 *CollisionBody) Overlap(b2 *CollisionBody) bool {
	return b1.min.x < b2.max.x && b1.max.x > b2.min.x &&
		b1.min.y < b2.max.y && b1.max.y > b2.min.y
}

// Inside
func (inner *CollisionBody) Inside(outer *CollisionBody) bool {
	return outer.min.x <= inner.min.x && outer.min.y <= inner.min.y &&
		outer.max.x >= inner.max.x && outer.max.y >= inner.max.y
}
