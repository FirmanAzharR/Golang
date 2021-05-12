package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Tracking struct {
	Status Status `json:"status"`
	Data   Data   `json:"data"`
}
type Status struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type Data struct {
	ReceivedBy string   `json:"receivedBy"`
	Histories  []string `json:"histories"`
}

type DataHistory struct {
	Description string `json:"description"`
	CreatedAt   string `json:"createAt"`
	//formated    Format `json:"formated"`
}
type Format struct {
	Formated string `json:"formated"`
}

func main() {
	res, err := http.Get("https://gist.githubusercontent.com/nubors/eecf5b8dc838d4e6cc9de9f7b5db236f/raw/d34e1823906d3ab36ccc2e687fcafedf3eacfac9/jne-awb.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	} else {

		stat := Status{Code: "060101", Message: "Delivery tracking detail fetched successfully"}
		dataTrack := Data{ReceivedBy: "PAK MURADI"}
		rows := make([]DataHistory, 0)

		doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
			tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
				rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
					row := new(DataHistory)
					row.Description = tablecell.Text()
					row.CreatedAt = tablecell.Text()
					rows = append(rows, *row)
				})
			})
		})
		track := Tracking{Status: stat, Data: dataTrack}
		bts, err := json.MarshalIndent(track, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(bts))
	}

}

// rows = append(rows, row)
// row = nil

// byteArr, err := json.MarshalIndent(track, "", "  ")
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Println(string(byteArr))

// doc.Find(".tracking").Children().Each(func(i int, sel *goquery.Selection) {
// 	row := new(Article)
// 	row.Title = sel.Find("td").Text()
// 	rows = append(rows, *row)
// })
