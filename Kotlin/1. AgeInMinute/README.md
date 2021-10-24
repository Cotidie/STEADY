# Age In Minute
 생년월일을 선택하면 나이를 분 단위로 표시해주는 앱


## 배운점
 * UI: Constraint Layout 연습, colors.xml, strings.xml 사용법


## 이슈
 * **Device Offline**: 실행 시 Device Offline으로 인하여 앱이 설치되지 않는 문제
    1. \[Tools\] > \[AVD Manager\]에서 장치 삭제 후 새로운 장치 생성
    2. \[Console\] > 'adb kill-server' 및 'adb start-server'
        - 단, 환경변수 PATH에 sdk\\platform-tools 경로가 지정되어 있어야 함. (Tools > SDK Manager에서 경로 확인 가능)
