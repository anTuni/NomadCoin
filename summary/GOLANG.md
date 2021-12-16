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

## 4.1 Our first block

해쉬 함수의 특성

블록간 체인으로 연결되는 원리
현재 블록의 데이터와 이전 블록의 Hash를 더해서 같이 Hash 함으로서
block struct 만들기 hash와 prevHash, data를 가짐

sha256 알고리즘으로 hash process

go에 내장된 sha256.Sha256 함수에 data+prevHash (string)를 ([]byte)로 전환하여 전달
-> string 변수를 slice of byte로 전화하는 방법 []byte(string)

return 값으로 받는 array를 다시 hexadecimal hash string으로 formatting 함
->fmt.Sprintf("%x",[int]byte)

## 4.2 Our first blockchain

blockchain struct에 block 저장하는 reciever function 만들기

getLastHash
addBlock
getList

## 4.3 Singleton Pattern

Package 만들기 blockchain package -> main.go 안에서 많은 작업이 이루어질 것이기 때문에 정리

Single ton pattern : sharing only one instance in the application
-> 변수를 직접 공유하지 않음(소문자로 시작함)
-> blockchain package 안에서만 접근 가능함

package를 만들고 변수에 blockchain instance를 할당한다.

package 의 method로 intance가 없으면 할당하고 있으면 인스턴스 반환하는 기능을 추가한다.

Go는 parallel 하게 작동할 수 있기 떄문에 if b == nil 조건문으로 한 번만 실행하도록 하면
동시에 일어나는 routine 에 대해서 여러번 실행하게 된다 때문에 sync package의 Once.Do func를 사용한다.

Q.What is parallel and routine?

## 4.4 Refactoring part One

sync package의 Once.Do func를 사용한다
=> 코드의 특정 부분이 한 번만 실행 되도록 하기 위해서

## 4.5 Refactoring part Two

New func

1) functiong to get blocks of the blockchain
2) functiong to append a block to the blockchain

## 5.0 Setup

make a  Server Side render Website with go and go standard

using http package

## 5.1 Rendering Templates

using template package
html templates 파일 parsing and execute

use variables in template

## 5.2 Rendering Blocks ( splice of Block's pointer)

ieterate by splice in template

     {{range Blocks}}
     {{.Var}}~~~
     {{end}}

using mvp.css

## 5.3 Using partials

html template file에서 반복해서 사용하는 부분을 나눠서 저장하고 불러오기

define partial template

    {{define "name"}}
    ...template html
    {{end}} 

Using partial template (Somedata에 변수나 값을 직접 전달할 수 있다.)

    {{tempalte "name" somedata}}
    ex 1) {{tempalte "name" .Pagetitle}}
    ex 2) {{tempalte "name" "Title"}}

load all template files in main func

    pattern
    templates = template.Must(template.ParseGlob(templateDIr + "pages/*.gohtml"))
    templates = template.Must(templates.ParseGlob(templateDIr + "partials/*. gohtml"))

excute template with defined name in HandleFunc

    templates.ExecuteTemplate(rw, "add", nil)

## 5.4  Adding Blocks

handling http request

r.Method 값에 따라 분기.
request body 값 읽기

    r.ParseForm()
    data := r.Form.Get("blockData")
Q.r.Form is a Map, then what is the map type?

## 5.6 Recap http sever code

http response code 308 -> redirect

## 6.0 REST API Setup

거래, 마이닝 등을 위한 확인
GO 에서 JSON 사용하기

json package 의 Marshal() 함수 go->JSON , Unmarshal() 함수 JSON ->go로 데이터 형식을 변환

## 6.1 Marshal and Field Tags

http 요청에 텍스트 response

    fmt.Fprint... 

JSON 형식으로 response 하기 -> response header 수정

     http.ResponseWriter.Header().Add("Content-Type","application/json")

response json in 3steps(3Lines of code)

    1)Marshal 함수로 go-> json 변환
    2)err 핸들링
    3)Fprintf() 로 출력

response json in 1steps(1Lines of code)

    json.NewEncoder().Encode(...data)

json key(field) 값을 다 lowercase로 바꾸는 방법 (Field tags)

1) struct type 정의 할 때 각 키마다 json 응답시에 보여질 field key 명시.
null 값일 때 field 생략, field  항상 생략 가능
{
    Data string `json:"data,omitempty"`
    Data string `json:"-"`
}

## 6.2 Marshal Text

응답을 보낼 때 URL 값 앞에 생략된 URL을 붙이고 싶음
ex ) '/add' -> 'http://localhost:4000/add'

Stringers interface를 사용한다.-> Struct의 reciever func을 정해진 규칙대로 만들면됨
여기서는 String() string {}
fmt의 print 함수로 Struct를 출력하려고 할 때 String Method 가 있는지 확인하고 있으면 그 method의 return value를 출력함

URL type 을 선언 하고 TextMarshal interface를 사용

## 6.3 JSON Decode

POST,GET method to /blocks

recieve data from POST Request body into struct
ex) json.NewDecoder(r.Body).Decode(&struct)

## 6.4 NewServeMux

* go routine으로 http server 동시에 실행하기

* 서로 다른 두 port를 열어 http 서버를 실행할 때
서로 다른 multiflexer를 정의 해줘야한다.
기본적으로 DefaultMultiflexer를 사용하기 때문에 같은 url pattern을 사용할 경우 충돌이 난다.
http package의 NewServeMux 함수로 새로운 multiflexer를 초기화 할 수 있고
http request handler를 정의 할 때 새로 초기화 한 multiflexer를 사용한다.

## 6.5 GorillaMux ( install and use first dependency)

to recieve http parameter by url

1) Install Grilla/mux package from github
2) http.ServeMux를 mux.NewRouter()로 수정
mux.NewRouter()에 의해 반환되는 Multiflexer는 URL에 규칙을 정할 수 있고, mux.Vars(r http.request) map로 url parameter를 가져올 수 있다.
또 Method를 정할 수 있다.

mux.Vars() 함수는 map을 반환한다.

## 6.6 Atoi

url parameter로 받은 string 을 interger로 변환하여 blockchain의 특정 block 조회하기

block에 Height 추가
GetBlock(height int) *block 함수 생성

url parameter를 int로 변환하고 함수 호출

## 6.7 error hnadling

GetBlock() 함수에서 전달받은 height 가 총 len보다 클 때 에러 보여주기
errors.New() 로 error type 변수 만들어서 반환하기
error 변수 formmatting 해서 json response하기

adapter pattern

http.Handler는 ServeHttp(http.ResponseWriter, *http.Request) 메소드로 구현되는 인터페이스이다.

http.Handler 인터페이스를 구현하려면
타입을 선언하고 ServeHttp(http.ResponseWriter, *http.Request)를 구현해야한다.

하지만 http.HandlerFunc 타입이 있다.
이 타입은 func(http.ResponseWriter, *http.Request) 타입의 value를 가진다

때문에 따로 타입을 선언하고 메소드를 구현할 필요없이 http.HandlerFunc 타입의 함수값을 선언해주면 된다.

## Extra explaing about interface

interface는 무엇인가.
Go 프로그램은 package 단위로 구성된다.
Package A,B,C 가 있다고 하자
package A에서는 type I interface를 M() 메소드를 가짐으로 정의하고
Af 함수는 value of type I interface 를 argument로 받아 I interface의 method를 실행한다.

package B 에서 type B가 있고 이 type B는 M()메소드를 가진다.
package C 에서 type C가 있고 이 type C는 M()메소드를 가진다.

B,C 모두 package A를 import 하고 함수 Af에 각각 type B,type C를 argumnet로 전달해 실행한다.

## 7.0 CLI 7.1Parsing Command

Command Line Interface
명령어를 입력해서 코드를 실행하는 것.

os.Args로 입력한 command parsing하기

## 7.2 FlagSet 7.3 flag

flag.NewFlagSet() 로 flag 파싱하기

    flag 란?
    명령줄에서 '-'을 붙여서 전달할 인자를 명시하는 부분
    ex) go run main.go -port=4000

command line flag 받아서 알맞은 작업 실행하기

## 8.0 Database introduction

블록체인 데이터를 메모리에 저장할 순 없기 때문에 데이터 베이스에 저장한다.
이 강의에서는 Bolt라는 데이터베이스를 사용한다.
"key":"value" 형태로 저장되는데 더 이상 업데이트가 없다.

## 8.1 Creating the Database

데이터 베이스 만들기
bolt.Open(...) 으로 데이터베이스를 초기화 한다.
.db 파일을 해당 directory에 만든다.

RDB의 테이블 개념으로 Bucket 이 있다.

## 8.2 A New Blockchain

divide blockchain code in two part
the one is part of Block and another is of Chain

Block part will have struct that have many information of the Block(like Data,Hash,Previous Hash,Transaction, Height ... )

Chain part will have NewestHash and Height of the BlockChain

## 8.3 Saving a Block

블록을 데이터베이스에 저장하기 위해 hash와 data를 []byte type으로 받아
hash를 key 값으로 data를 value로 저장
=> DB instance method 사용

block에서 db의 저장함수를 호출, hash와 data를 인자로 전달해야함
->data는 블록 자체를 []byte type 으로 encoding 해야함

gob  package로 encoding하기
byte.Buffer type의 변수 선언
gob.NewEncoder()로 Encoder initializing
encoder instancer 의 Encode() 메소드 사용= > error handling 후

byte.Buffer type에 저장된 []byte type 반환

## 8.4 Persisting Ther BlockChain

블록체인을 저장하기위해
똑같은 작업 필요 data를 byte로 전환하는 함수 만들기
=> utils 함수로 만듬

함수 선언시 argment로 어떤 형태든 다 받을 때
func NameOfFunc (i interface{}) {...}
interface는 base type 이라서 어떤 type이든 다 받을 수 있음
형태로 씀

## 8.5  Restoring the Blockchain

블록체인 initializing 할 때 이미 저장된 내용 확인하고
chain 값 업데이트하기

db로부터 checkpoint 존재여부 확인
=>nil 값일 땐 첫 블록 만들기
=>checkpoint 존재 시 []byte 로 저장 돼있기 때문에 복원 메모리상 블록체인으로 저장

## 8.6 Restoring Block

find a Block from a hash

1) requset a block with hash string from URL
2) Get block from database with that hash
3) restore a block  from []byte and return

database로부터 온 []byte type의 데이터로 부터 특정 struct type의 데이터를 얻기 위해서
gob.Newdecoder().Decode()를 사용한다.
Decode() 함수는 i interface{}의 "Pointer"를 인자로 받아서 해당 포인터에 값을 저장한다.

## 8.7 All Blocks

func for closing Database

Get all blocks using loop and func getting a block
with initial value of blockchain's NewestHash start looping to find all blocks

## 8.8 recap Chapter 8 Database

Implement persistence of Blockchain with database

Save Block and state of Chain on Database ( bucket of Bolt DB)

Restoring from Database's bytes data

## 9.0 Introduction to PoW

블록을 추가하기 어렵게 하기 위해서 작업 증명 방식을 사용한다.
답을 찾기는 어렵지만 검증하긴 쉽다.

작업 증명 vs 지분 증명이 있다.

modifying block structure

## 9.1 PoW Proof of Concept

해쉬 -> deterministic function, ono way function 으로
같은 입력은 항상 같은 출력, 입력으로부터 출력 추정 불가
=> 답이 맞는지 확인하긴 쉽지만, 답을 찾기는 어려움

작업 증명 find a hash start with certain number of Zeros
zero의 숫자(n)는 difficulty로 조정
nonce의 숫자만 변경하면서 답을 찾음.

## 9.2 Mining block

블록체인에 블록을 추가할 땐 이전 블록의 해쉬값과 새로 추가할 데이터로 새로운 해쉬를 만들어야한다.
이 때 어떤 해쉬든 다 받는 다면 누구나 블록을 추가할 수 있다. 그렇게 되면 거짓 데이터가 블록체인을 공격할 수 있다.

채굴한다는 것은 정해진 작업 증명 방식에 따라 다음 블록의 hash를 만들기 위한 nonce를 찾는다는 것이다.

무엇이 블록체인을 안전하게 하는가.
블록을 블록체인에 추가할 때 들이는 노력 즉, 컴퓨팅 파워가 필하다.
블록 체인의 모든 블록의 내용을 수정하기 위해 들이는 노력이 그에 의한 대가보다 커야한다.
그 노력의 크기를 결정하는 것이 작업증명 방식이다.
작업 증명은 난이도가 조절가능해야한다. 일반적으로 더 어려워 지는 게 맞다.

작업 증명의 방식이 중요하다.
가치를 만들어 내야한다. 기본적으로 가치는 에너지와 시간으로 부터 나온다.
나온다기보다 계산된다.

## 9.3 -9.4 Difficulty part One,Two

Difficulty를 network로부터 자동으로 결정되도록

add hashing function to utils package

difficulty defualt value
-> difficulty interval(난이도를 조절할 간격? 이 강의에서는 blockchain의 height 로 하였음)

기본(초기)난이도를 정하고 , 난이도를 조절할 간격을 정한다.
그리고 블록이 생성 될 때마다 타임스탬프를 저장해 블록이 얼마나 빠르게 추가되는지 확인한 후 원하는 만큼 난이도를 조절할 수 있도록 한다.

## 9.5 conclusion

check blochain network can adjust difficulty by itself as adding blocks

check blockcahin network can assign difficulty dynamically

## 10.0 - 10.1 Transaction Intro

블록체인 네트워크에서 이뤄지는 거래의 의미

코인을 주고 받는 다는 건 무슨 의미일까.
코인은 어디서 생성되고 어떻게 채굴자에게 보상으로 줄 것인가
누가 얼마나 가지고 있는가 확인 할 수 있을 것
비트코인의 accounting model

Unspent Transaction Output(uTxOut) 모델

Transaction is composed inputs and outputs
input has information that is about how much do someone have before the transaction

output has information for after the transaction

거래는 인풋과 아웃풋으로 이루어진다.
이 거래를 하기전 가지고 있는 코인?의 양과
거래가 끝나고 가지고 있는 코인의 양이 나온다.

예를 들어 A가 B에게 x coin을 보낸다하면.
인풋엔 원래 A가 가지고 있는 총금액이 있고
아웃풋에 A가 가진 금액에서 x만큼 빠진 금액, 그리고 B가 원래 가지고 있던 금액에서 x를 더한 금액이 있는 것이다.

채굴자가 보상을 받을 경우엔 인풋에 블록체인 네트워크에서 인풋이 만들어지고, => 화폐를 발행하는 것 과 같은 맥락=>coinbase transaction 이라 한다.

노드 간 거래일 경우 보내는 쪽 노드에서 인풋이 만들어진다.

## 10.2 coinbase transaction

make transaction to reward for Miner
It's gonna be made by blockchain network to the Miner

block을 생성할 때 transaction 정보를 담는다.
transaction 은 고유한 Id 를 가진다. 시간, 거래 Input과 Output으로 구성된다.

## 10.3

get TxOuts by owner

1 모든 블록의 거래 내역 중 Output을 가져오고
2 그 중 원하는 owner의 거래 내역만 불러온다.
3 불러온 거래 내역의 모든 값을 더한다.

## 10.4 Mempool

거래들이 0승인되기 전에(블록에 추가되기 전에) 저장되는 곳

## 10.5 AddTx

거래를 mempool에 저장하기

Mempool은 메모리상에 저장? 됨 . 블록에 추가 될 때 데이터 베이스에 추가됨

보내는 사람, 받는 사람, 거래량

보내는 사람의 잔고 확인 (이전 거래 Outputs 확인) -> 거래량 만큼 있는지

## 10.6 makeTx to AddTx

보내는 사람의 잔고 확인 (이전 거래 Outputs 확인) -> 거래량 만큼 있는지

1.TxIns 만들기
2.TxOuts 만들기
3.Tx 만들기

1.TxIns 만들기

TxIns에 들어간 amount의 양을 확인하기 위한 변수 total 초기화

보내는 사람의 Outputs slice 로 loop를 돌린다.
하나의 Txout으로 TxIn을 만들어 TxIns에 추가한다.
추가한 Txout의 Amount를 total에 더한다.
total이 보내고자하는 amount보다 크거나 같으면 loop를 break 한다.

2.TxOuts 만들기

2-1 보내는 사람 TxOut : 잔돈 확인
total과 amount의 차가 0보다 크면 그만큼을 TxOuts에 보내는 사람의 TxOut추가한다.

2-2 받는 사람 TxOut
amount 만큼 추가한다.

3.Tx 만들기
timestamp와 TxIns,TxOuts로 만들고 ID 초기화

Q. 보내는 사람, 마이너 는 어떻게 가지고 오나.
Q. 블록체인 네트워크의 아키텍쳐
각 노드는 어떤 정보를 가지고 있고 어떤걸 갖추고 있는가.
블록체인의 모든 정보를 어느 주기로 어떻게 검증하는가.
