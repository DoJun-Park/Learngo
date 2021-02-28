package main

import (
	"fmt"
	"time"
)

// var errRequestFailed = errors.New("Request failed")

// func hitURL(url string) error {
// 	fmt.Println("Checking", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		return errRequestFailed
// 	}
// 	return nil
// }

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
	go sexyCount("dojun")
	sexyCount("gildong")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
