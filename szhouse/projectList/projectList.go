package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
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
ProjectBrief 项目概述
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

var ProjectBriefList []ProjectBrief

var count int = 1

func main() {
	url := projectList

	log.Printf("start , request url %s", url)

	// resp, err := http.Get(url)

	resp := requestForData(url, 1, 10)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("==== parse start ====")

	parseData(doc)

	b, _ := json.Marshal(ProjectBriefList)

	log.Printf("result : %s", string(b))

	// writeToFile("20190804houst_projectlist.txt",ProjectBriefList)

}

func requestForData(posturl string, startPage, size int) *http.Response {
	form := url.Values{

		"scriptManager2":       {"updatepanel2|AspNetPager1"},
		"__EVENTTARGET":        {"AspNetPager1"},
		"__LASTFOCUS":          {""},
		"__VIEWSTATE":          {"cRfD7sgYr7sSj9QboW1/wH7HEspaOTpbunVzTZT4mVhGp9DX/MqCYx+zYb/C1BzGbHXsghK7zuB4qJtyqPP44l7A6iDOAPbFqYx5iyh3PAICziqdwqRKWwRRXsPWfoHD3/pWSff0KeHRSPmAjtBpL3svjvSZ4{xYmbSXoiF4xycQdooOio98XTic3cXxJK4myHmufiyPk7fOXFZ2l9xOuJbQ9AkkDqDNQHif0mRr4+gpOHGSrq58D{zdvLHbovgGE6sehHvyd8heWIy3GUeoriLsfWJTiz1Os2CdSRQT7nTN8vXgnDcsRKUeRi9qm0c8f9M5Z7qCYP3PYO7+eb9Y9PRQsVEik1c8yq{1qW++RajZnOBz2heWdUbz5AyQjjFkWL77z7mV{mRzg53kSkWnq5/3a+GUMdMG224u0tmj3X+HcHZGo7io7MhLrrAD72s0NJbJmAwArrUqDZes9I82VuevDXjhzEkFD7Qktk7R9eelMP7mEFAScF3VhGvVp6IkntA+yCzdqsu8Ya{ABDIOjehVCp8hDbwbMHQvHHfH/aDoG9Yji+EO9hXFq655w36Gp5zE+HXRLieoLTRonrXEWmGqASMzhnHvdEBP4crVojVIrRbVcEmJGHQEzrirlAnzqFOSkloVHK3UIFX5yi2xVI1dlNBx9Glh4jxcva4EIzyfPP28ScRTmuTMjgtQjiHRBbjdl21r1lerEZh5EYTkJiB{HxluIRU1fGQzZP7jqjbLEG{6yTfRsyzIhT1Dr3iwUFKN6sUD4OA66bPMXIArhmUNygeSVSZzzUnvQjiJjVKyzQQetAyBTJun1kB3TcKMHp5FlNyFkpHhwrYz449/UbkTONNgavvVtdaOc8Y7GzYu9gSs+TbmfbMwOfYRVhwndVbEHKRZ7W8cvcu6LvQRLfsqZfW74z6FfGUCYZd91VUN0KF0CJgcgJfGRbiWLjSXMHsfhcgtiFaXBA9M1iuzaQlt7WKnBeYRdFBL48SFuUTyzuhIZCvfMH0iC86ASL7AUUSkE2Ws{mWr6CP7I7J68EVUmMUJFgIE5pSFh4HH9TUL30fRp{Zd1mS+7TO1DhB8BvvUGVlDb5jFq30ZP3F+VOfuPbs+JNJUCx0ZN08xhMlYZUDjlBArQ4HlQtIwhD4Cr9SJtgUsb5LsJ9FbEdb8e354360dmUWtFCpfh8k2RqV1+8jm{MVeOcP+G5s/Nkuj4i0kN5pv1/kc/ismYn3BVp+sxZTtIYx/ZbadorlPizLlZGETeJA3FSdiKaIuuFxubQLOeiTGdeAAiHLGu+64oXvXNo20kpOB0POqw1IOxnsJdt5{6zIvb7mqkwbqRvM3FuTGqj+KCA6EenYdxLm0TWylJNsnXJxB1raCrEQXm/Yf6YTu+vok9aHgDx5mZy4e7BHweO8gAdGmLBSip77QObE3rs5jD7{BwoOVPdfSDwRiknBd+wkSlmebT5R3GN17P5g5IdLRlV3I4RaqgehRT/nt82pTmOMHFzsI9/OdVWaE1XCqFZV9Qiz50jHj9zauhv4yAs8cIV2+GYp7ESNoS0fU9H{oXmisa/f+1ITAvvguelQro5JV704wQMnIk/m5lVHGv/vfyFODkZclbygLxQJGbgd3fFIxM9D0npcM6QGpwbfFYcy+BL99ZaUNJSt2KLw+jGNTYu{a+ErUhEABh5mopIc6wButzatBoBH/Yoh3pErTpO4n7kvR65OdrPSFloFefaAV5IkqGHgspy2V11mL7WhpVd1LP054JRbJjLWHcAd/GZFGVs6B4iFmzdBLobhRC9L6Yt1e73TDSEc4fvgkEbkZmjsjceAX8YlfzTKhpdQIp7VghdUl3lILWspDu80vcR3YgG48ZOqG4X0bGiZmRUS6/Revs6Vb1JzruLWPh8FnIKs47xeEAIOmv4wReDsDDueA8e6KURS15lmwzKESbBnADMvMs3UZ5QjKQib8pJMuIHXbRSJd6SjYnlBOtnJCOCy7gZrXEcbgNx9dF{fbapkZ8KbNDDEEsMlcG7uwOPB3nxxQShNz14Dab6TcRGxeG7Bpy6ggIS5RbYgRMEnNYuqQWm0LgB1waxV0qk1mOjtCIVi7XsHsop1hxebNvaPX3IkP{6N28uS9lMdMWu7sat9sCzgYUiWYHoJiB2Dj{b{ETDuGeYS2CG7LRGyU1Z90SyymrIwNIP9cNHO923bFQcbgw2j0aLH2onplXfFPSOqig0nK3QxBTnwAcUU0OF6pRWfNDULZLMEMTmMwE1NqUxjq5fVJYHFyfGOxjMr/yqSlRNDL5d3JHwckgclIxmknoHF4CsH6gM7IMBPt1p{EvgUJ{B23t9WdtCloSW6ZEQ63i7YvzJRmt5pHsAB6sCHLMCWl+bn4ebHIrizT56rYXb65Es9uLy7PxqoJ/zfXoKqw"},
		"__VIEWSTATEGENERATOR": {"2A35A6B2"},
		"__VIEWSTATEENCRYPTED": {""},
		"__EVENTVALIDATION":    {"Q9T5K+q5GQok4lkig5Kosdhxc5R93ZaH0JRLjAeOeZqeGyur/++QO4IWJR5Uh/o6uiUQhecB+3U7ORX687faqFLC7NwTBxUM5ieILAZ8rW1RAi3tIjqvTglI1inh5ZNvaRcNk0/sAGyDl4VXcMKbXGXCJgQ8Wt7OT7LXs8zAgt8YOIAhjw{JHKZPg5gZ+iA3juhnthxmZfN/x7+4hpgKPZ5to+g{"},
		"tep_name":             {""},
		"organ_name":           {""},
		"site_address":         {""},
		"__EVENTARGUMENT":      {"1"},
		"ddlPageCount":         {"10"},
	}

	resp, e := http.PostForm(posturl, form)
	if e != nil {
		log.Fatal(e)
	}
	return resp

}

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

		n , e := fo.WriteString(string(l) + "\n")

		fo.Sync()
		defer fo.Close()

		if e != nil{
			log.Println(e)
		}
		log.Println("write line:" + string(n) + string(l))
	}
}

func parseData(doc *goquery.Document) {
	projects := doc.Find("table").Find("tr[class]")

	projects.Each(func(pNum int, selection *goquery.Selection) {
		log.Printf("count:%d \n", count)
		count++

		one := selection.Find("td")

		b := ProjectBrief{}

		one.Each(func(num int, sel *goquery.Selection) {

			val := strings.TrimSpace(sel.Text())

			log.Printf("num:%d %s \n", num, val)

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
