# Simple Data Analysis and Comparison Tool with Go

A simple program that will allow you to compare or analyze data sets of different types or the same type. I made this project to learn Go.

###

<img height="250" width="100%" src="https://cdn.discordapp.com/attachments/855396331194286081/1319672437977251891/goland64_8M90EZ0O56.gif?ex=6766d03b&is=67657ebb&hm=223ed9f9806916021547d7837bd473c268a39b5e4fdd5ae000c2ca79e8c36488&"  />

###


## Installation: 
```shell

# Fetching
git clone https://github.com/BotEkrem/data-analysis-comparison-tool.git
cd data-analysis-comparison-tool

# Running
go build
./dataAnalyzer

```

## Directories:
#### factories:
There is only one file here: `data.factory.go`. In this file, validations are made regarding the data entered into the program. I wrote 5 validation functions, you can add new ones or customize validations from this file.

```go
const (
	KeysLengthCheck ValidationType = iota
	ValuesLengthCheck
	PrimaryKeyCheck
	KeysEqualityCheck
	ValuesEqualityCheck
	// ...
)
```

#### adapters:  
There are the data types adapted according to `data.factory.go` here. I wrote for CSV and JSON. You can add new types here.
#### data:
Sample data used for the program is here.
#### utils:
Some utility functions used for the program are here.