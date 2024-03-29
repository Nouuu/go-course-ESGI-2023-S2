package car

import (
	"time"
)

// Car Structure Car
type Car struct {
	brand  string
	year   int
	color  string
	engine int
}

func New(brand string, year int, color string, engine int) *Car {
	return &Car{brand: brand, year: year, color: color, engine: engine}
}

// Equals compare deux voitures
func (c Car) Equals(other Car) bool {
	return c.brand == other.brand && c.year == other.year
}

// Age retourne l'age de la voiture
func (c Car) Age() int {
	return time.Now().Year() - c.year
}

// UpdateYear met à jour l'année de la voiture
func (c *Car) UpdateYear(newYear int) {
	c.year = newYear
}

// SetColor met à jour la couleur de la voiture
func (c *Car) SetColor(newColor string) {
	c.color = newColor
}

// UpgradeEngine met à jour la puissance du moteur
func (c *Car) UpgradeEngine() {
	c.engine++
}
