## Link
[별 찍기 11](https://www.acmicpc.net/problem/2448)

## Topic
- Recursion

## Approach

`3 x 2^k` (3, 6, 12...) 꼴로 이전 단계의 별 패턴을 바로 아래쪽 양 쪽에 붙여나간다. 최초 3줄의 별찍기를 준비하고, k가 증가할 때마다 이전 단계를 양쪽에 붙여 나간다. 이때, 붙인 만큼 위쪽의 공백을 새로 추가해야 한다.

1. 최초의 3 줄의 별을 준비한다.
2. 입력값 `3x2^k`에서 k번 재귀를 반복하며 다음을 수행한다.
    - 이전 별찍기의 양 옆에 `3x2^(k-2)`개의 공백을 추가한다.
    - 이전 별찍기 2개를 가로로 붙인 후, 기존 별찍기 아래쪽에 추가한다.
3. 재귀 이후 출력한다.
   