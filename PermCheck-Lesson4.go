package main

//https://app.codility.com/programmers/lessons/4-counting_elements/perm_check/

// you can also use imports, for example
import "fmt"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) (int){
	N := len(A)
	var answer int
	answer = 1
	
	actual := make([]int,N,N)
	
	for i:=0 ; i<N ; i++{
		actual[i] = i+1
	}
	
	actualTable := make(map[int]int)
	
	for _, entry := range actual{
		actualTable[entry]+=1
	}
	
	//fmt.Println(actualTable)
	
	
	for _,val := range A{
		actualTable[val]-=1
	}
	
	
	for _,val := range actualTable{
		if val==1{
			answer = 0
			break
		}
	}
	
	//fmt.Println(actualTable)
	
	
	
	return answer
}

func main() {
	A := []int{4,1,3,2}
	fmt.Println(Solution(A))
}
