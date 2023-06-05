package goarea

import "math"

func Circ(raio float64) float64 {
	return math.Pow(raio, 2)
}

func Rect(base, altura float64) float64 {
	return base * altura
}

func _Triang(base, altura float64) float64 {
	return (base * altura) / 2
}
