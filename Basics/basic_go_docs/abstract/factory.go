package abstract

import "fmt"

// Product Interface
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Concrete Product
type Gun struct {
	name  string
	power int
	years int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
	g.years = power * power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

// Concrete Product
type AK47 struct {
	Gun
	level int
}

func newAk47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47",
			power: 100,
		},
		level: 12,
	}
}

func (ak *AK47) getLevel() int {
	return ak.level
}

// Concrete Product
type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket",
			power: 23,
		},
	}
}

// Gun factory

func getGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("wrong gun type")
	}
}

// Details Printy

func printDetail(gun IGun) {
	fmt.Printf("Gun : %s \n", gun.getName())
	fmt.Printf("Power : %d \n", gun.getPower())
}

// Client

func TestClient() {
	ak47, _ := getGun("ak47")
	printDetail(ak47)
	if isAk47, ok := ak47.(*AK47); ok {
		fmt.Println("this gun power is :", isAk47.getLevel())
	}
	musket, _ := getGun("musket")
	printDetail(musket)

}
