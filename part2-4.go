package main

// import (
// 	"fmt"
// 	"time"
// 	// "os"
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
// 	for i := id; i < len(election_chans); i++ {
// 		fmt.Println("Node", id, "broadcasting election to Node", i+1)
// 		election_chans[i] <- id
// }
// }

// func normal_node(new bool, id int, election_chans []chan int, reject_chans []chan int, msg_chans []chan int, notif_chan chan notification) {
// 	coordinator_id := 5
// 	election_ongoing := false

// 	// emulate faulty coordinator behavior
// 	if id == coordinator_id {
// 		emulate_faulty_coordinator(id, coordinator_id, msg_chans)
// 		fmt.Println("(Break Down) Coordinator Node 5 broke down")	
// 		return
// 	}

// 	if new {
// 		go func() {
// 			// Send all info of new node to all other nodes 
// 			fmt.Println("(Notification) New node sending info to other nodes...")
// 			notif := notification{id, election_chans, reject_chans, msg_chans}
// 			for i := 0; i < len(msg_chans)-1; i++ {
// 					notif_chan <- notif
// 			}
// 		}()
	
// 		time.Sleep(time.Second * 3)
// 		fmt.Println("New node calling election to all nodes")
// 		// Call for elections
// 		// Broadcast election to all other nodes with higher ids
// 		go func() {
// 			election_ongoing = true
// 			for i := 0; i < len(election_chans); i++ {
// 				fmt.Println("Node", id, "broadcasting election to Node", i+1)
// 				election_chans[i] <- id
// 			}
// 		}()
// 	}

// 	// periodically send messages to coordinator
// 	go func() {
// 			for { 
// 				if id != coordinator_id && id != 6 {
// 					fmt.Println("(Sent) Node", id, "sent message to", coordinator_id)
// 					msg_chans[coordinator_id-1] <- id 
// 					time.Sleep(time.Millisecond * 500)
// 			}
// 		}
// 	}()

// 	// listen for election replies
// 	go func() {
// 		for { 
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
// 					go func() {
// 						<-msg_chans[old_coordinator_id-1]
// 					}()
// 					for i := 0; i < 5; i++ {
// 						if i != coordinator_id-1 {
// 							election_chans[i] <- id + 10
// 						}
// 					}
// 				}
// 				election_ongoing = false
// 			}
// 		}
// 	}()

// 	// get notified on new node
// 	go func() {
// 		for {
// 			new_node := <-notif_chan
// 			election_chans = new_node.election_chans
// 			reject_chans = new_node.reject_chans
// 			msg_chans = new_node.msg_chans
// 			time.Sleep(10*time.Millisecond)
// 			fmt.Println("Node", id, "received notification on New Node", new_node.id)
// 		}
// 	}()

// 	// receive messages
// 	go func() {
// 		for {
// 			if (id > 5) { // for new nodes
// 				sender_id :=  <- msg_chans[id-1]
// 				if (coordinator_id == id) {
// 					fmt.Println("(Received) Node", id, "received message from Node", sender_id)
// 					msg_chans[sender_id-1] <- ACK
// 				} else {
// 					fmt.Println("(Replied) Node", id, "received reply from Node", coordinator_id)
// 				}
// 			} else { // for existing nodes
// 				select {
// 				case sender_id := <- msg_chans[id-1]:
// 					if (coordinator_id == id) {
// 						fmt.Println("(Received) Node", id, "received message from Node", sender_id)
// 						msg_chans[sender_id-1] <- ACK
// 					} else {
// 						fmt.Println("(Replied) Node", id, "received reply from Node", coordinator_id)
// 					}
// 				case <-time.After(time.Second * 1):
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
// 		for {
// 			caller_id := <-election_chans[id-1]
// 			fmt.Println("Node", id, "received election call from Node", caller_id)
// 			if caller_id > 10 { // caller_id > 10: is elected node
// 				old_coordinator_id := *&coordinator_id
// 				coordinator_id = caller_id - 10
// 				<-msg_chans[old_coordinator_id-1]
// 				fmt.Println("(New Coordinator) Node", id, "is notified that Node", coordinator_id, "is now the new coordinator")
// 			} else if caller_id < id && caller_id != id { // caller_id < id: reject
// 				reject_chans[caller_id-1] <- REJECT
// 				fmt.Println("(Reject) Node", id, "rejected Node", caller_id)

// 				if (!election_ongoing) {
// 					fmt.Println("(New Election) Node", id, "calling for election...")
// 					election_ongoing = true
// 					go broadcast_elections(id, election_chans) // broadcast to nodes of higher ids 
// 				}

// 			}
// 		}
// 	}()

// }

// type notification struct {
// 	id int
// 	election_chans []chan int
// 	reject_chans []chan int
// 	msg_chans []chan int
// }

// func add_node(id int, election_chans []chan int, reject_chans []chan int, msg_chans []chan int, notif_chan chan notification) {
// 	// create new election, reject and message channel

// 	election_chans = append(election_chans, make(chan int, 5))
// 	reject_chans = append(reject_chans, make(chan int, 5))
// 	msg_chans = append(msg_chans, make(chan int, 5))

// 	// Create new node
// 	go normal_node(true, id, election_chans, reject_chans, msg_chans, notif_chan)

// }

// func main() {

// 	// create 5 election channels, 5 reject channels and 5 message channels
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

// 	msg_chan1 := make(chan int) 
// 	msg_chan2 := make(chan int)
// 	msg_chan3 := make(chan int)
// 	msg_chan4 := make(chan int)
// 	msg_chan5 := make(chan int)

// 	notif_chan := make(chan notification)

// 	election_chans := []chan int{election_chan1, election_chan2, election_chan3, election_chan4, election_chan5}
// 	reject_chans := []chan int {reject_chan1, reject_chan2, reject_chan3, reject_chan4, reject_chan5}
// 	msg_chans := []chan int{msg_chan1, msg_chan2, msg_chan3, msg_chan4, msg_chan5}

// 	go normal_node(false, 1, election_chans, reject_chans, msg_chans, notif_chan)
// 	go normal_node(false, 2, election_chans, reject_chans, msg_chans, notif_chan)
// 	go normal_node(false, 3, election_chans, reject_chans, msg_chans, notif_chan)
// 	go normal_node(false, 4, election_chans, reject_chans, msg_chans, notif_chan)
// 	go normal_node(false, 5, election_chans, reject_chans, msg_chans, notif_chan)

// 	time.Sleep(3 * time.Second)
// 	go add_node(6, election_chans, reject_chans, msg_chans, notif_chan)

// 	var input string
// 	fmt.Scanln(&input)
// }
