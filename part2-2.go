package main

// import (
// 	"fmt"
// 	"time"
// 	// "os"
// )

// const ACK = 200
// const REJECT = 0

// // Faulty coordinator
// func faulty_node(id int, node_chans []chan int) {

// 	// Working initially
// 	for i := 0; i < 3; i++ {
// 		sender_id := <-node_chans[4]
// 		fmt.Println("Node 5 received message from Node", sender_id)
// 		node_chans[sender_id-1] <- ACK
// 	}
// 	// Failed later. Stop replying to messages.
// 	fmt.Println("(Break down) Coordinator Node 5 broke down")
// }

// func election_call_handler(caller_id int, election_chan chan int, caller_chan chan int, coordinator_id int) int {
// 	for {
// 		select {
// 		case res := <-caller_chan:
// 			if res == REJECT { // rejection
// 				// Stop election process
// 				fmt.Println("Node", caller_id, "rejected as coordinator.")
// 				return coordinator_id
// 			}
// 		case <-time.After(1 * time.Second): // accepted
// 			// Set self as coordinator and broadcast
// 			fmt.Println("Node", caller_id, "elected as coordinator. Broadcasting...")
// 			coordinator_id = caller_id
// 			for i := 0; i < 3; i++ {
// 				election_chan <- caller_id + 5
// 			}
// 			return coordinator_id
// 		}
// 	}
// }

// func normal_node(id int, election_chan chan int, node_chans []chan int) {
// 	coordinator_id := 5
// 	count := 1
// 	// Simulate msg exchange with faulty coordinator
// 	if id == 1 { // set id == 1 for Worst case, id == 4 for Best case
		
// 		// send messages to faulty coordinator
// 		go func() { for { node_chans[coordinator_id-1] <- id } }()
		
// 		// receive ACK messages 
// 		loop := true
// 		for loop {
// 			count++
// 			fmt.Println("Node", id, "sent message to Node", coordinator_id)
// 			select {
// 			case <-node_chans[id-1]:
// 				fmt.Println("Node", id, "received ACK from Node", coordinator_id)
// 			case <-time.After(1 * time.Second):
// 				// Call for election
// 				fmt.Println("(Timeout) No response from Node", coordinator_id, "\nNode", id, "Calling election...")
// 				election_chan <- id
// 				// Listen for election responses
// 				// old_coordinator_id := coordinator_id
// 				coordinator_id = election_call_handler(id, election_chan, node_chans[id-1], coordinator_id)
// 				// clear old coordinator channel
// 			}
// 		}

// 	}

// 	// Simulate all other nodes listening to election
// 	for {
// 		caller_id := <-election_chan
// 		fmt.Println("caller_id", caller_id, "id", id)
// 		if caller_id >= 6 { // new coordinator found
// 			coordinator_id = caller_id-5
// 			fmt.Println("(New Coordinator) Node", id, "is notified that Node", coordinator_id, "is now the new coordinator")
// 			time.Sleep(1*time.Millisecond)
// 			// os.Exit(0)
// 		} else if caller_id < id && caller_id != id { // if id is smaller, reject
// 			node_chans[caller_id-1] <- REJECT
// 			// call for election
// 			fmt.Println("(Rejection) Node", id, "rejected Node", caller_id)
// 			fmt.Println("(New Election) Node", id, "calling for election...")
// 			election_chan <- id
// 			coordinator_id = election_call_handler(id, election_chan, node_chans[id-1], coordinator_id)
// 		}
// 	}

// }

// func main() {

// 	// create 1 election channel and 5 node-specific channels
// 	election_chan := make(chan int) // broadcast id of election caller to call for election
// 	chan1 := make(chan int)         // to reject election caller
// 	chan2 := make(chan int)
// 	chan3 := make(chan int)
// 	chan4 := make(chan int)
// 	chan5 := make(chan int)

// 	node_chans := []chan int{chan1, chan2, chan3, chan4, chan5}

// 	// create 1 faulty coordinator node and 4 normal nodes
// 	go faulty_node(5, node_chans)
// 	go normal_node(1, election_chan, node_chans)
// 	go normal_node(2, election_chan, node_chans)
// 	go normal_node(3, election_chan, node_chans)
// 	go normal_node(4, election_chan, node_chans)

// 	var input string
// 	fmt.Scanln(&input)
// }
