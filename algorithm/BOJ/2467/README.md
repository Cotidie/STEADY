## Link
[용액](https://www.acmicpc.net/problem/2467)

## Topic
- Two Pointers

## Approach

정렬된 상태로 들어오므로 한 방향으로 탐색이 가능하다. 가장 끝쪽부터 더하면서 탐색할 때, 왼쪽을 움직인다면 합은 더 커질 것이고, 오른쪽을 움직인다면 합은 더 작아질 것이다. 모든 수를 탐색하고 종료한다. 

1. 두 포인터를 양 끝에 위치시킨다.
2. 합이 0 이상이면 오른쪽을 음수 방향으로, 미만이면 왼쪽을 양수 방향으로 옮긴다.
3. 두 포인터가 겹치는 순간 탐색을 종료한다.

## Note
- 가장 0에 가까운 합을 저장할 때 절댓값을 취해야 한다.