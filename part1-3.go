package main

// import (
// 	"fmt"
// 	"time"
// 	"math/rand"
// )

// // Vector Clock with Detector for Potential Causality Violation

// func server_parse(id int, clock [4]int, msg [6]int) ([6]int, [4]int) {
// 	// update server clock
// 	if (msg[2] > clock[0]) { clock[0] = msg[2] }
// 	if (msg[3] > clock[1]) { clock[1] = msg[3] }
// 	if (msg[4] > clock[2]) { clock[2] = msg[4] }
// 	if (msg[5] > clock[3]) { clock[3] = msg[5] }
// 	clock[3]++

// 	fmt.Printf("(Server, Clock %d) Received Client %d's Message %d \n", clock, msg[0], msg[1])
// 	return [6]int{msg[0], msg[1], clock[0], clock[1], clock[2], clock[3]}, clock
// }

// // 1. Receive message. Increment order by 1.
// // 2. Send message to 2 other clients. Increment order by 2, 1 for each client
// // 3. Random delay
// func server(req1 chan [6]int, req2 chan [6]int, req3 chan [6]int, res1 chan [6]int, res2 chan [6]int, res3 chan [6]int) {
	
// 	// clock[0]: Client 1, clock[1]: Client 2, clock[2]: Client 3, clock[3]: Server
// 	clock := [4]int{0,0,0,0}

// 	for {
// 		select {
// 			case msg1 := <- req1:
// 				msg1, clock = server_parse(1, clock, msg1)
// 				res2 <- msg1
// 				clock[3]++
// 				msg1[5] = clock[3]
// 				res3 <- msg1
// 			case msg2 := <- req2:
// 				msg2, clock =  server_parse(2, clock, msg2)
// 				res1 <- msg2
// 				clock[3]++
// 				msg2[5] = clock[3]
// 				res3 <- msg2
// 			case msg3 := <- req3:
// 				msg3, clock = server_parse(3, clock, msg3)
// 				res1 <- msg3
// 				clock[3]++
// 				msg3[5] = clock[3]
// 				res2 <- msg3
// 		}
// 		clock[3]++
// 		amt := time.Duration(rand.Intn(500))
// 		time.Sleep(time.Millisecond * time.Duration(amt))
// 	}
// }

// // Client simultaneously does:
// // 1. Receives response from server & detects potential causality violation
// // 2. Periodically send message to buffer
// func client(id int, req chan [6]int, res chan [6]int) {

// 	// clock[0]: Client 1, clock[1]: Client 2, clock[2]: Client 3, clock[3]: Server
// 	clock := [4]int{0,0,0,0}

// 	// receive response from server
// 	go func() {
// 		for {
// 			msg := <- res
// 			// detect potential causality violation
// 			if (clock[msg[0]-1] > msg[msg[0]+1]) { fmt.Println("Potential Causality Violation Detected.")}
// 			//update clock
// 			if (msg[2] > clock[0]) { clock[0] = msg[2] }
// 			if (msg[3] > clock[1]) { clock[1] = msg[3] }
// 			if (msg[4] > clock[2]) { clock[2] = msg[4] }
// 			if (msg[5] > clock[3]) { clock[3] = msg[5] }
// 			clock[id-1]++
// 			fmt.Printf("(Client %d, Clock %v) Received Client %d's Message %d \n", id, clock, msg[0], msg[1])
// 		}
// 	}()

// 	// periodically send message into buffer
// 	go func() {
// 		for i := 0; ; i++ {
// 			clock[id-1]++
// 			msg := [6]int{id, i, clock[0], clock[1], clock[2], clock[3]}
// 			fmt.Printf("(Client %d, Clock %v) Sending Message %d \n", id, clock, i)
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
// 	req1 := make(chan [6]int) 
// 	req2 := make(chan [6]int) 
// 	req3 := make(chan [6]int) 
// 	res1 := make(chan [6]int)
// 	res2 := make(chan [6]int)
// 	res3 := make(chan [6]int)

// 	// create 1 server and 3 clients
// 	go server(req1, req2, req3, res1, res2, res3) // server listens to all channels 
// 	go client(1, req1, res1) 
// 	go client(2, req2, res2)
// 	go client(3, req3, res3)

// 	var input string
// 	fmt.Scanln(&input)
// }
