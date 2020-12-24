package main

import "fmt"

func main() {
	ui()
}

//function for displaying the base commandline UI
func ui() {
  ui := `
Please choose the structure type of the simple queue by typing the number of the option:
  1. Array
  2. Linked List
  3. Vector
  4. Array List

`
  fmt.Printf("%s", ui)

  var userchoice int
  fmt.Scanln(&userchoice)

  switch userchoice {
  case 1:
    fmt.Println("Simple Queue Using An Array")
    var newsimplequeue []int

		//variable to use for closing out of simple array demonstration
		var rundemoarray int
		rundemoarray = 1

		for rundemoarray == 1 {
    	queueArrayUI :=`
What would you like to do with the Queue?
  1. Enqueue
  2. Dequeue
  3. List the front value in the queue
  4. Display all members of the queue
  5. Exit simple queue with array demo

`
    	fmt.Printf("%s", queueArrayUI)

    	var queueArrayChoice int
    	fmt.Scanln(&queueArrayChoice)

    	switch queueArrayChoice {
    	case 1:
      	fmt.Println("Enqueuing...")
      	newsimplequeue = arrayEnqueue(newsimplequeue)
				fmt.Println("Current Queue Values: ", newsimplequeue)
    	case 2:
      	fmt.Println("Dequeuing...")
				newsimplequeue = arrayDequeue(newsimplequeue)
				fmt.Println("Current Queue Values: ", newsimplequeue)
    	case 3:
      	fmt.Println("The front item in the queue is:", newsimplequeue[0])
    	case 4:
      	fmt.Println("Display all elements")
				fmt.Println("Current Queue Values: ", newsimplequeue)
			case 5:
				fmt.Println("Exiting Simple Queue with Array Demo.")
				rundemoarray = 0
    	default:
    	}
		}
  case 2:
    fmt.Println("Simple Queue Using A Linked List")
  case 3:
    fmt.Println("Simple Queue Using A Vector")
  case 4:
    fmt.Println("Simple Queue Using An Array List")
  default:
    fmt.Println("No Option Selected. Closing....")
  }
}

//functions to demonstrate a simple queue with an array
func arrayEnqueue(a []int) []int {
    a = append(a, 2)
		return a
}

func arrayDequeue(a []int) []int {
	return a[1:]
}

//functions to demostrate a simple queue linked list
