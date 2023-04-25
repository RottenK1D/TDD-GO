package models

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rect Rectangle) float64 {
	return (rect.Width + rect.Height) * 2
}

func Area(area Rectangle) float64 {
	return area.Width * area.Height
}
