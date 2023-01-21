package visitor

import "fmt"

type Visitor struct {
	active   uint
	inactive uint
	total    uint
}

type ActiveType string

const (
	JoinActivity = ActiveType("join")
	LeftActivity = ActiveType("left")
)

func New() Visitor {
	return Visitor{
		active:   0,
		inactive: 0,
		total:    0,
	}
}

func (v *Visitor) Join() {
	v.active += 1
	v.total += 1
}

func (v *Visitor) Left() {
	v.inactive += 1
	v.active -= 1
}

func (v *Visitor) ActiveVisitors() uint {
	return v.active
}

func (v *Visitor) left(value uint) {
	v.active -= value
	v.inactive += value
}

func (v *Visitor) join(value uint) {
	v.active += value
	v.total += value
}

func (v *Visitor) BlukActivity(ac string) func(uint2 uint) {
	switch ac {
	case string(JoinActivity):
		return v.join
	case string(LeftActivity):
		return v.left
	default:
		return func(value uint) {
			fmt.Println("Activity not supported")
		}
	}
}
