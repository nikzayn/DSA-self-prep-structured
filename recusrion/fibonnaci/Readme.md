## Nth Fibonacci

- In this problem, we need to return the nth fibonacci number.

### Example - 1

Input:

```
n = 2
```

Output:

```
1
```

### Example - 2

Input:

```
n = 6
```

Output:

```
5
```

### Approach

1. Brute force approach

- First, we check if number of sequence is equal to 1, return 0
- If number of sequence is equal to 2, return 1
- Then, do a recursive check fib(n - 2) + fib(n - 1) and return it.

2. HashMap approach

- Initialize a map data structure and check if value is present in map, if it is then return value
- If value is not present, then update the map value with helper(n - 2) + helper(n - 1)

### Time and Space complexity

1. Brute force approach

- Time complexity - O(2\*n) - It will exceed the call stack as the number of recursive calls
- Space complexity - O(n) - due to n number of recursive calls

2. HashMap approach

- Time complexity - O(n) - to iterate over number of fibonacci calls
- Space complexity - O(n) - to store the value in map data structure
