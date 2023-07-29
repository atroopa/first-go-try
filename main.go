package main


import (
	"fmt"
)

func main (){

	// variables
	var annualSalary string = "golnaz"
	sayHello(annualSalary)
}


// #################### Function ##################
func sayHello(name string){
	fmt.Println("hello", name)
}