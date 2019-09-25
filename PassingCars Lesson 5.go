package main

import (
	"fmt"
)

func Solution(A []int) int{
	n:= len(A)
	carPass := 0
	goingLeft := 0
	for i:= n-1 ; i>=0 ; i--{
		
		if(A[i] == 1){
			goingLeft+=1		
		}else{
			carPass+=goingLeft
			
			if carPass > 1000000000{
				return -1
			}
		}
		//fmt.Println(carPass)
	
	}
	
	return carPass
}


func main() {
	A:=[]int{0,1,0,1,1}
	fmt.Println(Solution(A))
}
