## Invert Tree

- In this problem, we need to invert the left node to the right nodes in ancestral and child way.

Example:

Input:

```
     4
   /   \
  2     7
 / \   / \
1   3 6   9
```

Output:

```
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

- As, we can see the leftmost and rightmost leaf nodes swapped, in the case of go, we can do something like this: `a, b = b, a`,
  it's a short technique to swap the values.

- Meanwhile, using recursive DFS + swapping approach is not scalable, if number of node height or depth increases. Stacks and Queues would be the good data structures for this too support this edge case.

- Time complexity - O(n) - no. of nodes in a tree
- Space complexity - O(h) - extra stack calls due to height of the tree
