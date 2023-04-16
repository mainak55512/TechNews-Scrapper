package csvwriter

import (
	"encoding/csv"
	"log"
	"os"
)

func InitCsv() {
	f, err := os.Create("tectnews.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	if err := w.Write([]string{"News Header", "Article Link"}); err != nil {
		log.Fatalln("Can not set the headers ", err)
	}
}

func WriteCsv(records [][][]string) {
	f, err := os.OpenFile("tectnews.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()
	w := csv.NewWriter(f)

	for _, record := range records {
		err = w.WriteAll(record)
		if err != nil {
			log.Fatal(err)
		}
	}
}
