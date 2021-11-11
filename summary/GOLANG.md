# 3.0 Go Project 시작하기

## go mod 명령어 실행

    1.go.mod 파일 생성
    - 모듈의 위치 기록: Node.js 의 package.json 과 같은 역할
    - 모듈이 어디에 위치하는지 알려줌
    2.main.go 파일 작성
    -VScode Go extension 사용시 package가 자동으로 import 됨
    -"fmt" package 의 println 함수로 콘솔에 찍을 수 있음

---
Q.what is module?
Q.what is package in main.go?

# 3.1 Variables(변수)

## 변수만들기

     var name string = "val" 
     - '생성자' '이름' '타입' = "값" 순으로 선언 가능
     
     함수 내에서 변수를 선언할 땐 shortcut이 있다.
     '이름' := '값'
     -colon과 등호를 같이쓰면 컴파일러에서 값의 타입을 변수의 타입으로 인식한다

     함수 밖에서 선언할 땐 shorcut을 사용하지 못한다.

     타입은 boolean, string, int 등이 있다.

---