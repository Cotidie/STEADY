## Link
[구슬 탈출 2](https://www.acmicpc.net/problem/13460)

## Topic
- BFS

## Approach
 보드를 기울인다는 규칙으로부터 구슬이 도달할 수 있는 좌표는 한정적이다. BFS로 탐색하면서 다음 경우를 고려해야 한다.

 1. 더이상 탐색이 불가능한 경우 
    - 방문표시 (빨강, 파랑 모두)
    - 11회 이상 탐색했을 때 (출력 조건)
    - 큐가 비었을 때
 2. 파란 구슬이 구멍에 도달하는 경우 (=> 방문 표시에 파란 구슬의 좌표도 반영)
    - 파란 구슬만 도달하는 경우
    - 빨강 구슬과 함께 도달하는 경우
 3. 구멍에 최단 거리로 도달하는 방법 (=> BFS)