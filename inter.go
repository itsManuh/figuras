package figuras

import "fmt"

type Figura interface {
	CalcularPerimetro() float64
	CalcularArea() float64
}

func Calculos(f Figura) {
	fmt.Println(f)
	fmt.Println("El area es:", f.CalcularArea())
	fmt.Println("El perimetro es:", f.CalcularPerimetro())
}
