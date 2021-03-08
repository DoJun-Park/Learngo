package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term

	var jobs []extractedJob
	c := make(chan []extractedJob) //일자리 정보가 여러개 전달되므로 extractedJob이 아닌 []extractedJob
	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c)
	}

	for i := 0; i < totalPages; i++ { //전체 페이지 수만큼 메세지가 오기 때문에
		//jobs에 일자리를 저장하는 것을 같지만, return값을 받아서 저장하는 것이 아닌 메세지를 통해 받는다.
		extractJobs := <-c
		jobs = append(jobs, extractJobs...) // extractedJobs를 각각의 배열로 저장하는 것이 아닌 하나의 배열로 만들기 위해 contents를 가져온다. 이때 '...' 이용
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int, url string, mainC chan<- []extractedJob) { //하나의 페이지에 있는 일자리 추출
	var jobs []extractedJob
	C := make(chan extractedJob)

	pageURL := url + "&start=" + strconv.Itoa(page*10)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, C)
	})

	for i := 0; i < searchCards.Length(); i++ { //전달받을 메세지의 수는 카드의 수와 같음.
		job := <-C
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) { //카드에서 일자리 정보 추출
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int { //총 페이지 수 가져오는 함수
	pages := 0
	res, err := http.Get(url)
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

func writeJobs(jobs []extractedJob) { //일자리를 csv파일로 저장하기 위한 함수
	file, err := os.Create("jobs.csv") //jobs.csv 파일 생성
	checkErr(err)

	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	file.Write(utf8bom)

	w := csv.NewWriter(file) //Writer 생성
	defer w.Flush()          //함수가 끝날 때 파일에 데이터 입력

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
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
