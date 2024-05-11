package abstract

import "fmt"

type IDoor interface {
	getWidth() int
	getHeight() int
	getColor() string
}

type WoodenDoorx struct {
	width  int
	height int
	color  string
}

func (wd *WoodenDoorx) getWidth() int {
	return wd.width
}

func (wd *WoodenDoorx) getHeight() int {
	return wd.height
}

func (wd *WoodenDoorx) getColor() string {
	return wd.color
}

func (wd *WoodenDoorx) detail() string {
	return "There is a lot of details"
}

type MetalDoor struct {
	width  int
	height int
	color  string
}

func (md *MetalDoor) getWidth() int {
	return md.width
}
func (md *MetalDoor) getHeight() int {
	return md.height
}
func (md *MetalDoor) getColor() string {
	return md.color
}

type DoorInterface interface {
	MakeDoor(width, height int, color string) IDoor
}

type WoodenDoorFactory struct{}
type MetalDoorFactory struct{}

func (wdf *WoodenDoorFactory) MakeDoor(width, height int, color string) IDoor {
	return &WoodenDoorx{
		width:  width,
		height: height,
		color:  color,
	}
}

func (mdf *MetalDoorFactory) MakeDoor(width, height int, color string) IDoor {
	return &MetalDoor{
		width:  width,
		height: height,
		color:  color,
	}
}

func RunImpl2() {
	var woodDoorFactory WoodenDoorFactory
	woodDoor := woodDoorFactory.MakeDoor(1, 1, "brown")
	woodDoor.(*WoodenDoorx).detail()

	fmt.Println(woodDoor)

	var metalDoorFactory MetalDoorFactory
	metalDoor := metalDoorFactory.MakeDoor(1, 1, "red")
	fmt.Println(metalDoor)

}
