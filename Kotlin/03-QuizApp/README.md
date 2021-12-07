# Quiz App

# 배운점
## UI
* **res/.../themes.xml**: 액션바/상태표시줄 표시, 앱컬러 등을 설정한다.
```xml
<style name="Theme.QuizApp" parent="Theme.MaterialComponents.DayNight.NoActionBar">
    <!-- Primary brand color. -->
    <item name="colorPrimary">@color/purple_500</item>
    ...
    <item name="android:windowFullscreen">true</item>
</style>
```
* **ProgressBar**: 원형, 일자형 등 다양한 스타일로 진행상황을 표시한다.
```xml
<ProgressBar
    android:layout_width="0dp"
    android:layout_height="wrap_content"
    android:minHeight="50dp"
    android:layout_weight="1"
    style="@style/Widget.AppCompat.ProgressBar.Horizontal"
    // max, progress: 최대값, 현재값
    // indeterminate: true이면 단순 애니메이션으로 동작 (네트워크 딜레이 등과 같이 완료시점이 모호한 경우)
    android:max="9"
    android:indeterminate="false"
    android:progress="5"
    />
```
## Android
* **View Binding**: View 요소를 인스턴스로 매핑하는 방법. findById를 대체하며, onCreate에서 바인딩을 진행한다.
```kotlin
override fun onCreate(savedInstanceState: Bundle?) {
    // Inflator는 현재 Window 전용 Inflator이다
    // *Window: 안드로이드 UI에서 구분된 구역(contentView 등)
    val viewBinding = ActivityMainBinding.inflate(layoutInflater)
    val btnStart = viewBinding.btnStart
    ...
    // 바인딩된 View를 전달해야 한다.
    setContentView(viewBinding.root)
}
```
* **Intent**: Activity 간 이동, 통신을 위한 클래스
```kotlin
// 현재 -> QuizApp Activity
val intent = Intent(this, QuizQuestionsActivity::class.java)
startActivity(intent)
// 현재 Activity 종료
finish()
```


# 이슈
* **Android SDK Downgrade**
    * 문제: 기기 SDK 버전과 Android Studio SDK 버전이 일치하지 않아 테스트 불가능하던 문제
    * 해결: build.gradle에서 sdk 버전을 수정한다.
    ```gradle
    android {
        compileSdk 31

        defaultConfig {
            ...
            minSdk 25
            targetSdk 31
            ...
        }
    }
    ```
* **List Slicing**
    * 문제: slice(listOf(1, 3))은 개별 인덱스 1, 3번 두 개만 slice한다.
    * 해결: 범위로 지정하려면 1..3(=[1,3])을 사용한다.