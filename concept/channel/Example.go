package channel

import (
	"fmt"
)

type Numbers struct {
	Nums     map[int]string
	NumsChan chan int
}

func NewNumbers() *Numbers {
	return &Numbers{
		Nums:     map[int]string{},
		NumsChan: make(chan int),
	}
}

func (n *Numbers) Add(num int) {
	n.Nums[num] = fmt.Sprintf("Added num is %d", num)
}

func (n *Numbers) Start() {
	go n.Loop()
}

func (n *Numbers) Loop() {
	for {
		fmt.Println("Executin...")
		num := <-n.NumsChan
		n.Add(num)
	}

	/* select {
	case num := <-n.NumsChan:
		{
			if num > 5 {
				log.Println("number is greater than 5")
			}
			n.Add(num)
		}

	default:
		{
			log.Println("executing just like default..")
		}
	} */
}
