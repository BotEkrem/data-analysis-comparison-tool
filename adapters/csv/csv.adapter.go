package csv

import (
	"dataAnalyzer/factories"
	"dataAnalyzer/utils"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type CSVType [][]string

func ReadCSV(fileName string) CSVType {
	dirname := utils.GetCurrentPath()
	completePath := fmt.Sprintf("%s/../data/%s", dirname, fileName)

	csvFile, err := os.Open(completePath)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully Opened: %s\n", completePath)
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			panic(err)
		}
	}(csvFile)

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	return records
}

func (csv CSVType) CSVtoRecord() factories.DataRecords {
	var keyData []factories.DataKey
	var valueData [][]factories.DataValue

	for i := 0; i < len(csv); i++ {
		var rowData []factories.DataValue
		values := strings.Split(csv[i][0], ";")
		if i == 0 {
			for _, k := range values {
				keyData = append(keyData, factories.DataKey{DisplayName: k, MinifiedName: strings.ReplaceAll(strings.ToLower(k), " ", "")})
			}
		} else {
			for vi, v := range values {
				value := factories.DataValue{KeyName: keyData[vi].MinifiedName, Value: v}
				rowData = append(rowData, value)
			}
			valueData = append(valueData, rowData)
		}
	}

	return factories.DataRecords{Keys: keyData, Values: valueData}
}
