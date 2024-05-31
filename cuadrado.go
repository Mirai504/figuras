package figuras

type Cuadrado struct {
	Lado float64 //Esta variable debe ser publica porque vamos a manipular desde otros archivos
}

func (cua *Cuadrado) area() float64 { //recibe una variable de tipo cuadrado y retorna un float
	return cua.Lado * cua.Lado
}

func (cua *Cuadrado) perimetro() float64 {
	return 4 * cua.Lado
}
