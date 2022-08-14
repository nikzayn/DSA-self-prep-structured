## Merge two sorted linked lists

- In this problem, we need to merge two sorted linked lists

### Example

Input:

```
list1 = 1 -> 2 -> 4
list2 = 1 -> 3 -> 4
```

Output:

```
list = 1 -> 1 -> 2 -> 3 -> 4 -> 4
```

### Approach

1. Recursive

- First check with null checks for both list1 and list2
- Compare, the list1 with list2, check if list1 value is less than or equal to list2 value
- If above statement is true, then update the list1 next to recursive call of mergeTwoList(list1.next, list2)
- vice versa with list2 and then return both list according to the logic.

2. Iterative

- First initialize a dummy variable with linked list data structure and also initialize a new variable named as pre and assign it with dummy.
- Then, use while loop to iterate over list1 and list2 to check it they are null
- and then check if list1 value is less than list2, if it is then swap pre next and list1 -> list1, list1 next
- inversely proportional, if list1 is not less than list2, then swap pre next and list2 -> list2, list2 next
- after that update pre to pre next
- outside of while loop, do check if list1 is null or not, if it is then update pre next -> list1 else pre next -> list2
- return dummy next

### Time and Space complexity

1. Recursive

- Time complexity - O(n + m) - At every call of recursion, we are adding one node to the output.
- Space complexity - O(n + m) - Stack space will be used in recursion.

2. Iterative

- Time complexity - O(n + m) - where N and M are the sizes of linked lists.
- Space complexity - O(1)
