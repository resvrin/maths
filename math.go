package main

import (
	"fmt"
	"math/rand"
	"os"
	//"os/signal"
	//"syscall"
	e "github.com/enescakir/emoji"
	"time"
)
type count struct {
	total_count int
}

//Multiplication
func (c *count) multiplication(a, b int) {
	number1 := randomRun(a)
	number2 := randomRun(b)
	c.total_count++
	c.total()
	sum := number1 * number2 
	for {
		fmt.Printf("What would be the product of %v x %v ?: ", number1, number2) 
		var inp int
		fmt.Scanln(&inp)
		if inp == sum {
			fmt.Printf("Very Good!!! \n", e.WavingHand)
			number1 = randomRun(a)
			number2 = randomRun(b)
			c.multiplication(100, 100)
			c.total_count++
		}else{
			for {
			fmt.Printf("Oops That is not right, Try again: ")
			fmt.Scanln(&inp)
			if inp == sum {
				fmt.Println("Good!!!\n")
				number1 = randomRun(a)
				number2 = randomRun(b)
				c.multiplication(100, 100)
				c.total_count++
					}
			}
		}
	}
}



//Addition
func (c *count) addition(a, b int) {
	number1 := randomRun(a)
	number2 := randomRun(b)
	c.total_count++
	c.total()
	sum := number1 + number2 
	for {
		fmt.Printf("What would be the sum of %v + %v ?: ", number1, number2) 
		var inp int
		fmt.Scanln(&inp)
		if inp == sum {
			fmt.Println("Very Good!!! \n")
			number1 = randomRun(a)
			number2 = randomRun(b)
			c.addition(1000, 1000)
			c.total_count++
		}else{
			for {
			fmt.Printf("Oops That is not right, Try again: ")
			fmt.Scanln(&inp)
			if inp == sum {
				fmt.Println("Good!!!\n")
				number1 = randomRun(a)
				number2 = randomRun(b)
				c.addition(1000, 1000)
				c.total_count++
					}
			}
		}
	}
}	
// Subtraction 
func (c *count) subtraction(a, b int) {
	number1 := randomRun(a)
	number2 := randomRun(b)
	if number1 < number2 {
		number1, number2 = number2 , number1
	}
	diff := number1 - number2 
	for {
		fmt.Printf("What would be the diff of %v - %v ?: ", number1, number2) 
		var inp int
		fmt.Scanln(&inp)
		if inp == diff {
			fmt.Println("Very Good!!! \n")
			number1 = randomRun(a)
			number2 = randomRun(b)
			c.subtraction(1000, 1000)

		}else{
			for {
			fmt.Printf("Oops That is not right, Try again: ")
			fmt.Scanln(&inp)
			if inp == diff {
				fmt.Println("Good!!!\n")
				number1 = randomRun(a)
				number2 = randomRun(b)
				c.subtraction(1000, 1000)
				}
			}

		}
	}
}

func (c *count) division(a, b int) {
	number1 := randomRun(a)
	number2 := randomRun(b)
	c.total_count++
	c.total()
	if number1 < number2 {
		number1, number2 = number2 , number1
	}
	div := number1 / number2 
	for {
		fmt.Printf("What would be the d of %v divide %v ?: ", number1, number2) 
		var inp int
		fmt.Scanln(&inp)
		if inp == div {
			fmt.Println("Very Good!!! \n")
			number1 = randomRun(a)
			number2 = randomRun(b)
			c.addition(100.0, 1000.0)
			c.total_count++
		}else{
			for {
			fmt.Printf("Oops That is not right, Try again: ")
			fmt.Scanln(&inp)
			if inp == div {
				fmt.Println("Good!!!\n")
				number1 = randomRun(a)
				number2 = randomRun(b)
				c.division(100.0, 100.0)
				c.total_count++
					}
			}
		}
	}
}

// Common Function to use for random number
func randomRun(a int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(a)
	return random
}

func (c count) total() {
    fmt.Printf("Total question done %v \n", c.total_count)
}


// func (c *count) final() {
// 	channel := make(chan os.Signal)
//     signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
//     go func() {
//         <-channel
//         c.total()
//         os.Exit(1)
//     }()
	
	
// }
func main() {
out := count{}
what := os.Args[1]
	fmt.Println("\nUsage: go run main.go sum or go run main.go mul or go run main.go sub or go run main.go div'\n\n\n")

if what == "sum" {
	out.addition(1000, 1000)
}else if what == "sub" {
	out.subtraction(1000, 1000)
}else if what == "mul" {
	out.multiplication(100, 100)
}else if what == "div" {
	out.division(100.0, 100.0)
}else{
	fmt.Println("Run Again :-) ")
}

}















