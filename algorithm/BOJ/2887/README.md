## Link
[행성 터널](https://www.acmicpc.net/problem/2887)

## Topic
- MST (Kruskal)
- Priority Queue

## Approach
행성을 잇는 모든 간선에 대해 최소 스패닝 트리를 구해야 하나, 행성의 개수가 10만개가 될 수 있으므로 간선을 미리 구할 수는 없다. 점들을 X, Y, Z 기준으로 각각 정렬한 후 가장 차이가 작은 두 행성부터 길을 잇는다.

1. X, Y, Z 좌표를 저장할 3차원 벡터를 준비한다.
    - `planets[0][i]`: i번째 행성의 x 좌표
    - `planets[1][i]`: i번째 행성의 y 좌표
    - `planets[2][i]`: i번째 행성의 z 좌표
2. planets[0:2]를 각각 오름차 순으로 정렬한다.
3. 정렬한 x, y, z 좌표에 대해 인접한 두 행성 간 터널 비용을 우선순위 큐에 넣는다.
    - x, y, z 좌표 각각 넣는다. (5개 행성을 때 넣는 횟수는 3 * 4)
    - 우선순위 큐에 넣을 자료형 `Edge{planetA, planetB, cost}`를 정의한다.
4. 우선순위 큐에서 하나씩 꺼내 두 행성을 연결한다.
    - Disjoint Set을 이용한다.
    - 사이클이 발생하면 이미 경로가 존재하므로 연결하지 않는다. 
5. N-1개를 연결한 시점에 그만둔다.