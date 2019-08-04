package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var (
	index       = "http://zjj.sz.gov.cn/ris/bol/szfdc"
	projectList = "http://zjj.sz.gov.cn/ris/bol/szfdc/index.aspx"
)

/*
ProjectBrief , brief information
*/
type ProjectBrief struct {
	Seq         string
	PreNumber   string
	NumberLink  string
	Name        string
	NameLink    string
	Company     string
	District    string
	ApproveTime string
}

/*
  ProjectBriefList,  project list
*/
var ProjectBriefList []ProjectBrief

func main() {
	url := projectList

	log.Printf("start , request url %s", url)

	// first requst . In some situation, the method should be GET
	// firstResp, err := http.Get(url)
	// if err != nil{
	// 	log.Println(err)
	// }

	// parseData(firstResp.Body)

	// requestForData(url, "2", "10")

	var page int

	for page = 1; page <= 50; page++ {
		resp := requestForData(url, page, 10)

		log.Println("==== parse start ====")

		parseData(resp.Body)

		// for debug
		// b, _ := json.Marshal(ProjectBriefList)
		// log.Printf("result : %s", string(b))

	}

	writeToFile("500_house_projectlist.txt", ProjectBriefList)

}

func requestForData(posturl string, curPage, size int) *http.Response {

	p := strconv.Itoa(curPage)
	s := strconv.Itoa(size)

	form := url.Values{
		"scriptManager2":       {"updatepanel2|AspNetPager1"},
		"__EVENTTARGET":        {"AspNetPager1"},
		"__LASTFOCUS":          {""},
		"__VIEWSTATE":          {"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"},
		"__VIEWSTATEGENERATOR": {"2A35A6B2"},
		"__VIEWSTATEENCRYPTED": {""},
		"__EVENTVALIDATION":    {"QXE4Dtr9D9E0btYt+fxiqZ/Ed8a60n5SNXjXavMY9CaSSJiEJZwRWpkx4SlAtDTnbT95GxQvqgHSnBOJDpxYCh+1Ow8DoXRJEAqTpLUYzUPvrNIQ8irNC394rEO01KdZA+4gE6ASepTUlW8H6pJBGeXiZ8BL9r5hJcU6N2fk6RxALG0qskvi36vkt+neDru9m+1eFsPA3lFErwexeYHYYS/Wubo="},
		"tep_name":             {""},
		"organ_name":           {""},
		"site_address":         {""},
		"__EVENTARGUMENT":      {p},
		"ddlPageCount":         {s},
	}

	log.Printf("current page: %s, size:%s", p, s)

	resp, e := http.PostForm(posturl, form)
	if e != nil {
		log.Print(e)
	} else {
		log.Println("request succeed")
	}

	// // debug
	// body, e := ioutil.ReadAll(resp.Body)
	log.Printf("status:%s, length:%d  ", resp.Status, resp.ContentLength)

	return resp

}

/*
 save to file
*/
func writeToFile(fileName string, list []ProjectBrief) {
	log.Printf("begin write file , size:%d", len(list))

	fo, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("create file success")
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for _, b := range list {

		l, _ := json.Marshal(b)

		_, e := fo.WriteString(string(l) + "\n")

		// fo.Sync()
		defer fo.Close()

		if e != nil {
			log.Println(e)
		}
		// log.Println("write line:" + string(n) + string(l))
	}
}

func parseData(respBody io.ReadCloser) {
	// dom
	doc, err := goquery.NewDocumentFromReader(respBody)
	if err != nil {
		log.Fatal(err)
	}

	// parse
	projects := doc.Find("table").Find("tr[class]")

	projects.Each(func(pNum int, selection *goquery.Selection) {

		one := selection.Find("td")

		b := ProjectBrief{}

		one.Each(func(num int, sel *goquery.Selection) {

			val := strings.TrimSpace(sel.Text())

			// log.Printf("num:%d %s \n", num, val)

			switch num {
			case 0:
				if val == "" {
					return
				}

				b.Seq = val
				break
			case 1:
				b.PreNumber = val
				v, _ := sel.Find("a").Attr("href")
				b.NumberLink = strings.Trim(v, ".") // remove dot

				break
			case 2:
				b.Name = val
				b.PreNumber = val
				v, _ := sel.Find("a").Attr("href")
				b.NameLink = v
				break
			case 3:
				b.Company = val
				break
			case 4:
				b.District = val
				break
			case 5:
				b.ApproveTime = val
				break
			default:
				break
			}

		})

		ProjectBriefList = append(ProjectBriefList, b)

	})
}

// func parseOne(num int, sel *goquery.Selection,) {
// 	log.Printf("count: %d \n", count)
// 	count++

// 	val := strings.TrimSpace(sel.Text())

// 	log.Printf("num:%d %s \n", num, val)

// 	switch num {
// 	case 0:
// 		b.Seq = val
// 		break
// 	case 1:
// 		b.PreNumber = val
// 		v, _ := sel.Find("a").Attr("href")
// 		b.NumberLink = strings.Trim(v, ".") // remove dot

// 		log.Print("number link  "  + b.NumberLink)
// 		break
// 	case 2:
// 		b.Name = val
// 		break
// 	case 3:
// 		b.Company = val
// 		break
// 	case 4:
// 		b.District = val
// 		break
// 	case 5:
// 		b.ApproveTime = val
// 		break
// 	default:
// 		break
// 	}

// }
