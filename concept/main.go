package main

import (
	"concept/articles"
	"concept/channel"
	"fmt"
)

func Log(str ...string) {
	fmt.Println(str)
}

func main() {

	/* // Runes
	const str = "ABCD EFGH"
	fmt.Println("Rune Count", utf8.RuneCountInString(str))

	// Encoded runes
	for idx, runeValue := range str {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Decode runes
	for i, w := 0, 0; i < len(str); i += w {
		runeVal, width := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("-- %#U start at %d\n", runeVal, i)
		w = width
		fmt.Printf("-- The width after %d & the w after %d \n", width, w)
	} */

	// Channels

	/* type Student struct {
		Name  string
		Roll  string
		Class string
	}
	messages := make(chan Student)
	go func() {
		fmt.Println("Creating student.....")
		time.Sleep(5 * time.Second)
		createdStudent := Student{
			"Sefat Anam",
			"123",
			"Eight",
		}
		messages <- createdStudent
	}()
	msg := <-messages
	json, _ := json.Marshal(&msg)
	fmt.Println(string(json)) */

	// Buffer channel

	/* messages := make(chan string, 2)

	go func() {
		fmt.Println("Printing message...")
		time.Sleep(3 * time.Second)
		messages <- "3 seconds done !"
		time.Sleep(1 * time.Second)
		messages <- "1 seconds done !"
		time.Sleep(5 * time.Second)
		messages <- "5 seconds done !"
		time.Sleep(1 * time.Second)
		messages <- "1 seconds done !"
	}()

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages) */

	// Channel Select

	/* c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			{
				fmt.Println("received", msg1)
			}
		case msg2 := <-c2:
			{
				fmt.Println("received", msg2)
			}
		}

	}*/

	/* // Channel Timeout
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		{
			fmt.Println(res)
		}
	case <-time.After(1 * time.Second):
		{
			fmt.Println("timeout")
		}
	} */

	// Non Blocking Channel
	/* messages := make(chan string)

	signals := make(chan bool)

	select {
	case msg := <-messages:
		{
			fmt.Println("received message", msg)
		}
	default:
		fmt.Println("no message recived")
	}

	msg := "hi"

	select {
	case messages <- msg:
		{
			fmt.Println("sent message", msg)
		}
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		{
			fmt.Println("Received message", msg)
		}
	case sig := <-signals:
		{
			fmt.Println("Received signal", sig)
		}
	default:
		fmt.Println("no activity")
	} */

	/* // Tickers

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
	done <- true

	fmt.Println("Timer stop") */

	/* // Worker pool

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	fmt.Println("** Spinning Up Worker Thread **")
	for w := 1; w <= 3; w++ {
		fmt.Println("Going for worker fn....", w)
		go Worker(w, jobs, results)
		fmt.Println("Going for worker fn....done")
	}

	fmt.Println("** Job Channel Value Asssigning **")
	for j := 1; j <= numJobs; j++ {
		fmt.Println("going through jobs <-", j)
		jobs <- j
		fmt.Println("going through jobs <- done", j)
	}

	fmt.Println("Close job..")
	close(jobs)
	fmt.Println("Close job..END")

	fmt.Println("** Calling Worker Thread to Spinning UP the Works **")
	for r := 1; r <= numJobs; r++ {
		fmt.Println("Results loop", r)
		<-results
		fmt.Println("Result found..: ", r)
	} */

	// Worker Pool 2

	// WorkerPool()

	// Basic Channel
	channel.MainChannel()

	// newNums := channel.NewNumbers()
	// newNums.Start()
	// done := make(chan bool)
	// for i := 1; i <= 10; i++ {
	// 	go func(i int) {
	// 		newNums.NumsChan <- i
	// 		done <- true
	// 	}(i)
	// }
	// for i := 1; i <= 10; i++ {
	// 	<-done
	// }

	// close(newNums.NumsChan)

	// Articles Code
	articles.RunArticles()
}

/* func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Duration(id) * time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- id
	}
}
*/
