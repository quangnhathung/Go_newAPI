package database

import (
	"log"
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"

)

type vcb struct{
	Vocab 		string
	identify 	string
}
func List_vocab()  []vcb{
		url := "https://emojiflashcards.com/english/school-things/picture-dictionary?collection_id=31&view=grid"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal(res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
    var vocabs []vcb
	doc.Find("li.item").Each(func(index int, element *goquery.Selection) {
		var vocab vcb
		vocab.identify,_ = element.Attr("data-collection-ids")
		vocab.Vocab = element.Find("span").Text()
		fmt.Printf("%s:%s\n\n",vocab.Vocab,vocab.identify)
        vocabs = append(vocabs,vocab)
	})
	return vocabs
}

type ctgr struct{
	Id 		string
	Name 	string
}

func List_category()  []ctgr{
		url := "https://emojiflashcards.com/english/school-things/picture-dictionary?collection_id=31&view=grid"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal(res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
    var category []ctgr
	doc.Find("option").Each(func(index int, element *goquery.Selection) {
		var ca ctgr
		ca.Name = element.Text()
		ca.Id, _ = element.Attr("value")
		fmt.Printf("%s:%s\n\n",ca.Id,ca.Name)
        category = append(category,ca)
	})
	category[0].Id="0";
	return category
}

