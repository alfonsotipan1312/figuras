package figuras

type Figura interface {
	GetArea() int
}

func GetAreaFigura(c Figura) int {
	return c.GetArea()
}
