## Bubble Sort

In the algorithm of bubble sort we repeatedly compare adjecent elements and swap them when they are unnordered.

Time complexity - worst scenario: O(n²)

Used Data Structure: Array

### Example

considering the following array as the array we want to order [n=5]:

| 64 | 25 | 12 | 22 | 11 |

- In the first external loop we will swap the first element until it reaches its ordered position:

We will use i to external loop counter and j as the internal loop counter:

**1º pass** [i=0 and 0 < j < n-1]:
````
64 <-> 25 12 22 11
25 64 <-> 12 22 11
25 12 64 <-> 22 11
25 12 22 64 <-> 11
25 12 22 11     64 [64 is ordered]
````

**2º pass** [i=1 and 0 < j < n-2]:
````
25 <-> 12 22 11 64
12 25 <-> 22 11 64
12 22 25 <-> 11 64
12 22 11     25 64 [25 and 64 are ordered]
````

**3º pass** [i=2 and 0 < j < n-3]:
````
12 <-> 22 11 25 64
12 22 <-> 11 25 64
12 11     22 25 64 [22, 23 and 64 are ordered]
````

**4º pass** [i=3 and 0 < j < n-4]:
````
12 <-> 11 22 25 64
11 12 22 25 64 [ordered array]
````

Since we have **n-1** external loops and to every loop we have to compare with **n-i-1** elements we have a time complexity of **O(n²)**


To see the algorithm: [bubbleSort.go](bubbleSort.go)

### How to execute:

inside bubbleSort folder run:

````
go run bubbleSort.go
````