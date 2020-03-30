package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Rules struct which contains
// an array of rules
type Rules struct {
	Rules []Rule `json:"rules"`
}

// Rule struct which contains two array
type Rule struct {
	Test  string   `json:"test"`
	XAxis []string `json:"xAxis"`
	YAxis []string `json:"yAxis"`
}

// Find - Search element in a list
// Return index and status from a list
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// findMatched - recursive function. Find [x,y] from a list
// Return status
func findMatched(rules Rules, index int, x string, y string) bool {

	if index >= len(rules.Rules) {
		return false
	}

	_, xExisted := Find(rules.Rules[index].XAxis, x)
	_, yExisted := Find(rules.Rules[index].YAxis, y)

	if xExisted && yExisted {
		fmt.Printf("\n***** Capture the data [%s,%s] ******\n", x, y)
		fmt.Println("Rule X-Axis Value: ", rules.Rules[index].XAxis)
		fmt.Println("Rule Y-Axis Value: ", rules.Rules[index].YAxis)

		return true
	} else {
		index++
		return findMatched(rules, index, x, y)
	}
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("rules.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened rules.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var rules Rules

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'rules' which we defined above
	json.Unmarshal(byteValue, &rules)

	// we iterate through every rule within our rules array
	for i := 0; i < len(rules.Rules); i++ {
		fmt.Println("Rule X-Axis Value: ", rules.Rules[i].XAxis)
		fmt.Println("Rule Y-Axis Value: ", rules.Rules[i].YAxis)
	}

	test := findMatched(rules, 0, "111", "ccc")
	fmt.Println("\nHi, the result is - ", test)
}