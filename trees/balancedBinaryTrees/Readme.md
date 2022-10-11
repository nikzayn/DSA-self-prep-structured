## Invert Tree

- In this problem, we need to invert the left node to the right nodes in ancestral and child way.

### Example

Input:

```
     3
   /   \
  9     20
        / \
       15  7
```

Output:

```
true
```

### Approach

1. Bottom-up approach

- In this approach, we are recursively checking the balancing via depth and isBalanced function, in depth function, we are

2. Top-down approacj

- Meanwhile, using recursive DFS + swapping approach is not scalable, if number of node height or depth increases. Stacks and Queues would be the good data structures for this too support this edge case.

### Time and space complexity

1.  Bottom-up approach

- Time complexity - O(n) - no. of nodes in a tree
- Space complexity - O(h) - extra stack calls due to height of the tree

1.  Top-down approach

- Time complexity - O(n) - no. of nodes in a tree
- Space complexity - O(h) - extra stack calls due to height of the tree
