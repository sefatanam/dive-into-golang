package abstract

import "fmt"

type Door interface {
	getWidth() int
	getHeight() int
}

type WoodenDoor struct {
	width  int
	height int
}

func (wd WoodenDoor) getWidth() int {
	return wd.width
}

func (wd WoodenDoor) getHeight() int {
	return wd.height
}

type DoorFactory struct {
}

func (df DoorFactory) makeDoor(width int, height int) Door {
	return &WoodenDoor{
		width:  width,
		height: height,
	}
}

func RunImpl1() {
	var df DoorFactory
	door := df.makeDoor(1, 1)
	door.getHeight()
	fmt.Println(door)
}
