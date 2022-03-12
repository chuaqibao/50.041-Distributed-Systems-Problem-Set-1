package main

// import (
// 	"fmt"
// 	"time"
// 	"strconv"
// 	"math/rand"
// )

// func server(req1 chan string, req2 chan string, req3 chan string, res1 chan string, res2 chan string, res3 chan string) {

// 	// 1. Receive message
// 	// 2. Send message to other clients
// 	// 3. Random delay 
// 	for {
// 		select {
// 			case msg1 := <- req1:
// 				fmt.Printf("Server received from %s \n", msg1)
// 				res2 <- msg1
// 				res3 <- msg1
// 			case msg2 := <- req2:
// 				fmt.Printf("Server received from %s \n", msg2)
// 				res1 <- msg2
// 				res3 <- msg2
// 			case msg3 := <- req3:
// 				fmt.Printf("Server received %s \n", msg3)
// 				res1 <- msg3
// 				res2 <- msg3
// 		}
// 		amt := time.Duration(rand.Intn(500))
// 		time.Sleep(time.Millisecond * time.Duration(amt))
// 	}
// }

// // Client simultaneously does:
// // 1. Receives response from server
// // 2. Periodically send message to buffer
// func client(id int, req chan string, res chan string) {
// 	// receive response from server
// 	go func() {
// 		for {
// 			msg := <- res
// 			fmt.Printf("Client %d received from %s \n", id, msg)
// 		}
// 	}()

// 	// periodically send message into buffer
// 	go func() {
// 		for i := 1; ; i++ {
// 		fmt.Printf("Client %d sending value %d \n", id, i)
// 		req <- "Client " + strconv.Itoa(id) + " value " + strconv.Itoa(i)
// 		time.Sleep(time.Millisecond * 1000)
// 		}
// 	}()

// }

// // Creates 
// // 1. 3 "Request" channels
// // 2. 3 "Response" channels
// // 3. 1 Server, 3 Clients
// func main() {

// 	// create 3 request channels and 3 response channels
// 	req1 := make(chan string) 
// 	req2 := make(chan string) 
// 	req3 := make(chan string) 

// 	res1 := make(chan string)
// 	res2 := make(chan string)
// 	res3 := make(chan string)

// 	// create 1 server and 3 clients
// 	go server(req1, req2, req3, res1, res2, res3)
// 	go client(1, req1, res1)
// 	go client(2, req2, res2)
// 	go client(3, req3, res3)

// 	var input string
// 	fmt.Scanln(&input)
// }