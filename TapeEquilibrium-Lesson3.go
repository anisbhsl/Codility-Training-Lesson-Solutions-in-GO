//LESSON 3 : TIME COMPLEXITY
//TapeEquilibrium


/**
A non-empty array A consisting of N integers is given. Array A represents numbers on a tape.

Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].

The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|

In other words, it is the absolute difference between the sum of the first part and the sum of the second part.

For example, consider array A such that:

  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3
We can split this tape in four places:

P = 1, difference = |3 − 10| = 7 
P = 2, difference = |4 − 9| = 5 
P = 3, difference = |6 − 7| = 1 
P = 4, difference = |10 − 3| = 7 
Write a function:

class Solution { public int solution(int[] A); }
that, given a non-empty array A of N integers, returns the minimal difference that can be achieved.

For example, given:

  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3
the function should return 1, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [2..100,000];
each element of array A is an integer within the range [−1,000..1,000].

**/
package solution

import (
   // "fmt"
    "math"
    )

func getSum(M []int) int{
    sum:=0
    for _,val :=range M {
        sum+=val
    }
    return sum //return the sum
    
} 

func Solution(A []int) int {
    
    var minDiff int
    length:=len(A)  //get length of A
   
    //if array is of size 1 or 2 return 1
    if length==1 || length==2 {
        return 1
    }else{
        
        for p:=1;p<=length;p++{
          leftSum:=getSum(A[:p])
          rightSum:=getSum(A[p:])
          diff:=int(math.Abs(float64(leftSum-rightSum)))
          if p==1{  //run only once to initialize minDiff
              minDiff=diff
          }
         // fmt.Println(diff)
          if diff<minDiff{  //check for minimum difference
              minDiff=diff
          }
        }
    }
    
    return minDiff
}

// func main() {
//     A:=[]{3,1,2,4,3}
//     fmt.Println(Solution(A))
// }
