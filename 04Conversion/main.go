package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Transaction App")
	fmt.Println("Please Rate our Transaction App b/w 1 to 5, 5 is outstanding")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') //End when some one press next line
	fmt.Println("Thank you for giving the rating", input)
	//Now we want to add 1 in given rating for this we need to do the conversion
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // here if we will not trim the input it add extra trailing char is the input
	if err != nil {
		fmt.Println(err)
		//panic() // this we can youse to end the program
	} else {
		fmt.Println("Adding 1 to provided rating", numRating+1)
	}
}
