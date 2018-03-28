package crawler

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	model "github.com/visheratin/ico-crawler/model/icorating"
	"github.com/visheratin/ico-crawler/writer"
)

type ICORatingWorker struct {
	id       int
	finished bool
	pageType string
	links    []string
}

func (worker *ICORatingWorker) Start() error {
	for _, link := range worker.links {
		entity, _ := worker.GetDetails(link)
		outputPath := "./data3/icorating/"
		outFilename := entity.Title + ".json"
		writer.WriteToFS(outputPath, outFilename, entity)
	}
	return nil
}

func (worker *ICORatingWorker) GetDetails(detailsLink string) (model.ICORatingCompany, error) {
	doc, err := goquery.NewDocument("https://ru.investing.com" + detailsLink + "/markets")
	if err != nil {
		return model.ICORatingCompany{}, err
	}
	result := model.ICORatingCompany{}
	titleNode := doc.Find("h1")
	if len(titleNode.Nodes) > 0 {
		result.Title = clearText(titleNode.Text())
	}

	tableCells := doc.Find("td:first-child + td + td, td:first-child + td + td + td + td, td:first-child + td + td + td + td + td, td:first-child + td + td + td + td + td +td + td")

	for i := 1; i < len(tableCells.Nodes); i++ {

		m := model.Market{}

		cell1 := tableCells.Eq(i)
		m.Name = clearText(cell1.Text())

		i += 1
		cell2 := tableCells.Eq(i)
		m.Max = clearText(cell2.Text())

		i += 1
		cell3 := tableCells.Eq(i)
		m.Min = clearText(cell3.Text())

		i += 1
		cell4 := tableCells.Eq(i)
		m.Volume = clearText(cell4.Text())

		result.Markets = append(result.Markets, m)
	}

	return result, nil
}

func clearText(input string) string {
	output := strings.Replace(input, "\n", "", -1)
	output = strings.TrimSpace(output)
	return output
}

// func (worker *ICORatingWorker) GetNews(link string) (interface{}, error) {

// }

// func (worker *ICORatingWorker) GetReview(link string) (interface{}, error) {

// }
