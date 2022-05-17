# Weather App

## 배운점
### Event Loop
![Event Loop](https://i.imgur.com/zCgCjYs.png)
Node.js는 기본적으로 싱글 스레드이며 병렬적으로 처리되지 않는다. 싱글 스레드 위에서 병렬 처리를 위해 Node.js는 Call Stack과 구분되는 Node API, Event Loop를 둔다. 
* **Node API**: C++ 라이브러리를 활용하는 멀티 스레드 영역
  * 시간을 재거나 I/O 작업 등이 이루어진다.
  * 작업이 끝나면 콜백 함수가 Event Loop로 전달된다. 
* **Event Loop**: Call Stack이 비워지면 도착한 순서대로 콜백 함수를 실행한다.
```js
console.log('Starting')

// Basic Async Function
setTimeout(() => {
    console.log('After 2 seconds')
}, 2000)
setTimeout(() => {
    console.log('After 0 seconds')
}, 0)

console.log('Stopping')

// 실행결과
Starting
Stopping
After 0 seconds
After 2 seconds
```

### HTTP Request
```js
import requst from 'postman-request'

requenst(url, (error, response, body) => {
  const json = JSON.parse(body)
})
```
[postman-request](https://www.npmjs.com/package/postman-request) 라이브러리를 이용하면 간편하게 요청을 보낼 수 있다. 