# Basic Abstractions
 Abstractions are models we use to set basis of a distributed system and to explain how it performs computations and communications. Two kinds of abstractions are of primary importance: **processes** and **links**(messaging). 

## Computation
![process](https://i.imgur.com/cfzBuNZ.png)  
 Computation through a process follows three steps: **receiving**, **executing** and **sending**. It is considered deterministic, which its steps and output are uniquely determined by its input i.e. received message.
 1. Process receives a message from another process.
 2. Modules execute a local computation
 3. Process then sends a message to some process.
 
 In computation, we seek to garantee two important properties of a system: **safety** and **liveness**. Understanding those properties will help to devise an adequate algorithm for implementation.
 - **Safety**: "Nothing bad will happen". 
   - If safety is violated, will never be satistfied again.
   - Messages should not be invented nor duplicated.
   - Unexpected behavior/output should not happen. 
 - **Liveness**: "Something good will happen".
   - Messages should be delivered and not lost.
   - Process should respond to messages.

## Process
