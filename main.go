package main


import (
	"fmt"
	"time"

	Jalaali "github.com/jalaali/go-jalaali"
)

func main (){

	// variables
	var annualSalary string
	var bigNumber int

	annualSalary = "golnaz"
	sayHello(annualSalary)

	bigNumber = 12
	fmt.Println(bigNumber)

	fmt.Println(Jalaali.ToJalaali(time.Now().Date()))

}


// #################### Function ##################
func sayHello(name string){
	fmt.Println("hello", name)
}