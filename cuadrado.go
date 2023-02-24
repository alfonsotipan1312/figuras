package figuras

type Cuadrado struct {
	Lado int
}

func (c *Cuadrado) GetArea() int {
	return c.Lado * 4
}
