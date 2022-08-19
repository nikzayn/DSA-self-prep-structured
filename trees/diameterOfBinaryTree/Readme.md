## Diameter of Binary Tree

- In this problem, we need to find the diameter of the binary tree

Input:

```
     1
   /   \
  2     7
 / \
3   4
```

Output:

```
3 - length of the path [3, 2, 1, 7] between edges of two nodes
```

### Approach

- First, I will create a TreeInfo struct which consists diameter and height
- Then, I will create a recursive function which will first check the tree empty check, if it is then it will return empty struct of TreeInfo
- Otherwise, we will create variables of left and right info values which recursively get left and right tree height
- Then, I calculate max height with sum of left and right height
- After that, then I will find the maximum diameter from left and right tree
- Then, will initialize a currentDiameter which we will get the max from maxHeight and maxDiameter
- At last, we will add +1 to the find the currentHeight with max of left and right tree
- and then will return the TreeInfo with diameter and height, and we will return the diameter in main function

### Time and space complexity

- Time complexity - O(n) - no. of nodes in a tree
- Space complexity - O(h) - height of a tree
