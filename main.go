package main

import (
	cw "github.com/mainak55512/TechNews-Scrapper/csvwriter"
	ml "github.com/mainak55512/TechNews-Scrapper/emailer"
	sc "github.com/mainak55512/TechNews-Scrapper/scrapper"
)

func main() {
	cw.InitCsv()
	cw.WriteCsv(sc.NewsScrapper())
	ml.SendMail()
}
