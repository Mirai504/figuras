package figuras

import "math"

type Circulo struct {
	Radio float64 //esta variable dede ser publica ya que vamos a manipular desde otros archivos
}

func (cir *Circulo) area() float64 {
	return math.Pi * (cir.Radio * cir.Radio) //pi es constante y public
}

func (cir *Circulo) perimetro() float64 {
	return 2 * math.Pi * cir.Radio
}
