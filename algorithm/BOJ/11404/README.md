## Link
[플로이드](https://www.acmicpc.net/problem/11404)

## Topic
- Shortest Path (Floyd-Warshall)

## Approach
플로이드 알고리즘은 새로운 경유지를 추가했을 때, 그것이 기존보다 더 적은 비용인지를 비교하는 알고리즘이다. 2차원 테이블에 모든 노드 사이의 최단거리를 구한다.

1. 도시의 개수에 맞추어 n^2 테이블을 준비한다.
    - `cities`: 두 도시 간 최단거리를 기록할 테이블
    - 초기값은 음수로 초기화한다.
    - 자기 자신에 대한 값은 0으로 초기화한다.
2. 임의의 경유지에 대하여 기존 비용보다 더 작은지 검사하고 덮어쓴다.
    - `cities[a][k] + cities[k][b] < cities[a][b]`
3. `cities` 배열을 출력한다.

## Trials
- 코드 작성 후 floyd()를 돌리지 않고 그대로 출력했다.
- 기존 경로와 갱신 경로의 대소관계를 거꾸로 했다.

