package typeswitch

import "fmt"

type Ninja struct {
	Name      string
	Specialty string
	Level     int
}

func (Ninja) attack() {
	fmt.Println("Ninja attacking....")
}

type Newbie struct {
	Name  string
	Age   int
	Track int
}

func (Newbie) attack() {
	fmt.Println("Newbie attacking....")
}

func (Newbie) pickUp() {
	fmt.Println("Newbie pick up sword....")
}

type Fighter interface {
	attack()
}

func TypeSwitchMain() {
	fighters := []Fighter{
		Ninja{"Wallace 1", "Sword", 1},
		Ninja{"Wallace 2", "Sword", 3},
		Ninja{"Wallace 3", "Sword", 4},
		Newbie{"Bew 3", 28, 4},
		Newbie{"Eew 3", 23, 4},
	}
	for _, fighter := range fighters {
		switch fighter.(type) {
		case Newbie:
			nf, ok := fighter.(Newbie)
			if !ok {
				fmt.Println("Error: Unable to type assert to Newbie")
				continue
			}
			nf.attack()
			nf.pickUp()
		case Ninja:
			n, ok := fighter.(Ninja)
			if !ok {
				fmt.Println("Error: Unable to type assert to Ninja")
				continue
			}
			n.attack()
		default:
			fmt.Println("Doing nothing....")
		}
	}

	/* for _, fighter := range fighters {

		switch f := fighter.(type) {
		case Newbie:
			nf := f.(Newbie)
			nf.attack()
			nf.pickUp()
		case Ninja:
			fighter.attack()
		default:
			fmt.Println("Doing nothing....")
		}

	} */
}
