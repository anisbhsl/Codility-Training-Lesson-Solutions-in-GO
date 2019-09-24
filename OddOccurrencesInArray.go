package main

//https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/

// you can also use imports, for example:
import "fmt"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) (int){
	var answer int
    keys := make(map[int]int)
	
    for _, entry := range A {
        keys[entry] +=1
          
    }    
	
	for key,val:=range keys{
		if val%2==1{
		 answer = key
		 break
	}
}
	return answer
}

func main() {
	A := []int{9,3,9,3,9,7,9}
	fmt.Println(Solution(A))
}
