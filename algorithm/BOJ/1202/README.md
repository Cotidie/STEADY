## Link
[보석 도둑](https://www.acmicpc.net/problem/1202)

## Topic
- Greedy
- Priority Queue

## Approach
 가방에 넣는 문제여서 냅색 및 DP로 착각할 수 있는 문제이다. 이 문제는 정해진 최대 개수만큼 보석을 나열하는 문제이므로 정렬로 해결할 수 있다. 기존의 DP문제는 같은 가방에서 이전 단계의 최대값을 활용하지만 이 문제는 가방이 다르기 때문에 이전 단계의 최대값을 활용할 수 없다. 새로 넣기로 결정했다면, 기존에 있던 보석을 어느 가방에 넣어야할지 다음 단계에서 **다시 결정**해야 한다. 그러므로 DP가 성립하지 않는다.

 1. 보석을 하나씩 탐색하며 가방에 들어갈 수 있는지 검사한다.
 2. 어떤 한 가방에 들어갈 수 있으면 후보군에 넣는다.
 3. 모두 탐색한 후 후보군을 정렬한다.
 4. 정렬한 후보군에서 차례차례 꺼내어 가방에 배정한다.
    - n개를 순서대로 꺼낼 수 있는가?


 ## Note
- 냅색 문제에 우선순위큐 방식의 접근이 먹히지 않는 이유?
