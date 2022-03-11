package main

import (
	"fmt"
	"os"
	"strings"
)

var books = [...]string{"Clean code", "The C programming language", "Effective Java", "C language", "C++ Prime"}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Printf("\nYou have to enter valid commands after .go \n")
		fmt.Printf("Valid commands are: \n search => search a single book \n list => list all books \n")
		return
	}
	wantedBook := strings.Join(args[2:], " ") //joins strings after  to a single string
	if args[1] == "search" {                  //Searches command search a single book and prints its index
		for i, book := range books {
			if book == wantedBook {
				fmt.Printf("\n%s is in index : %d\n ", book, i)
				return
			}
		}
		fmt.Printf("%s is not found", wantedBook)
	} else if args[1] == "list" { //Lists all books
		for _, book := range books {
			fmt.Printf("\n%s", book)
		}
	} else {
		fmt.Printf("\nValid commands are: \n search => search a single book \n list => list all books \n")
	}
}
