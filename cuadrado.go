package figuras

import (
	"fmt"
	"math"
)

type Cuadrado struct {
	Lado float64
}

func (c Cuadrado) sacarArea() string {
	return fmt.Sprintf("El círculo tiene %v en cada lado", c.Lado)
}

func (c Cuadrado) perimetro() float64 {
	return c.Lado * 4
}

func (c Cuadrado) area() float64 {
	return math.Pow(c.Lado, 2)
}
