package main

import (
	"errors"
	"fmt"
	"net/http"
)

// go루틴과 채널의 rule을 아는 것이 중요
// 1. main 함수가 끝나면 go루틴은 무의미해진다. go루틴이 끝나지 않았더라도
// 2. 채널을 통해 받을 데이터 그리고 보낼 데이터의 타입을 정확하게 지정해줘야 한다.
// 3. 메세지 보내는 방법은 <- 이용
// 4. <- 는 blocking operation으로 프로그램이 이 동작을 하기 전까지 멈추는 것을 뜻한다.

var errRequestFailed = errors.New("Request failed")

// -------------Go routine을 사용하지 않은 URL Checker-----------------
// func hitURL(url string) error {
// 	fmt.Println("Checking", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		return errRequestFailed
// 	}
// 	return nil
// }
// ----------------------------------------------------------------------

// -----------------------Go routine & channel 실습-----------------------------
// func isSexy(person string, c chan string) bool {
// 	time.Sleep(time.Second * 10)
// 	c <- person + " is sexy"
// 	return true
// }
// ----------------------------------------------------------------------

// -------------Go routine을 사용한 URL Checker-----------------------------
type requestResult struct {
	url    string
	status string
}

func hitURL(url string, c chan<- requestResult) { //chan<- 는 send-only를 뜻한다.
	resp, err := http.Get(url)
	status := "Ok"
	if err != nil || resp.StatusCode >= 400 {
		status = "Failed"
	}
	c <- requestResult{url: url, status: status}
}

// ----------------------------------------------------------------------

func main() {

	// -------------Go routine을 사용하지 않은 URL Checker-----------------
	// // var results map[string]string
	// // 만약 위와 같이 results를 하고 값을 넣으면 panic 에러가 발생한다.	초기화되지 않은 map에는 어떠한 값을 넣을 수 없기 때문이다.
	// // 초기화하는 방법은 아래와 같이 두가지 방법이 있다.

	// // 1.
	// // var results = map[string]string{}

	// // 2. make() 이용
	// var results = make(map[string]string)

	// urls := []string{
	// 	"https://www.airbnb.com/",
	// 	"https://www.google.com/",
	// 	"https://www.amazon.com/",
	// 	"https://www.reddit.com/",
	// 	"https://www.google.com/",
	// 	"https://soundcloud.com/",
	// 	"https://www.facebook.com/",
	// 	"https://www.instagram.com/",
	// 	"https://academy.nomadcoders.co/",
	// }

	// for _, url := range urls { //range를 사용하면 인덱스, 요소값 사용할 수 있다.
	// 	result := "OK"
	// 	err := hitURL(url)
	// 	if err != nil {
	// 		result = "FAILED"
	// 	}

	// 	results[url] = result
	// }

	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }
	// ----------------------------------------------------------------------

	// -----------------------Go routine & channel 실습-----------------------------
	// c := make(chan string)
	// people := [2]string{"dojung", "gildong"}
	// for _, person := range people {
	// 	go isSexy(person, c) //isSexy()함수가 return값이 있다고 해서 result := go isSexy(person)  으로 사용할 수 없다.
	// }
	// // channel에서 메시지 받아와서 출력하는 방법
	// // 방법 1.
	// // result := <-c
	// // fmt.Println(result)
	// // 방법 2.
	// // fmt.Println("Received this message: ", <-c) // '<-'는 채널로부터 메세지를 받아올때 사용한다.
	// // fmt.Println("Received this message: ", <-c)

	// // 만약 받으려는 메세지가 많아지면?? -> loop 사용!!
	// for i := 0; i < len(people); i++ {
	// 	fmt.Println(<-c)
	// }
	// ----------------------------------------------------------------------

	// -------------Go routine을 사용한 URL Checker-----------------------------

	c := make(chan requestResult)
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls { //range를 사용하면 인 덱스, 요소값 사용할 수 있다.
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

	// ----------------------------------------------------------------------

}
