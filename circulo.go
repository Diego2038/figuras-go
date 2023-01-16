package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	Radio float64
}

func (c Circulo) sacarArea() string {
	return fmt.Sprintf("El circulo tiene %v metros de radio", c.Radio)
}

func (c Circulo) perimetro() float64 {
	return c.Radio * math.Pi * 2
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.Radio, 2)
}
