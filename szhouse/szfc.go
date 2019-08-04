package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


var projectDetail = "http://zjj.sz.gov.cn/ris/bol/szfdc/projectdetail.aspx?id=39156"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {

	log.Println("======== start ========")

	url := projectDetail

	resp, err := http.Get(url)

	body := resp.Body

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("======== parse start ========")

	// s := doc.Find("form").Find("div[class=wrap]").
	// 	Find("table").First().Find("tbody").Find("tr").
	// 	Find("td")

	s := doc.Find("table")

	log.Print(s.Length())

	//s := doc.Find("form")

	// s.Each(func(i int, s *goquery.Selection) {
	// 	out := strings.TrimSpace(s.Text())
	// 	//out = strings.Replace(out, "\n", "", -1)
	// 	fmt.Printf("%d  content:%s length:%d \n", i, out, len([]rune(out)))
	// })

}

// func getBody(url string) io.Reader {
// 	resp, err := http.Get(url)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	body, e := ioutil.ReadAll(resp.Body)
// 	defer resp.Body.Close()

// 	if e != nil {
// 		log.Fatal(e)
// 		fmt.Printf("%s", body)

// 	}

// 	return resp.Body
// }
