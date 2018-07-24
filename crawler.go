package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {
	resp, err := soup.Get("https://www.ptt.cc/bbs/nba/index.html")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("div", "class","r-ent")
	for _, link := range links {
		fmt.Println("標題:",link.Find("a").Text()+"\n", 
					"網址 :", "https://www.ptt.cc/"+link.Find("a").Attrs()["href"]+"\n",
					"作者 :", link.Find("div","class","author").Text()+"\n",
					"推文數 ：", link.Find("span").Text()+"\n",
					"Date :", link.Find("div","class","date").Text()+"\n",
				)
	}
}