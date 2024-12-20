package factories

import (
	"dataAnalyzer/utils"
	"fmt"
	"reflect"
	"sort"
)

type DataFactory struct {
	Datasets   []DataRecords
	PrimaryKey string
}

type DataRecords struct {
	Keys   []DataKey
	Values [][]DataValue
}

type DataKey struct {
	DisplayName  string
	MinifiedName string
}

type DataValue struct {
	KeyName string
	Value   any
}

func (factory *DataFactory) Adapter(data DataRecords) {
	factory.Datasets = append(factory.Datasets, data)
}

func (factory *DataFactory) Verify() {
	if len(factory.Datasets) == 0 {
		panic("No datasets found.")
	}

	for index, dataset := range factory.Datasets {
		utils.ColoredLog(true, fmt.Sprintf("No. %v Dataset: %v", index+1, dataset), utils.White)

		factory.Validate(KeysLengthCheck, dataset, index)
		factory.Validate(ValuesLengthCheck, dataset, index)
		factory.Validate(PrimaryKeyCheck, dataset, index)
		factory.Validate(KeysEqualityCheck, dataset, index)
		factory.Validate(ValuesEqualityCheck, dataset, index)
		utils.ColoredLog(true, fmt.Sprintf("No. %v dataset verified.", index+1), utils.Green)
	}

	utils.ColoredLog(true, "All datasets verified.", utils.Green)
}

type ValidationType int

const (
	KeysLengthCheck ValidationType = iota
	ValuesLengthCheck
	PrimaryKeyCheck
	KeysEqualityCheck
	ValuesEqualityCheck
	// ...
)

func (factory *DataFactory) Validate(validationType ValidationType, dataset DataRecords, datasetIndex int) {
	var totalKeysCount = len(factory.Datasets[0].Keys)
	var totalValuesCount = len(factory.Datasets[0].Values[0])
	var keys = utils.Map(factory.Datasets[0].Keys, func(k DataKey) string { return k.MinifiedName })
	sort.Strings(keys)

	var values [][]DataValue
	for _, value := range factory.Datasets[0].Values {
		values = append(values, value)
	}

	switch validationType {
	case KeysLengthCheck:
		if len(dataset.Keys) != totalKeysCount {
			panic(fmt.Sprintf("Keys Count Check Exception in No. %v | Dataset: %v", datasetIndex+1, dataset))
		}
		break
	case ValuesLengthCheck:
		for index, value := range dataset.Values {
			if len(value) != totalValuesCount {
				panic(fmt.Sprintf("Values Count Check Exception in No. %v | Index: %v | Dataset: %v", datasetIndex+1, index, dataset))
			}
		}
		break
	case PrimaryKeyCheck:
		for index, value := range dataset.Values {
			idx := len(value)
			for innerIndex, innerValue := range value {
				if innerValue.KeyName == factory.PrimaryKey {
					idx = innerIndex
				}
			}

			//idx := sort.Search(len(value), func(i int) bool {
			//	println(i)
			//	fmt.Printf("%v == %v | Index: %v | Len: %v\n", value[i].KeyName, factory.PrimaryKey, i, len(value))
			//	return value[i].KeyName == factory.PrimaryKey
			//})

			if idx == len(value) {
				panic(fmt.Sprintf("Missing Primary Key Exception in No. %v | Index: %v | Dataset: %v", datasetIndex+1, index, dataset))
			}
		}
		break
	case KeysEqualityCheck:
		var datasetKeys = utils.Map(dataset.Keys, func(k DataKey) string { return k.MinifiedName })
		sort.Strings(datasetKeys)
		if !reflect.DeepEqual(keys, datasetKeys) {
			panic(fmt.Sprintf("Keys Equality Exception in No. %v | Dataset: %v", datasetIndex+1, dataset))
		}
		break
	case ValuesEqualityCheck:
		for index, value := range values {
			var globalValueStrings = utils.Map(value, func(value DataValue) any { return value.Value })
			var localValueStrings = utils.Map(dataset.Values[index], func(value DataValue) any { return value.Value })

			if !utils.Equal(globalValueStrings, localValueStrings) {
				panic(fmt.Sprintf("Values Equality Exception in No. %v | Index: %v | Dataset: %v", datasetIndex+1, index, dataset))
			}
		}
		break
	}
}
