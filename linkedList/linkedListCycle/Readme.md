## Linked list cycle

- In this problem, we need to merge two sorted linked lists

### Example

Input 1:

```
[head1](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)
```

Output 1:

```
true
```

Input 2:

```
[head2](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)
```

Output 2:

```
false
```

### Approach

1. Brute force approach - HashMap

- First, we initialize our map data structure
- Then, we check if head is null or not, if not then we range over map and then check if our map contains head val, we return true, otherwise
  we will update map with empty struct and also update head to head.Next
- After that, just return false when head is null

2. Floyd’s Cycle Detection Algorithm

- This approach uses two pointers **fast** and **slow** in which fast nodes moves 2 nodes at a time and slow moves only 1 node at a time.
- Null check for head, it is equal to null, then return false
- Initialize two pointers with fast = head.Next and slow = head
- Iterate over nodes to check if fast == null && fast.Next == null, just return false
- Else do update fast = head.Next.Next, slow = head.Next
- return true at last if all conditions with null checks didn't meet

### Time and Space complexity

1. Recursive

- Time complexity - O(n) - iterate over each node value to check
- Space complexity - O(n) - Using hashmap to store every unmatched value of head

2. Iterative

- Time complexity - O(n) - iterate over each node value to check
- Space complexity - O(1)
