# BFS-DFS
This document explains basics of BFS and DFS, and their applications in problem solving. **BFS**(Breadth First Search) and **DFS**(Depth First Search) are the two most common graph searching algorithms. Queue is typically used for BFS, and Stack/recursion for DFS. 

## BFS
```cpp
BFS (G, s)
  let Q be queue.
  Q.enqueue( s )

  mark s as visited.
  while ( Q is not empty)
    //Removing that vertex from queue,whose neighbour will be visited now
    v  =  Q.dequeue()

    for all neighbours w of v in Graph G
      if w is not visited 
        Q.enqueue( w )             
        mark w as visited.
```
1. Prepare a queue and an array
    - **queue**: stores nodes to visit
    - **array**: marks already visited nodes
2. Put a starting node into the queue
3. Pop from the queue and
    - put adjacent nodes
    - mark adjacent nodes as visited
3. Iterate until the queue is empty

### Level order traversal
![level-order-traversal](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS8kfsogDGt7Xe3g5KyevRUM8uP1-9fWJg9AA&usqp=CAU)  
BFS itself traverses a tree in a level order. In this case the queue should use a data structure that records both a node and its level. 
