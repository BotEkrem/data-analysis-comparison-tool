package json

import (
	"dataAnalyzer/factories"
	"dataAnalyzer/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type JSONType []map[string]any

func ReadJSON(fileName string) JSONType {
	dirname := utils.GetCurrentPath()
	completePath := fmt.Sprintf("%s/../data/%s", dirname, fileName)

	jsonFile, err := os.Open(completePath)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully Opened: %s\n", completePath)
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	var output JSONType

	err = json.Unmarshal(byteValue, &output)

	if err != nil {
		panic(err)
	}

	return output
}

func (json JSONType) JSONtoRecord() factories.DataRecords {
	var keyData []factories.DataKey
	var valueData [][]factories.DataValue

	if len(json) == 0 {
		panic("No data found.")
	}

	keys := utils.Keys(json[0])

	for _, k := range keys {
		key := factories.DataKey{DisplayName: k, MinifiedName: strings.ReplaceAll(strings.ToLower(k), " ", "")}
		keyData = append(keyData, key)
	}

	for i := 0; i < len(json); i++ {
		var rowData []factories.DataValue
		for l := 0; l < len(keyData); l++ {
			value := factories.DataValue{KeyName: keyData[l].MinifiedName, Value: json[i][keyData[l].MinifiedName]}
			rowData = append(rowData, value)
		}
		valueData = append(valueData, rowData)
	}

	return factories.DataRecords{Keys: keyData, Values: valueData}
}
