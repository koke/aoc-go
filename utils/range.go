package utils

type Range struct {
	Low, High int
}

func (r Range) Contains(n int) bool {
	return r.Low <= n && n <= r.High
}

func (r Range) Overlaps(other Range) bool {
	return r.Contains(other.Low) || r.Contains(other.High)
}
