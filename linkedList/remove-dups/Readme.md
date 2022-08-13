## Remove duplicates nodes from sorted linked list

- In this problem, we need to remove the duplicate nodes from the sorted linked list.

### Example

Input:

```
linkedList = 1 -> 1 -> 3 -> 4 -> 4 -> 4 -> 5 -> 6 -> 6
```

Output:

```
linkedList = 1 -> 3 -> 4 -> 5 -> 6
```

### Approach

- Iterate over the linked list and search for duplicate nodes in list and remove them and return the updated list

### Time and Space complexity

- Time complexity - O(n) - because we are searching a particular node
- Space complexity - O(1) - since updating value to the node
