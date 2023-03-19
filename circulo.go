package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (c *Circulo) CalcularPerimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func (c *Circulo) CalcularArea() float64 {
	return math.Pi * (c.Radio * c.Radio)
}
