package main

//https://app.codility.com/programmers/lessons/4-counting_elements/missing_integer/

// you can also use imports, for example
import "fmt"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) (int){

	var answer int 
	arrayTable:= make(map[int]bool)
	
	
	for _, entry := range A{
		
		arrayTable[entry] = true
		
	}
	
	for i:=1; i<=1000000; i++{
		if(!arrayTable[i]){
			answer = i
			break
		}
	}
	
	return answer
}

func main() {
	A := []int{1, 3, 6, 4, 2}
	fmt.Println(Solution(A))
}