package figuras

import "fmt"

type Figura interface {
	area() float64
	perimetro() float64
}

func Medidas(f Figura) {
	fmt.Println("Área:", f.area())
	fmt.Println("Perímetro:", f.perimetro())
}
