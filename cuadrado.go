package figuras

type Cuadrado struct {
	Lado float64
}

func (c *Cuadrado) CalcularPerimetro() float64 {
	return 4 * c.Lado
}

func (c *Cuadrado) CalcularArea() float64 {
	return c.Lado * c.Lado
}
