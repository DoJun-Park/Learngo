package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=golang"

func main() {
	totalPages := getPages()

	for 
}

func getPages() int { //페이지 수 가져오는 함수
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	// 문장에서 에러가 발생하더라도 res.Body를 Close하여 메모리가 새어 나가는것을 막을 수 있다.
	defer res.Body.Close()

	// query docu 로드
	doc, err := goquery.NewDocumentFromReader(res.Body) //res.Body는 기본적으로 byte
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length() //pagination 클래스 내에 링크 갯수
	})

	return pages

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
