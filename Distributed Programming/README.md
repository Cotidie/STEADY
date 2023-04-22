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
![Failures](https://i.imgur.com/ISw7LuG.png)  
 Abstraction on types of process failures sets an important basis on implementation. If it's likely to crash for some processes, the system wouldn't want to try to recover them and simply do best with the processes alive. 

- **Crash**: The process never recovers after a crash.
  - if it crashes during the execution, the process is said to be *faulty*
- **Omission**: A fault where a process doesn't send a message.
  - can be caused by buffer overflow or network congestion
- **Crashes with Recoveries**: The process tries to recover from a crash
  - Central controller initiates `Recovery` procedure, which is atomic with `Init` procedure.
  - Recovered process is considered as the same process, otherwise it's assumed as the number of processes increased.
- **Eavesdropping Faults**: An unauthorized process trying to act like a normal process.
- **Arbitrary Faults**: When no clear behavior of fault is defined and can deceive other processes. (a.k.a. Byzantine)
  - Byzantine comes from ['Byzantine Generals Problem`](https://river.com/learn/what-is-the-byzantine-generals-problem/).
  - Arbitrary faults refer to intentional and malicious behaviors of a process.