package main

import (
	"fmt"
	"github.com/diogorodriguesc/lutetium-go"
	"gopkg.in/yaml.v2"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

func printValues(file File, wg *sync.WaitGroup) {
	defer wg.Done()

	f := &lutetiumgo.Sitemap{
		lutetiumgo.Xml{
			Location: file.path,
		},
	}

	UrlSet := f.Read()

	for _, c := range UrlSet.Urls {
		fmt.Fprintf(os.Stdout, "%s\n", c.Loc)
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

	for _, f := range files {
		go printValues(f, &wg)
	}

	wg.Wait()
}
