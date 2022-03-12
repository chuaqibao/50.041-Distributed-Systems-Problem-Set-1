package main

// import (
// 	"fmt"
// 	"time"
// )

// const ACK = 200
// const REJECT = 0

// func emulate_faulty_coordinator(id int, coordinator_id int, msg_chans []chan int) {
// 	for i := 0; i < 3; i++ {
// 		sender_id := <-msg_chans[coordinator_id-1]
// 		fmt.Println("(Received) Coordinator Node", coordinator_id, "received message from Node", sender_id)
// 		msg_chans[sender_id-1] <- ACK
// 	}
// }

// func broadcast_elections(id int, election_chans []chan int) {
// 	// Broadcast election to all other nodes with higher ids
// 	for i := id + 1; i <= 5; i++ {
// 		//fmt.Println("Node", id, "broadcasting election to Node", i)
// 		election_chans[i-1] <- id
// 	}
// }

// func normal_node(id int, election_chans []chan int, reject_chans []chan int, msg_chans []chan int) {
// 	coordinator_id := 5
// 	election_ongoing := false

// 	// close_chan := make(chan int)
// 	node_alive := true

// 	// emulate faulty coordinator behavior
// 	if id == coordinator_id {
// 		emulate_faulty_coordinator(id, coordinator_id, msg_chans)
// 		fmt.Println("(Break Down) Coordinator Node 5 broke down")
// 		return
// 	}

// 	// periodically send messages to coordinator
// 	go func() {
// 		for node_alive {
// 			if id != coordinator_id {
// 				fmt.Println("(Sent) Node", id, "sent message to", coordinator_id)
// 				msg_chans[coordinator_id-1] <- id
// 				// Make this value as small as possible to avoid Timeouts in stable state. The Timeouts only because context switching doesn't switch to some nodes :(
// 				time.Sleep(time.Millisecond*10) 
// 			}	
// 		}
// 	}()

// 	// listen for election replies
// 	go func() {
// 		for node_alive {
// 			if election_ongoing {
// 				select {
// 				case <-reject_chans[id-1]:
// 					// Stop election process
// 					//fmt.Println("Node", id, "rejected as coordinator.")
// 				case <-time.After(1 * time.Second): // accepted
// 					// Set self as coordinator and broadcast
// 					fmt.Println("(Elected) Node", id, "elected as coordinator. Broadcasting...")
// 					old_coordinator_id := *&coordinator_id
// 					coordinator_id = id
// 					//fmt.Println("old_coordinator_id", old_coordinator_id, "coordinator_id", coordinator_id)
// 					<-msg_chans[old_coordinator_id-1]

// 					for i := 0; i < 5; i++ {
// 						if i != coordinator_id-1 {
// 							election_chans[i] <- id + 5
// 						}
// 					}
// 				}
// 				election_ongoing = false
// 			}
// 		}
// 	}()

// 	// receive messages
// 	go func() {
// 		for node_alive {
// 			select {
// 			case sender_id := <-msg_chans[id-1]:
// 				if coordinator_id == id {
// 					fmt.Println("(Received) Node", id, "received message from Node", sender_id)
// 					msg_chans[sender_id-1] <- ACK
// 				} else {
// 					fmt.Println("(Replied) Node", id, "received reply from Node", coordinator_id)
// 				}
// 			case <-time.After(time.Second * 1):
// 				if (coordinator_id != id)  {
// 					fmt.Println("(Timeout) No response from Node", coordinator_id, "Node", id, "Calling election...")
// 					if !election_ongoing {
// 						election_ongoing = true
// 						go broadcast_elections(id, election_chans)
// 					}
// 				}
				
// 			}
// 		}
// 	}()

// 	// listen to election channel
// 	go func() {
// 		for node_alive {
// 			caller_id := <-election_chans[id-1]
// 				//fmt.Println("Node", id, "received election call from Node", caller_id)
// 				if node_alive {
// 					if caller_id > 5 { // caller_id > 5: is elected node
// 						old_coordinator_id := *&coordinator_id
// 						coordinator_id = caller_id - 5
// 						fmt.Println("(New Coordinator) Node", id, "is notified that Node", coordinator_id, "is now the new coordinator")
// 						<-msg_chans[old_coordinator_id-1]
// 					} else if caller_id < id && caller_id != id { // caller_id < id: reject
// 						reject_chans[caller_id-1] <- REJECT
// 						fmt.Println("(Reject) Node", id, "rejected Node", caller_id)
// 						if !election_ongoing {
// 							fmt.Println("(New Election) Node", id, "calling for election...")
// 							election_ongoing = true
// 							go broadcast_elections(id, election_chans) // broadcast to nodes of higher ids
// 						}
// 					}
// 				}
// 		}
// 	}()

// 	for {
// 		if election_ongoing && id == 1 {
// 			node_alive = false
// 			fmt.Println("I killed Node 1")
// 			return
// 		}
// 	}
// }

// func main() {

// 	// create 5 election channels, 5 reject channels and 5 node-specific channels
// 	election_chan1 := make(chan int, 5)
// 	election_chan2 := make(chan int, 5)
// 	election_chan3 := make(chan int, 5)
// 	election_chan4 := make(chan int, 5)
// 	election_chan5 := make(chan int, 5)

// 	reject_chan1 := make(chan int, 5)
// 	reject_chan2 := make(chan int, 5)
// 	reject_chan3 := make(chan int, 5)
// 	reject_chan4 := make(chan int, 5)
// 	reject_chan5 := make(chan int, 5)

// 	chan1 := make(chan int)
// 	chan2 := make(chan int)
// 	chan3 := make(chan int)
// 	chan4 := make(chan int)
// 	chan5 := make(chan int)

// 	election_chans := []chan int{election_chan1, election_chan2, election_chan3, election_chan4, election_chan5}
// 	reject_chans := []chan int{reject_chan1, reject_chan2, reject_chan3, reject_chan4, reject_chan5}
// 	msg_chans := []chan int{chan1, chan2, chan3, chan4, chan5}

// 	go normal_node(1, election_chans, reject_chans, msg_chans)
// 	go normal_node(2, election_chans, reject_chans, msg_chans)
// 	go normal_node(3, election_chans, reject_chans, msg_chans)
// 	go normal_node(4, election_chans, reject_chans, msg_chans)
// 	go normal_node(5, election_chans, reject_chans, msg_chans)

// 	// time.Sleep(time.Second * 5)
// 	// go new_node(6, election_chans, reject_chans, msg_chans)

// 	var input string
// 	fmt.Scanln(&input)
// }