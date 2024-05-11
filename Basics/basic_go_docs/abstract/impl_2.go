package abstract

import "fmt"

type IDoor interface {
	getWidth() int
	getHeight() int
	getColor() string
}

type DoorFittingExpert interface {
	getDescription() string
}

type Welder struct {
	comment string
}

func (w *Welder) getDescription() string {
	return w.comment
}

type Carpenter struct {
	comment string
}

func (c *Carpenter) getDescription() string {
	return c.comment
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
	MakeFittingExpert() DoorFittingExpert
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

func (wdf *WoodenDoorFactory) MakeFittingExpert() DoorFittingExpert {
	return &Carpenter{
		comment: "I can only fit wood door.",
	}
}

func (mdf *MetalDoorFactory) MakeDoor(width, height int, color string) IDoor {
	return &MetalDoor{
		width:  width,
		height: height,
		color:  color,
	}
}
func (wdf *MetalDoorFactory) MakeFittingExpert() DoorFittingExpert {
	return &Welder{
		comment: "I can only fit iron door.",
	}
}
func RunImpl2() {
	var woodDoorFactory WoodenDoorFactory
	woodDoor := woodDoorFactory.MakeDoor(1, 1, "brown")

	fmt.Println(woodDoor)
	fmt.Println(woodDoorFactory.MakeFittingExpert())

	var metalDoorFactory MetalDoorFactory
	metalDoor := metalDoorFactory.MakeDoor(1, 1, "red")
	fmt.Println(metalDoor)
	fmt.Println(metalDoorFactory.MakeFittingExpert())

}
