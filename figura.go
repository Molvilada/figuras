package figuras

import "fmt"

type figura interface {
	area() float64
	perimetro() float64
}

func Medidas(f figura) {
	fmt.Println("Área:", f.area())
	fmt.Println("Perímetro:", f.perimetro())
}
