// LESSON 3 : Time Complexity
// Problem 1 : FrogJmp

/*****
Count the minimal number of jumps that the small frog must perform to reach its target.

Write a function:
func Solution(X int, Y int, D int) int

that, given three integers X, Y and D, returns the minimal number of jumps from position X to a position equal to or greater than Y.

For example, given:

  X = 10
  Y = 85
  D = 30
the function should return 3, because the frog will be positioned as follows:

after the first jump, at position 10 + 30 = 40
after the second jump, at position 10 + 30 + 30 = 70
after the third jump, at position 10 + 30 + 30 + 30 = 100
Write an efficient algorithm for the following assumptions:

X, Y and D are integers within the range [1..1,000,000,000];
X ≤ Y.
**/

package solution

import (
    "math"
)

func Solution(X int, Y int, D int) int {
    //answer is simple (Y-X)/D
    return int(math.Ceil(float64(Y-X)/float64(D)))
}
