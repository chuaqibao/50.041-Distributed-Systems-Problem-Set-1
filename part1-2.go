package main

// import (
// 	"fmt"
// 	"time"
// 	"strconv"
// 	"math/rand"
// )

// // Lamport's Logical Clock 

// // 1. Receive message. Increment order by 1.
// // 2. Send message to 2 other clients. Increment order by 2, 1 for each client
// // 3. Random delay 
// func server(req1 chan [2]string, req2 chan [2]string, req3 chan [2]string, res1 chan [2]string, res2 chan [2]string, res3 chan [2]string) {
	
// 	// The order increments by 3 at server. 1 for receive, 2 for sending out to 2 different clients 
// 	order := 0

// 	for {
// 		select {
// 			case msg1 := <- req1:
// 				temp_order, _ := strconv.Atoi(msg1[1])
// 				if (temp_order > order) { order = temp_order }
// 				order++
// 				fmt.Printf("(Server, Order %d) Received \"%s\" \n", order, msg1[0])
// 				msg1 = [2]string{msg1[0], strconv.Itoa(order)}
// 				res2 <- msg1
// 				order++
// 				msg1 = [2]string{msg1[0], strconv.Itoa(order)}
// 				res3 <- msg1
// 			case msg2 := <- req2:
// 				temp_order, _ := strconv.Atoi(msg2[1])
// 				if (temp_order > order) { order = temp_order }
// 				order++
// 				fmt.Printf("(Server, Order %d) Received \"%s\" \n", order, msg2[0])
// 				msg2 = [2]string{msg2[0], strconv.Itoa(order)}
// 				res1 <- msg2
// 				order++
// 				msg2 = [2]string{msg2[0], strconv.Itoa(order)}
// 				res3 <- msg2
// 			case msg3 := <- req3:
// 				temp_order, _ := strconv.Atoi(msg3[1])
// 				if (temp_order > order) { order = temp_order }
// 				order++
// 				fmt.Printf("(Server, Order %d) Received \"%s\" \n", order, msg3[0])
// 				msg3 = [2]string{msg3[0], strconv.Itoa(order)}
// 				res1 <- msg3
// 				order++
// 				msg3 = [2]string{msg3[0], strconv.Itoa(order)}
// 				res2 <- msg3
// 		}
// 		order++
// 		// Delays received message while broadcasting to a registered client
// 		amt := time.Duration(rand.Intn(500))
// 		time.Sleep(time.Millisecond * time.Duration(amt))
// 	}
// }

// // Client simultaneously does:
// // 1. Receives response from server
// // 2. Periodically send message to buffer
// func client(id int, req chan [2]string, res chan [2]string) {
// 	order := 0

// 	// receive response from server
// 	go func() {
// 		for {
// 			msg := <- res
// 			temp_order, _ := strconv.Atoi(msg[1])
// 			if (temp_order > order) { order = temp_order }
// 			order++
// 			fmt.Printf("(Client %d, Order %d) Received \"%s\" \n", id, order, msg[0])
// 		}
// 	}()

// 	// periodically send message into buffer
// 	go func() {
// 		for i := 0; ; i++ {
// 			order++
// 			msg := [2]string{"Client " + strconv.Itoa(id) + "'s Message " + strconv.Itoa(i), strconv.Itoa(order)}
// 			fmt.Printf("(Client %d, Order %d) Sending \"%s\" \n", id, order, msg[0])
// 			req <- msg
// 			time.Sleep(time.Millisecond * 1000)
// 		}
// 	}()
// }

// // Creates 
// // 1. 3 "Request" channels
// // 2. 3 "Response" channels
// // 3. 1 Server, 3 Clients
// func main() {

// 	// create 3 request channels and 3 response channels
// 	req1 := make(chan [2]string) 
// 	req2 := make(chan [2]string) 
// 	req3 := make(chan [2]string) 
// 	res1 := make(chan [2]string)
// 	res2 := make(chan [2]string)
// 	res3 := make(chan [2]string)

// 	// create 1 server and 3 clients
// 	go server(req1, req2, req3, res1, res2, res3)
// 	go client(1, req1, res1)
// 	go client(2, req2, res2)
// 	go client(3, req3, res3)

// 	var input string
// 	fmt.Scanln(&input)
// }
