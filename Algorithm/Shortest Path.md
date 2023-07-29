# Shortest Path
- [Dijkstra](#dijkstra)
    - [Algorithm](#algorithm)
    - [Pseudo Code](#pseudo-code)
    - [Time Complexity](#time-complexity)
- [Floyd Warshall](#floyd-warshall)
    - [Algorithm](#algorithm-1)
    - [Pseudo Code](#pseudo-code-1)
    - [Time Complexity](#time-complexity-1)
- [Questions](#questions)


This document explains popular algorithms for finding the shortest path between vertices. 

## Dijkstra
 Given a set of vertices and edges, Dijkstra algorithm finds the shortest paths **from the source** to all vertices. It's a greedy algorithm that finds the shortest subpaths from the source `A` to the destination `D`, meaning that any subpath in between `A -> D` is also the shortest path from `A`.  

### Algorithm
![Algorithm](https://i.imgur.com/73tLrVZ.gif)  
1. Prepare two 1d arrays and a priority queue.
   - **distance**: updates shortest distance
   - **previous**: updates shortest path (previous vertex)
   - **priority queue**: queue of unvisited vertices, sorted by distance
2. Start from the source, examine distances to adjacent vertices.
   - Update distances of the `distance` array
   - Update previous vertex of the `previous` array
3. Pick the shortest one and pop from the priority queue.
   - It's the shortest subpath from the source
4. Repeat until no vertex is left in the priority queue.

### Pseudo Code
The pseudo code below is gotten from [Wikipedia](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm).  
```
 1  function Dijkstra(Graph, source):
 2      
 3      for each vertex v in Graph.Vertices:
 4          dist[v] ← INFINITY
 5          prev[v] ← UNDEFINED
 6          add v to Q
 7      dist[source] ← 0
 8      
 9      while Q is not empty:
10          u ← vertex in Q with min dist[u]
11          remove u from Q
12          
13          for each neighbor v of u still in Q:
14              alt ← dist[u] + Graph.Edges(u, v)
15              if alt < dist[v]:
16                  dist[v] ← alt
17                  prev[v] ← u
18
19      return dist[], prev[]
```

### Time Complexity
- `O(V logE)` if priority queue is used
- `O(V^)` if vector is used


## Floyd Warshall
Floyd Warshall finds the shortest paths **between all the pairs of vertices**. It follows dynamic programming approach for its implementation. 

### Algorithm
1. Prepare a distance table and a path table.
   - path table records previously visited vertex
2. Iterate through all the intermediate vertices `k`
3. Update the tables if the direct distance is greater than the intermediate distance.
   - `dist[a][b] > dist[a][k] + dist[k][b]`

| Graph | Matrix |
|:------:|:------:|
|![Graph](https://cdn.programiz.com/sites/tutorial2program/files/fw-Graph.png) | ![Step 2](https://cdn.programiz.com/sites/tutorial2program/files/fw-Matrix-1.png) |

| Step 1 | Step 2 |
|:------:|:------:|
| ![Step1](https://cdn.programiz.com/sites/tutorial2program/files/fw-Matrix-2.png) | ![Step 2](https://cdn.programiz.com/sites/tutorial2program/files/fw-Matrix-3.png) |

| Step 3 | Step 4 |
|:------:|:------:|
| ![Step 3](https://cdn.programiz.com/sites/tutorial2program/files/fw-Matrix-4.png) | ![Step 4](https://cdn.programiz.com/sites/tutorial2program/files/fw-Matrix-5.png) |

### Pseudo Code
The pseudo code below is gotten from [Wikipedia](https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm#Pseudocode).  
```cpp
let dist be a |V| × |V| array of minimum distances initialized to ∞ (infinity)

for each edge (u, v) do
    dist[u][v] ← w(u, v)  // The weight of the edge (u, v)
for each vertex v do
    dist[v][v] ← 0
for k from 1 to |V|
    for i from 1 to |V|
        for j from 1 to |V|
            if dist[i][j] > dist[i][k] + dist[k][j] 
                dist[i][j] ← dist[i][k] + dist[k][j]
            end if
```

### Time Complexity
`O(V^3)`

## Questions
> Time compexities for Dijkstra and Floyd Warshall are `O(V logE)` and `O(V^3)`. Isn't it better to iterate all the vertices using Dijkstra than using Floyd?
