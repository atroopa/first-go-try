package main


import "fmt"

func main (){

	var annualSalary string
	var bigNumber int

	annualSalary = "golnaz"
	sayHello(annualSalary)

	bigNumber = 12
	fmt.Println(bigNumber)

}

func sayHello(name string){
	fmt.Println("hello", name)
}