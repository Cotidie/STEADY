# Topological Sort

## Overview
![topological sort](.images/topological%20sort.png)  
 Topological sort is a technique that finds a dependency path(a.k.a topological ordering) of the given graph. It only works with Directed Acyclic Graph(DAG) which doesn't have **any cycle** in a **directed** graph. Topological sort is widely used in dependency searching and scheduling.
  - To detect a cycle, use **DFS** on a directed graph or **Union-Find** on a undirected graph
  - Topological sort is a exhaustive search by nature, which costs `O(V+E)` complexity

## Algorithm
A reverse order of DFS from any given node is itself a topological sort.
1. Pick a random node from the given graph
    - Preferably a node of zero indegree
2. Search paths using DFS
    - can use a stack or recursion
3. When returning from DFS, put searched nodes to a result array in reverse order
    - mark the nodes as visited
4. After emptying a stack, repeat until all nodes are visited
    - pick a node from unvisited list
5. Finally the resulting array is a topological sort

```cpp
L ← Empty list that will contain the sorted nodes
while exists nodes without a permanent mark do
    select an unmarked node n
    visit(n)

function visit(node n)
    if n has a permanent mark then
        return
    if n has a temporary mark then
        stop   (graph has at least one cycle)

    mark n with a temporary mark

    for each node m with an edge from n to m do
        visit(m)

    remove temporary mark from n
    mark n with a permanent mark
    add n to head of L
```

## Implementation