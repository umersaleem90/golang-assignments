package htmlParser

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

func ParseHtml(fileName string) LinkData{
    filePath := "htmlParser/" + fileName
	file, error := ioutil.ReadFile(filePath)
	if error != nil {
		log.Fatal(error)
	}
	
	r := strings.NewReader(string(file))
	document, error := html.Parse(r)
	if error != nil {
		log.Fatal(error)
	}	

	htmlNodes := cascadia.MustCompile("a").MatchAll(document)

	var href, text string
	for _, node := range htmlNodes {
		for _, attribute := range node.Attr {
			if attribute.Key == "href" {
				href += attribute.Val + " "
			} 
		}
		text += strings.TrimSpace(node.FirstChild.Data) + " "
	}

	return LinkData{href, text}
}