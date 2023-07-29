package main


import "fmt"

func main (){

	fmt.Println("Hello go")

	sayHello("omid")

}

func sayHello(name string){
	fmt.Println("hello", name)
}