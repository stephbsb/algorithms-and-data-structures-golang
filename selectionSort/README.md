## Selection Sort

In the algorithm of selection sort we repeatedly select the minimum( or maximum) element of an unordered list and move it to a ordered position of the list.

Time complexity - worst scenario: O(n²)

Used Data Structure: Array

### Example

considering the following array as the array we want to order:

| 64 | 25 | 12 | 22 | 11 |

- In the first pass the first element will be compared to every element until the end of the array to verify if there is an smaller element. If there is a smaller element then they are swapped. Now the first element is the smalled and is considered ordered.

resulting array after first pass:

| 11 | 25 | 12 | 22 | 64 |

- On the second pass we take the second element and compare it to all the following elements of the array, saving the index of the minimum element. The first elemente is already ordered so we dont compare with it. At the end of the pass we will have the first and second element of our ordered list. 

| 11 | 12 | 25 | 22 | 64 |

- We follow the same steps until the last element, then we will get our ordered array

| 11 | 12 | 22 | 25 | 64 |

Since we have **n** elements and to every loop we have to compare with **n-1** elements we have a time complexity of **O(n²)**


To see the algorithm: [selection.go](selectionSort.go)

### How to execute:

inside selectionSort folder run:

````
go run selectionSort.go
````