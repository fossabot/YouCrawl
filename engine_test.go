package youcrawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"testing"
)

func TestRequestWithURL(t *testing.T) {
	_, err := RequestWithURL(&Task{
		Url: "https://www.zhihu.com",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestEngine_Run(t *testing.T) {
	e := NewEngine(&EngineOption{MaxRequest: 3})
	e.AddURLs("https://www.zhihu.com")
	e.AddHTMLParser(func(doc *goquery.Document, ctx Context) {
		title := doc.Find("title").Text()
		fmt.Println(fmt.Sprintf("%s [%d]",ctx.Request.URL.String(),ctx.Response.StatusCode))
		fmt.Println(title)
	})
	stopChannel := make(chan struct{})
	e.Run(stopChannel)
	<-stopChannel
}

func TestParseHTML(t *testing.T) {
	bodyReader, err := RequestWithURL(&Task{
		Url: "https://www.zhihu.com",
	})
	if err != nil {
		t.Error(err)
	}
	err = ParseHTML(bodyReader, func(doc *goquery.Document, ctx Context) {
		title := doc.Find("title").Text()
		fmt.Println(title)
	}, Context{content: map[string]interface{}{}})
	if err != nil {
		t.Error(err)
	}
}
