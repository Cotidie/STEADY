## Link
[팰린드롬 분할](https://www.acmicpc.net/problem/1509)

## Topic
- DP (2차원)
- 백트래킹

## Approach
 
`Text[i:j]`의 팰린드롬 여부를 알아보려면 바로 직전 범위 `Text[i+1, j-1]`의 팰린드롬 여부를 알아보면 빠르게 구할 수 있다. 직전에 팰린드롬이었다면 i와 j가 서로 같은 문자인지만 확인하면 되고, 팰린드롬이 아니었다면 이번 범위 또한 팰린드롬이 아니다. 

1. **Subproblem**: `DP[i][j]`, Text[i:j]의 팰린드롬 여부
2. **Guess**: 직전 범위 `DP[i+1][j-1]`이 팰린드롬이면 i, j만 비교하면 된다.
3. **Recurrence**: `DP[i+1][j-1]`의 팰린드롬 여부에 따라
    - `DP[i][j] = (Text[i] == Text[j])`, 직전에 팰린드롬이었을 때
    - `DP[i][j] = False`, 직전에 팰린드롬이 아니었을 때
4. **Solve**: `DP[i][j]`

DP 배열은 대칭이며, 초기값으로 DP[i][i]와 DP[i][i+1]을 채워넣어야 한다.

DP 이후 두 가지 접근으로 분할의 개수를 확인할 수 있다.

**(1)** 백트래킹 => 시간초과
1. 몇 개로 분할할지 1부터 문자열 개수만큼 하나씩 검사한다.
2. 가장 처음 분할이 가능해지는 개수를 출력한다.

**(2)** DP
DP[n]을 [0:n]까지 부분 문자열의 최대 분할 개수라 하면 i=1..n까지 분할이 가능한지 검사하고 DP를 갱신한다.
- i=0일 때 분할이 가능하다면 이때 최대 분할 개수는 1이다. => 이후 순회는 생략
- `DP[n] = min(DP[n], DP[i]+1)`, [i:n]이 팰린드롬일 경우