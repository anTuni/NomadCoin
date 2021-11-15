# Go lang to Encrypto coin

## 3.0 Go Project 시작하기

### go mod 명령어 실행

    1.go.mod 파일 생성
    - 모듈의 위치 기록: Node.js 의 package.json 과 같은 역할
    - 모듈이 어디에 위치하는지 알려줌
    2.main.go 파일 작성
    -VScode Go extension 사용시 package가 자동으로 import 됨
    -"fmt" package 의 println 함수로 콘솔에 찍을 수 있음'
    - main in file main.go function will be execured by default when we run main.go

---
Q.what go.mod is for?
for dependency tracking when we use package from other module

Q.what is package in main.go?
a package is a way to group functions, and it's made up of all the files in the same directory

---

## 3.1 Variables(변수)

### 변수만들기

     var name string = "val" 
     - '생성자' '이름' '타입' = "값" 순으로 선언 가능

#### 함수 내에서 변수를 선언할 땐 shortcut이 있다

     '이름' := '값'
     #### -colon과 등호를 같이쓰면 컴파일러에서 값의 타입을 변수의 타입으로 인식한다

#### 함수 밖에서 선언할 땐 shorcut을 사용하지 못한다

#### 타입은 boolean, string, int 등이 있다

---

## 3.2 Functions(함수)

### 함수의 선언

#### 함수를 선언할 땐 함수의 이름, argument의 타입, return 값의 타임을 선언한다

     ex)
     func nameoffunction (arg int,arg string) int{
         return arg
     }

#### 같은 tpye의 arguments 여러 개를 배열로 받을 수 있다

#### 배열,string 으로 반복 작업을 하려면 다음과 같이 쓴다

    for index,elem := range "iterator" {
         /// do something
     }

#### print 할 때 값을 formatting 할 수 있음 binart,digit,hexa,string 등

---

## 3.3 fmt

### package fmt

    package for formatting data

    fmt.Sprintf()

## 3.4 Array and Slice

    * an Array has certain length (Can't be infinite)

    * a Slice is unlimited array

    * declare and assign

    Array : name := [length] type {el1,el2,...}
    ex) sampleArray := [3]sting{"str1","str2","str3"}

    Slice name := [] type {el1,el2,...}
    ex) sampleSlice := []sting{"str1","str2","str3"}

    * we can add more element in a slice by append function
        append(slice,value) will return a slice that add the new element to original slice
        인수로 받은 Slice 자체를 바꾸는 게 아니라 요소를 추가한 Slice를 반환함

## 3.5 Pointers

    * memory address
    변수 이름 앖에 &를 붙이면 그 값을 참조하는 게 아니라 메모리 주소(포인터)를 참조함
    포인터 변수 앞에 *을 붙이면 해당 주소의 값을 참조함

## 3.6 Structs and reciever function

### structs

     * type name struct {    var type }
     * func (n name) recieverFunction(arg type){    }
     * func (n *name) recieverFunction(arg type){    }

## 3.7 Structs with Pointers

* package 를 새로 만들었을 때 첫 번째 알파벳이 uppercase(대문자) 인 것만 export 됨

* reciever function 선언할 때 structs 이름 앞에 *을 안 붙이면 stucts 값을 참조함.
structs instance를 참조하려면 * 을 꼭 붙여야 해당하는 reciever function을 호출한 instance를 수정할 수 있음