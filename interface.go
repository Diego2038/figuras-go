package figuras

import "fmt"

type Medidas interface {
	sacarArea() string
	perimetro() float64
	area() float64
}

func SacarElArea(m Medidas) {
	fmt.Println(m.sacarArea())
}

func SacarPerimetro(m Medidas) {
	fmt.Println("El perímetro es:", m.perimetro())
}

func SacarArea(m Medidas) {
	fmt.Println("El área es:", m.area())
}
