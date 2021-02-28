// main은 컴파일을 위해 필요한 것

// go 이론 main

package main

import (
	"fmt"
	"strings"
)

func mulitfly(a, b int) int {
	return a * b
}

// func lenAndUpper(name string) (int, string) {
//   return len(name), strings.ToUpper(name)
// }

//argument
func repeatMe(words ...string) { //argument를 무제한으로 받고 싶은 경우, '...' 점 3개를 사용한다
	fmt.Println(words)
}

// naked return

func lenAndUpper(name string) (lenght int, uppercase string) { //return할 variable을 여기서 미리 정의
	defer fmt.Println("I'm done!") //defer는 함수가 끝나고 나서 실행시키는 코드
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return // return할 variable을 굳이 명시하지 않아도 됨
}

// for문, range

func superAdd(numbers ...int) int {
	fmt.Println(numbers)
	//number안에서 조건에 따라 반복 실행
	// for index, number := range numbers { //range는 for 안에서만 적용, range는 index도 넘겨준다.
	//  fmt.Println(index, number)
	// }

	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
	return 1
}

//if,else  / switch

func canIDrink(age int) bool {
	// if age < 20 {
	//  return false
	// }
	// return true
	// if koeran_Age := age + 2; koeran_Age < 20 { //if부분에 변수를 생성할 수 있다.
	//  return false
	// }

	// return true
	// switch {
	// case age < 20:
	//  return false
	// case age >= 20:
	//  return true
	// }
	// return true

	switch koeran_Age := age + 2; koeran_Age {
	case 10:
		return false
	case 20:
		return true
	}
	return true
}

type person struct {
	name    string
	age     int
	favFood []string //slice
}

func main() {
	// // ------------export--------------------------------------------------------------
	// fmt.Println("Hello World") //Go에서 function을 export하고 싶은 경우 function을 대문자로 시작해준다.
	// something.SayHello()       //대문자로 시작했기 때문에 export된 func이다.
	// something.sayBye() 는 private func(시작이 대문자 아님)이므로 export할 수 없다.
	// // --------------------------------------------------------------------------------
	// // ------------variable and Constants-----------------------------------------------
	// const name = "dojun" //const는 상수로 값을 변경할 수 없다. , 이 코드는 타입이 없는 상수이다.
	// Go는 type언어로, '타입이 무엇이다'라는걸 알려줘야 한다.
	// const name string = "dojun"
	// fmt.Println(name)
	// // var name string = "dojun" //var 변수로 수정 가능
	// name := "dojun" // 위 코드의 축약형으로 같은 의미이다.
	// // := 를 사용하면 Go가 type을 추측하여 찾아줌. 정해진 type은 임의로 변경이 불가능
	// // 축약형은 function 안에서만 가능하고 변수에만 적용 가능
	// name = "gildong"
	// fmt.Println(name)
	// // ---------------------------------------------------------------------------------
	// // ------------multiply--------------------------------------------------------------
	// fmt.Println(mulitfly(2, 2))
	// // --------------------------------------------------------------------------------

	// // ------------lenAndUpper-------------------------------------------------------------
	// totalLength, upperName := lenAndUpper("dojun")
	// fmt.Println(totalLength, upperName)
	// 만약 여러개의 값을 return받을 때, 다 받지 않고 원하는 값만 받고 싶을 때는 받고 싶지 않는 값은 '_'로 처리한다.
	// totalLength, _ := lenAndUpper(("dojun"))
	// fmt.Println(totalLength)
	// // ---------------------------------------------------------------------------------

	// // ------------여러 인자를 받고 싶을 때--------------------------------------------------------------
	// repeatMe("I", "love", "you", "so", "much")
	// // ---------------------------------------------------------------------------------
	// // ------------lenAndUpper--------------------------------------------------------------
	// totalLength, up := lenAndUpper("dojun")
	// fmt.Println(totalLength, up)
	// // ---------------------------------------------------------------------------------

	// // ------------superAdd--------------------------------------------------------------
	// superAdd(1, 2, 3, 4, 5, 6)
	// // ---------------------------------------------------------------------------------

	// // ------------canIDrink--------------------------------------------------------------
	// fmt.Println(canIDrink(2))
	// // ---------------------------------------------------------------------------------

	// // ------------Pointers--------------------------------------------------------------
	// a := 2
	// b := &a
	// *b = 12
	// fmt.Println(a, b)
	// // ---------------------------------------------------------------------------------

	// // ------------Arrays and Slices--------------------------------------------------------------
	//array의 크기를 정하고 생성
	// names := [5]string{"a", "b", "c"}
	// names[3] = "d"
	// names[4] = "e"
	// fmt.Println(names)
	//array의 크기를 정하지 않고 생성
	// numbers := []int{1, 2, 3}    //slice
	// numbers = append(numbers, 4) //append에서 첫번째
	// fmt.Println(numbers)
	// // ---------------------------------------------------------------------------------

	// // ------------Maps--------------------------------------------------------------
	// Map := map[string]string{"name": "dojun", "age": "25"}
	// fmt.Println(Map)
	// for key, value := range Map {
	//  fmt.Println(key, value)
	// }

	// // ---------------------------------------------------------------------------------

	// // ------------Struct--------------------------------------------------------------
	//방법 1 - 어떤 field가 어떤 것인지 알기 어려움
	// favFood := []string{"ramen"}
	// dojun := person{"dojun", 25, favFood}
	// fmt.Println(dojun)

	// 방법 2 - 좀 더 명확하게
	// favFood := []string{"ramen"}
	// dojun := person{name: "dojun", age: 25, favFood: favFood}
	// fmt.Println(dojun)
	// // ---------------------------------------------------------------------------------

}
