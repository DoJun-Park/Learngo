package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=golang"

func main() {
	getPages()
}

func getPages() int { //페이지 수 가져오는 함수
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	// 문장에서 어떤 에러가 발생하더라도 항상 파일을 Close할 수 있도록 하여 메모리가 새어 나가는것을 막을 수 있다.
	defer res.Body.Close()

	// query docu 만드는 방법
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each()

	fmt.Println(doc)

	return 0
}

// 에러를 계속 체크해줘야 하기 때문에 따로 함수를 만들어서 처리

func checkErr(err error) {
	if err != nil { //에러가 있다면
		log.Fatalln(err) //프로그램 끝내기
	}
}

func checkCode(res *http.Response) { //http.Get()에서 http.Response의 포인터이기 때문에 앞에 * 를 붙여준다.
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}
