// Interface
package figuras

import "fmt"

type Geometria interface {
	area() float64
	perimetro() float64
}

// funcion para manejar metodos por interface
func Medidas(g Geometria) { //funcion Publica con M mayuscula porque esta funcion abriremos desde otros archivos
	fmt.Println("Medidas: ", g)
	fmt.Println(g.perimetro())
	fmt.Println(g.area())
}
