package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=golang"

type extractedJob struct {
	id       string
	location string
	titile   string
	salary   string
	summary  string
}

func main() {
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) { //하나의 페이지에 대해 정보
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		extractJob(card)
	})
}

func extractJob(card *goquery.Selection) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	fmt.Println(id, title, location, salary, summary)
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
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
