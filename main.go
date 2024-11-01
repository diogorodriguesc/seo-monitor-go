package main

import (
	"fmt"
	"github.com/diogorodriguesc/lutetium-go"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

func GetMetaRobotsContent(res *http.Response) string {
	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var robotsContent string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			// Check if the meta tag has name="robots"
			var nameAttr, contentAttr string
			for _, attr := range n.Attr {
				if attr.Key == "name" && strings.ToLower(attr.Val) == "robots" {
					nameAttr = attr.Val
				}
				if attr.Key == "content" {
					contentAttr = attr.Val
				}
			}
			// If found, set robotsContent
			if nameAttr == "robots" {
				robotsContent = contentAttr
				return
			}
		}
		// Recursively apply to children
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	if robotsContent == "" {
		fmt.Errorf("robots meta tag not found")
	}

	return robotsContent
}

func ProcessXmlFile(XmlFile XmlFile, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%+v\n", XmlFile)

	f := &lutetiumgo.Sitemap{
		lutetiumgo.Xml{
			Location: XmlFile.File,
		},
	}

	UrlSet := f.Read()

	for _, c := range UrlSet.Urls {
		requestUrl := fmt.Sprintf(c.Loc)
		res, err := http.Get(requestUrl)

		if err != nil {
			fmt.Println(err)
		}

		defer res.Body.Close()

		robotsContent := GetMetaRobotsContent(res)

		fmt.Printf("url: %s client: status code: %d robots: %s\n", c.Loc, res.StatusCode, robotsContent)
	}
}

func getConf() interface{} {

	obj := make(map[string]interface{})

	yamlFile, err := os.ReadFile("parameters.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, obj)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}

	return obj
}

func main() {
	conf := getConf()
	files := getAllFiles(conf)

	var wg sync.WaitGroup

	wg.Add(len(files))

	for _, XmlFile := range files {
		go ProcessXmlFile(XmlFile, &wg)
	}

	wg.Wait()
}
