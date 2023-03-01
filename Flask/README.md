# Flask Backend Practice
 Flask를 이용한 백엔드 설계 실습입니다. 
 
## 주요 노트
#### 인증
* 암호화: 비밀번호 암호화는 단방향 해시함수 라이브러리인 bcrypt 활용  
   => database/db.py::add_new_user()
* 인증토큰: JWT(JSON Web Token)으로 인증토큰 발급  
   => blueprint/miniter.py::login()

#### 기타
* 데코레이터: 
* 
## 고민
* **MongoDB의 아키텍처**  
  : 최상위 클래스인 MongoClient가 게이트웨이가 되어 여러 가지 기능을 실행한다. 다양한 컬렉션이 있는데, 기능이 늘어나면 MongoClient의 코드양도 길어질 것이다. 어떻게 분리할 수 있을까?

* **데코레이터 아키텍처**  
  : miniter.py에는 로그인을 요구하는 데코레이터가 있다. 이러한 데코레이터들은 어떻게 분리할 수 있을까?