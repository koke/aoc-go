package utils

type Range struct {
	Low, High int
}

func (r Range) Contains(n int) bool {
	return r.Low <= n && n <= r.High
}
