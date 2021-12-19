package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
)

type rowData struct {
	D_1 string  `parquet:"name=d_1, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	D_2 string  `parquet:"name=d_2, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	D_3 string  `parquet:"name=d_3, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	D_4 string  `parquet:"name=d_4, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Day string  `parquet:"name=day, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	M_1 float64 `parquet:"name=m_1, type=FLOAT"`
	M_2 float64 `parquet:"name=m_2, type=FLOAT"`
	M_3 float64 `parquet:"name=m_3, type=FLOAT"`
}

func main() {
	columnMaps := map[string]string{
		"media_account_uid": "d_1",
		"campaign_id":       "d_2",
		"adset_id":          "d_3",
		"ad_id":             "d_4",
		"day":               "day",
		"clicks":            "m_1",
		"impressions":       "m_2",
		"cost":              "m_3",
	}

	rf, err := os.Open("../data/large.log")
	if err != nil {
		log.Fatal(err)
	}
	defer rf.Close()

	reader := bufio.NewReader(rf)

	wf, err := local.NewLocalFileWriter("../data/output-large-golang.parquet")
	if err != nil {
		log.Fatal(err)
	}

	pw, err := writer.NewParquetWriter(wf, new(rowData), 4)
	if err != nil {
		log.Fatal(err)
	}

	pw.RowGroupSize = 128 * 1024 * 1024
	pw.PageSize = 8 * 1024
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	count := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		var lineData map[string]interface{}
		json.Unmarshal(line, &lineData)

		var tD = make(map[string]string)
		var tM = make(map[string]float64)
		for c, k := range columnMaps {
			v := lineData[c]
			if v == nil {
				continue
			}

			switch x := v.(type) {
			case string:
				tD[k] = x
			case float64:
				tM[k] = x
			}
		}

		item := rowData{
			D_1: tD["d_1"],
			D_2: tD["d_2"],
			D_3: tD["d_3"],
			D_4: tD["d_4"],
			Day: tD["day"],
			M_1: tM["m_1"],
			M_2: tM["m_2"],
			M_3: tM["m_3"],
		}

		if err = pw.Write(item); err != nil {
			log.Println("Write error", err)
		}

		count++

		fmt.Printf("Row: %d \n", count)
	}

	if err = pw.WriteStop(); err != nil {
		log.Println("WriteStop error", err)
		return
	}

	log.Println("Write Finished")
	wf.Close()
}
