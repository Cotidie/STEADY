# Dynamic Programming
![DP](https://i.ytimg.com/vi/aPQY__2H3tE/maxresdefault.jpg)  
 Dynamic Programming is a technique for problem-solving by breaking the problem down into smaller sub-problems. The key to DP is to avoid redundant calculations by storing the mid-solutions to the DP table or using memoization. 

# Steps
1. Define Subproblems
   - Get the number of sub-problems
   - Is it suffixed(`x[i:]`), prefixed(`x[:i]`) or sub-fixed(`x[i:j]`)?
2. Guess the part of solution
   - Get the sample solution to a sub problem
3. Find recurrence relation using min, max
   - Define DP equation
4. Ensure topological order
   - Must be acyclic: left to right, smaller to larger, and so on.
5. Solve the original problem

# Examples
## Parenthesization
![Matrix Multiplication](https://bruceoutdoors.files.wordpress.com/2015/11/parenthesized-matrices.png?w=640)  
 Fine the fastest way of calculation to associattive expression A0 * A1 * A2 * ... * An
1. **Subproblems**: optimal evaluation of Ai...A(j-1) => sub-fixed
2. **Guess**: What is the result of the outermost/last multiplication? 
![guess](https://i.imgur.com/iabN50N.png)  
3. **Recurrence**  
![recurrence](https://i.imgur.com/JQDnllT.png)  
4. **Topological Order**: smaller to larger (increasing the subproblem size)
5. **Original Problem**: `DP(0, n)`
