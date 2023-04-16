package main

import (
	sc "tectnews-scrapper/scrapper"
	cw "tectnews-scrapper/csvwriter"
	ml "tectnews-scrapper/emailer"
)

func main() {
	cw.InitCsv()
	cw.WriteCsv(sc.NewsScrapper())
	ml.SendMail()
}
