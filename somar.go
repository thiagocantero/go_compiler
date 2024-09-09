package main

import "fmt" 

func somar(x int, y int){ 
	resultado := x + y 
	fmt.Println("A soma de x + y é: ", resultado) 
} 

func main(){ 
	somar(1, 1) 
}