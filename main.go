package main

import (
	"dataAnalyzer/adapters/csv"
	"dataAnalyzer/adapters/json"
	"dataAnalyzer/factories"
	"dataAnalyzer/utils"
	"fmt"
	"os"
)

func main() {
	factory := factories.DataFactory{Datasets: make([]factories.DataRecords, 0), PrimaryKey: "id"}
	adapter, _ := utils.GetDirectories("./adapters")

	for {
		dataType := utils.Selection(append(utils.Map(adapter, func(t os.DirEntry) string { return t.Name() }), "Finish"), fmt.Sprintf("Select a data type (Use Enter to select) \n Current Dataset: %v", factory.Datasets), utils.Cyan, utils.Blue, utils.White)

		if dataType == "Finish" {
			utils.ClearConsole()
			fmt.Println("The analysis process will be started soon.")
			break
		}

		files := utils.FindByExtension("./data", fmt.Sprintf(".%v", dataType))
		dataFile := utils.Selection(files, "Select a data file (Use Enter to select)", utils.Cyan, utils.Blue, utils.White)

		switch dataType {
		case "csv":
			csvData := csv.ReadCSV(dataFile)
			factory.Adapter(csvData.CSVtoRecord())
			break
		case "json":
			jsonData := json.ReadJSON(dataFile)
			factory.Adapter(jsonData.JSONtoRecord())
			break
		}

		fmt.Printf("%v", factory.Datasets)
	}

	factory.Verify()
}
