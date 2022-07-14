package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	read()
	readAll()
}

func readAll() {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	_ = records
}

func read() {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		_ = record
	}
}
