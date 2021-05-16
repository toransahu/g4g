# Problems

<!-- ## 5. tbd -->

<!-- #### Input -->

<!-- ``` -->
<!-- ``` -->

<!-- #### Output -->

<!-- ``` -->
<!-- ``` -->

## 1. Find the maximum sum of a contiguous sub array, less than or equal to K.

#### Input

```
arr: 1,3,5
k: 6
```

#### Output

```
contiguous subarray: [1,3]
max sum: 4
```

#### Input

```
arr: 1,2,5,7
k: 10
```

#### Output

```
contiguous subarray: [1,2,5]
max sum: 8
```

## 1.1. check if there exists a non-contiguous subarray of K elements in an array of length N whose sum is equal to a given sum.

Ref: https://stackoverflow.com/questions/64125807/check-if-there-exists-a-non-contiguous-subarray-of-k-elements-in-an-array-of-n-e

#### Input

```
Array = [1,2,3,4,5,6,7,8,9]
length of Subarray(K) = 3
target sum = 7.
```

#### Output
```
A non-contiguous subarray of length 3 with sum=7 is [1,2,4].
```

#### Constraint
```
1<=n<=10^6  
1<=k<=100
```

## 1.2. Check if a non-contiguous subsequence same as the given subarray exists or not

Ref: https://www.geeksforgeeks.org/check-if-a-non-contiguous-subsequence-same-as-the-given-subarray-exists-or-not/
 
## 1.3. Find the sum of a non-contiguous sub array, less than or equal to K and subarray with max length.

#### Input

```
arr: 1,3,5
k: 6
```

#### Output

```
(non) contiguous subarray: [1,5]
max sum: 6
```

#### Input

```
arr: 1,3,4,7
k: 10
```

#### Output

```
(non) contiguous subarray: [1,3,4]
max sum: 8

Notice: subarray [3,7] whose sum=10 > 8 is not the answer.
```


## 1.4. Find the maximum sum of a non-contiguous sub array, less than or equal to K.

#### Input

```
arr: 1,3,5
k: 6
```

#### Output

```
(non) contiguous subarray: [1,5]
max sum: 6
```

#### Input

```
arr: 1,3,4,7
k: 10
```

#### Output

```
(non) contiguous subarray: [3,7]
max sum: 10
```


## 1.5. Find the sum of a contiguous sub array, less than K, and subarray with max length.

## 1.6. Given an array of positive integers, what's the most efficient algorithm to find non-consecutive elements from this array which, when added together, produce the maximum sum?

Ref: https://stackoverflow.com/questions/4487438/maximum-sum-of-non-consecutive-elements

## 1.7. Given an array of positive integers, find the sum of all its subarrays in O(n) and in constant space?
Ref: 
- https://www.quora.com/Why-are-sub-array-problems-so-hard-to-solve-What-are-some-tricks-to-solve-it-in-O-n
- https://web.stanford.edu/class/cs9/sample_probs/SubarraySums.pdf

## 1.8. Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Ref: https://leetcode.com/problems/maximum-subarray/

#### Input
```
nums = [-2,1,-3,4,-1,2,1,-5,4]
```

#### Output

6

#### Explanation

```
[4,-1,2,1] has the largest sum = 6.
```


#### Input
```
nums = [1]
```

#### Output
1

#### Input
```
nums = [5,4,-1,7,8]
```

#### Output
23

## 1.9. Given an integer array nums, find the contiguous subarray which has the largest sum and return its sum.

Ref: https://en.wikipedia.org/wiki/Maximum_subarray_problem

#### Input
```
nums = [-2, 1, -1, 3, -2, 4, -5]
```

#### Output

5

#### Input
```
nums = [-2, -1, -1, -3, -2, -4, -5]
```

#### Output

0

## 2. Solve the Logical Expression given by string

#### Input

```
str: “[!,[[0,&,[!,1]],|,[!,[[!,0],&,1]]]]”
```

#### Output

```
1
```

#### Input

```
str = “[[0,&,1],|,[!,1]]”
```

#### Output

```
0
```

## 3. Given an expression string expression of parenthesis ( assume numbers from 0 to 9 each represent a type of parenthesis), write a program to examine whether the pairs and the orders are valid in expressions.

NOTE: Negative numbers are opening and positive are closing

#### Input

```
'-9-999'
```

#### Output

```
VALID
```

#### Input

```
'-1-1-22-4411'
```

#### Output

```
VALID
```

## 4. Given N tasks, and the dependencies among them, please provide an execution sequence, which make sure jobs are executed without violating the dependency.

#### Input

```
5
1<4
3<2
4<5

First line is the number of total tasks. 1<4 means Task 1 has to be executed before task 4.
```

#### Output

```
One possible sequence would be: 1 4 5 3 2
```

## 5. In a party of N people, only one person is known to everyone. Such a person may be present in the party, if yes, (s)he doesn’t know anyone in the party. We can only ask questions like “does A know B? “. Find the stranger (celebrity) in the minimum number of questions.

#### Input

```
MATRIX = { {0, 0, 1, 0},
           {0, 0, 1, 0},
           {0, 0, 0, 0},
           {0, 0, 1, 0} }
```

#### Output

id = 2

#### Explanation

The person with ID 2 does not know anyone but everyone knows him.

#### Input
```
MATRIX = { {0, 0, 1, 0},
           {0, 0, 1, 0},
           {0, 1, 0, 0},
           {0, 0, 1, 0} }
```

#### Output
No celebrity

#### Explanation
There is no celebrity.

## 6. Merge K sorted linkedlists.
#### Constraints
Auxilary Space `<= O(K)`

## 7. Consider a field(assume a M*N matrix), you have to find the shortest path from any cell in the first column to any cell in the last column.


Condition:

There are trespasser sensor in the field, which are marked with value 0. The cell with value 0 and also it's 8 adjacent cells can activate sensors. You have to find the shortest safe path.

Possible moves:

```
Go Up: (x, y) ——> (x – 1, y)
Go Left: (x, y) ——> (x, y – 1)
Go Down: (x, y) ——> (x + 1, y)
Go Right: (x, y) ——> (x, y + 1)
```

#### Input

```
{ 0, 1, 1, 1, 0, 1, 1, 1, 1, 1 },
{ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1 },
{ 1, 1, 1, 1, 1, 1, 1, 1, 0, 1 },
{ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1 },
{ 1, 1, 1, 1, 1, 0, 1, 1, 1, 1 },
{ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1 },
{ 1, 0, 1, 1, 1, 1, 1, 1, 1, 1 },
{ 1, 1, 1, 1, 1, 1, 1, 1, 1, 0 },
{ 1, 1, 1, 1, 1, 0, 1, 1, 1, 1 },
{ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1 }
```

#### Output

```
shortest path has length: 11
```

## 8. Find if a given linkedlist has a loop?

```
[1] → [2] → [3]
       ↑     ↓
      [5] ← [4]
```

## 8.1. Find if a given singly linkedlist is circular?

```
[1] → [2] → [3]
 ↑           ↓
[6] ← [5] ← [4]
```

## 8.2. Find if a given doubly linkedlist is circular?

## 9. How Dynamic Array / Vector / List / Slice / ArrayList works internally?

## 10. Detect Cycle in a Directed Graph?

## 11. Remove duplicate nodes from a unsorted linked list

## 11.1 Remove duplicate nodes from a unsorted linked list without using extra space

## 12. Suppose you have a input of continuous stream of integers, and at any point you want to get the median of the elements you have received so far. Write an algorithm for that.

## 13. Count the number of ways you can climb a stair, if the number of steps you can take is either 1 step, 2 step or 3 step at a time.

Input: n = 3

Output: 4
- 1 step + 1 step + 1 step
- 1 step + 2 step
- 2 step + 1 step
- 3 step

## 14. Let two numbers represented by two linked lists, write a function that returns the new list. The new  linked list is the representation of the addition of two input numbers. You can not modify the given input linked list and you can not use any explicit extra space other than the new linked list.

## 15. design a data structure that can do following operations in O(1) time
Insert(X), delete(X), search (X) and getRandom()

## 16. Design an efficient data structure that can do following operations:

- findMin() returns minimum item, frequency of operation: mostly
- findMax() return maximum item, frequency : mostly
- deleteMin(): delete Minimum item, frequency: moderate
- deleteMax: delete maximum item, frequency: moderate, 
- insert: insert an item, frequency: least
- delete: delete an item, frequency: least

## 17. Trees can be implemented either using linked list or array. What is the first thing you consider when you choose either of the option for your problem

## 18. Suppose you are getting stream of words like army, Mary, tea, coffee, eat, kit, tik. Then you have to design a data structure to store these values such that if a query is made to search  tik and it should also return it's anagram.

Write the function prototype as well.

## 19. Two nodes of a binary search tree is swapped, how do you find which two value is swapped and how can you restore the original BST
