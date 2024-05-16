package channel

import (
	"fmt"
	"log"
	"time"
)

func MainChannel() {
	log.Println("Program execution is start")

	ch := make(chan int, 2)
	time.Sleep(5 * time.Second)
	ch <- 2
	time.Sleep(5 * time.Second)
	ch <- 3
	close(ch)
	/* go func() {
		ch <- 42
	}()
	go func() {
		ch <- 99
	}()
	*/
	for {
		val, ok := <-ch
		if !ok {
			log.Println("Channel Closed!")
			break
		}
		log.Println(val)
	}

	/* ch := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	log.Println("Program execution is start")
	go func() {
		ch <- "1 :: ONE"
		close(ch)
	}()
	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- "2 :: TWO"
		close(ch2)
	}()
	go func() {
		ch3 <- "3 :: THREE"
		close(ch3)
	}()

	log.Println(<-ch3)
	log.Println(<-ch)
	log.Println(<-ch2)

	log.Println("Program execution is complete") */
}

func SendMessage(message chan<- string) {
	message <- "Hello" // OK assign value
	// msg := <-message   // NOK cannot receive from send-only channel message
}

func ReceiveMessage(message <-chan string) {
	msg := <-message // OK read value
	// message <- "Hellow" // NOK cannot send to receive-only channel message
	fmt.Println(msg)
}

func ForwardMessage(message chan string) {
	message <- "Hellow" // OK
	msg := <-message    // OK
	fmt.Println(msg)
}
