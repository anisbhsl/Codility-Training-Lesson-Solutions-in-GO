package main

//https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/

// you can also use imports, for example
import "fmt"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) (int){
	N := len(A)+1
	var answer int
	
	actual := make([]int,N,N)
	
	for i:=0 ; i<N ; i++{
		actual[i] = i+1
	}
	
	actualTable := make(map[int]int)
	
	for _, entry := range actual{
		actualTable[entry]+=1
	}
	
	
	
	for _,val := range A{
		actualTable[val]-=1
	}
	
	//fmt.Println(actualTable)
	
	for key,val:=range actualTable{
		if val==1{
		 answer = key
		 break
		}
	}
	
	
	return answer
}

func main() {
	A := []int{2,3,1,5}
	fmt.Println(Solution(A))
}
