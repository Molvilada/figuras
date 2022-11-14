package figuras

type Rectangulo struct {
	Ancho  float64
	Altura float64
}

func (r *Rectangulo) area() float64 {
	return r.Ancho * r.Altura
}

func (r *Rectangulo) perimetro() float64 {
	return 2*r.Ancho + 2*r.Altura
}
