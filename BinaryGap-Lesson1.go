/*
A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

Write a function:

func Solution(N int) int
that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..2,147,483,647].
*/

package solution

import "math"

func Solution(N int) int {
    //initiliaze variables
    currentGap:=0
    maxGap:=0
    start:=0
    //get number of bits required to represent the number in binary
    bitlength :=int(math.Ceil(math.Log2(float64(N))))
    //start counting for gaps
    for i:=0;i<bitlength;i++{
       
        if N & 1 !=0{  //bitwise AND
            //if true, 1 has been encountered
            //start counting by setting start = 1
            if start==0{
                start=1
            }else if start==1{
                if maxGap<currentGap{
                    maxGap=currentGap
                    currentGap=0 //reset counter
                }else{
                    currentGap=0 //reset counter
                }
            }
        }else{  //for each '0' encounter after '1' has been encountered 
            if start==1{
                currentGap+=1 //increase counter each time
            }
        }
        N=N>>1 //right shift by 1 bit
    }
    return maxGap
}

