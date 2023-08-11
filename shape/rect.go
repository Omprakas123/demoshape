package shape

type Rect struct {
	L, B float32
}

// method contains a reciver
func (r *Rect) Area() float32 {
	return r.L * r.B
}

// function does not contain  reciver

func AreaOf(r *Rect) float32 {
	return r.L * r.B
}
