# Note App


## 배운점
#### import/export
기본 모듈의 사용법은 [공식 문서](https://nodejs.org/docs/latest-v17.x/api/)에서 확인 가능하다.
```js
// lib\의 라이브러리는 모듈명으로 import한다.
// require 순간에 모듈이 import/실행된다.
const fs = require('fs')

const name = "myName"
const add = function (a, b) { return a + b } // 함수도 export 가능

// require(module)의 반환값을 설정
// exports는 하나만 설정 가능하다.
module.exports = name
```

#### File System
```js
const fs = require('fs')

// 파일쓰기 : fs.writeFile
fs.writeFileSync(fileName, data);  // 파일 쓰기
fs.appendFileSync(fileName, data); // 내용 추가하기
```

#### NPM
Node.js Package Manager. 터미널을 통해 설치하고 패키지 간 종속성을 관리한다.
* 패키지 저장소: [https://www.npmjs.com/](https://www.npmjs.com/)
```js
> npm init
// 패키지 설치 및 삭제 (-g: 전역 패키지)
> npm install <package name>@version
> npm install <package name>@version -g
> npm uninstall <package name>
> npm uninstall <package name> -g

// 종속성 설치 (package.json)
> npm install 
```

#### Third Party Packages
[**Validator**](https://www.npmjs.com/package/validator): 이메일, 전화번호 등 여러 양식을 검사해주는 패키지
```js
> npm install validator
import validator from 'validator';
```

[**chalk**](https://www.npmjs.com/package/chalk): 로그 메시지에 색상을 입히는 패키지. 로그 메시지에 대해서 Chaining으로 여러 스타일링이 가능하다.
```js
// Modifier, Color, Background로 스타일링 가능
// - Modifier: bold, dim, reset, italic, underline, inverse, strikethrough...
// - Color: green, blue, black, gray ...
// - Background: bgGreen, bgBlue, bgBlack ...
import chalk from 'chalk'
const successMsg = chalk.white.bgGreen.bold("Success!!")
console.log(successMsg)
```

[**nodemon**](https://www.npmjs.com/package/nodemon): 

## 오류
### **import 오류**
 최상위 js 파일이 아닌, 모듈 또는 폴더 내 js 파일에서 ```import``` 문을 사용할 경우 다음과 같은 오류를 출력한다.  
| ```Cannot use import statement outside a module```
=> package.json에 다음과 같이 type 속성을 추가하거나 파일 확장자 명을 mjs로 바꾸어야 한다. mjs는 해당 파일이 node.js에서 사용할 ES 모듈임을 명시한다.
```js
// package.json
{
  "type": "module"
}
```

### **Powershell 보안 오류**
Windows Powershell이 제3자 스크립트 실행을 제한하는 문제. Powershell을 관리자의 권한으로 실행하여 ```Set-ExecutionPolicy Unrestricted```를 입력하여 외부 스크립트를 실행 가능하도록 한다.
